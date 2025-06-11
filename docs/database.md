# ZCalendar Database Design

## Database Technology
- **Primary Database**: PostgreSQL 14+
- **Caching Layer**: Redis  
- **Connection Pooling**: PgBouncer
- **Local Storage**: Core Data (iOS), Room (Android)

## Architecture Overview

ZCalendar supports dual storage modes:

### Guest Mode
- **Local only**: Data stored in device database (Core Data/Room)
- **No server storage**: No user accounts or cloud data
- **Privacy focused**: Zero network data transmission

### Registered User Mode  
- **Hybrid storage**: Local database + PostgreSQL cloud storage
- **Synchronization**: Background sync between local and cloud
- **Multi-device**: Access from multiple registered devices

## Cloud Database Schema (PostgreSQL)

### Core Tables

#### users
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    timezone VARCHAR(50) DEFAULT 'UTC',
    subscription_tier VARCHAR(20) DEFAULT 'free', -- 'free', 'premium'
    registration_source VARCHAR(20) DEFAULT 'app', -- 'app', 'web', 'migration'
    migrated_data_count INTEGER DEFAULT 0, -- Number of items migrated from guest mode
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    last_sync_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_subscription_tier ON users(subscription_tier);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
```

#### calendars
```sql
CREATE TABLE calendars (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    color VARCHAR(7) DEFAULT '#2196F3', -- Hex color code
    is_primary BOOLEAN DEFAULT FALSE,
    is_shared BOOLEAN DEFAULT FALSE,
    share_permissions JSONB DEFAULT '{}', -- Sharing configuration for premium users
    migration_source_id VARCHAR(255), -- Original local ID for migrated calendars
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX idx_calendars_user_id ON calendars(user_id);
CREATE INDEX idx_calendars_migration_source ON calendars(migration_source_id);
CREATE INDEX idx_calendars_deleted_at ON calendars(deleted_at);
CREATE INDEX idx_calendars_shared ON calendars(is_shared) WHERE is_shared = true;
```

#### events
```sql
CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    calendar_id UUID NOT NULL REFERENCES calendars(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE NOT NULL,
    all_day BOOLEAN DEFAULT FALSE,
    location VARCHAR(500),
    recurrence_rule TEXT, -- RRULE format
    migration_source_id VARCHAR(255), -- Original local ID for migrated events
    invited_users JSONB DEFAULT '[]', -- Email list for premium sharing feature
    attachments JSONB DEFAULT '[]', -- File attachments for premium users
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    
    CONSTRAINT events_time_check CHECK (end_time > start_time)
);

CREATE INDEX idx_events_calendar_id ON events(calendar_id);
CREATE INDEX idx_events_start_time ON events(start_time);
CREATE INDEX idx_events_end_time ON events(end_time);
CREATE INDEX idx_events_migration_source ON events(migration_source_id);
CREATE INDEX idx_events_deleted_at ON events(deleted_at);
CREATE INDEX idx_events_time_range ON events(start_time, end_time);
CREATE INDEX idx_events_invited_users ON events USING GIN(invited_users) WHERE jsonb_array_length(invited_users) > 0;
```

#### event_reminders
```sql
CREATE TABLE event_reminders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL, -- 'notification', 'email', 'sms'
    minutes_before INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_event_reminders_event_id ON event_reminders(event_id);
```

#### calendar_shares
```sql
CREATE TABLE calendar_shares (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    calendar_id UUID NOT NULL REFERENCES calendars(id) ON DELETE CASCADE,
    shared_with_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    permission_level VARCHAR(20) NOT NULL, -- 'read', 'write', 'admin'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(calendar_id, shared_with_user_id)
);

CREATE INDEX idx_calendar_shares_calendar_id ON calendar_shares(calendar_id);
CREATE INDEX idx_calendar_shares_user_id ON calendar_shares(shared_with_user_id);
```

#### refresh_tokens
```sql
CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    revoked_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);
```

#### user_sync_metadata
```sql
CREATE TABLE user_sync_metadata (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    last_sync_timestamp TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    sync_version INTEGER DEFAULT 1,
    device_id VARCHAR(255), -- Unique device identifier
    device_info JSONB DEFAULT '{}', -- Device type, OS version, app version
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(user_id, device_id)
);

CREATE INDEX idx_user_sync_metadata_user_id ON user_sync_metadata(user_id);
CREATE INDEX idx_user_sync_metadata_last_sync ON user_sync_metadata(last_sync_timestamp);
```

#### external_integrations (Premium Feature)
```sql
CREATE TABLE external_integrations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    provider VARCHAR(50) NOT NULL, -- 'google', 'outlook', 'apple'
    provider_user_id VARCHAR(255) NOT NULL,
    access_token_hash VARCHAR(255) NOT NULL,
    refresh_token_hash VARCHAR(255),
    scope TEXT, -- OAuth scopes granted
    sync_enabled BOOLEAN DEFAULT TRUE,
    sync_direction VARCHAR(20) DEFAULT 'bidirectional', -- 'import', 'export', 'bidirectional'
    last_sync_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(user_id, provider, provider_user_id)
);

CREATE INDEX idx_external_integrations_user_id ON external_integrations(user_id);
CREATE INDEX idx_external_integrations_provider ON external_integrations(provider);
CREATE INDEX idx_external_integrations_sync_enabled ON external_integrations(sync_enabled) WHERE sync_enabled = true;
```

#### shared_event_invitations (Premium Feature)
```sql
CREATE TABLE shared_event_invitations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    inviter_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    invitee_email VARCHAR(255) NOT NULL,
    invitee_user_id UUID REFERENCES users(id) ON DELETE SET NULL, -- NULL if not registered
    status VARCHAR(20) DEFAULT 'pending', -- 'pending', 'accepted', 'declined'
    message TEXT,
    invitation_token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    responded_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(event_id, invitee_email)
);

CREATE INDEX idx_shared_event_invitations_event_id ON shared_event_invitations(event_id);
CREATE INDEX idx_shared_event_invitations_invitee_email ON shared_event_invitations(invitee_email);
CREATE INDEX idx_shared_event_invitations_token ON shared_event_invitations(invitation_token);
CREATE INDEX idx_shared_event_invitations_expires ON shared_event_invitations(expires_at);
```

## Local Database Schema (Mobile Apps)

### Core Data (iOS) / Room (Android)

The local mobile database mirrors the cloud schema but includes additional sync metadata:

#### Local Events Table
```sql
CREATE TABLE events (
    id TEXT PRIMARY KEY, -- Local UUID
    server_id TEXT, -- Cloud UUID when synced
    calendar_id TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    all_day BOOLEAN DEFAULT 0,
    location TEXT,
    recurrence_rule TEXT,
    sync_status TEXT DEFAULT 'local', -- 'local', 'synced', 'pending', 'conflict'
    last_sync_version INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    FOREIGN KEY (calendar_id) REFERENCES calendars(id)
);
```

#### Local Calendars Table
```sql
CREATE TABLE calendars (
    id TEXT PRIMARY KEY, -- Local UUID
    server_id TEXT, -- Cloud UUID when synced  
    name TEXT NOT NULL,
    description TEXT,
    color TEXT DEFAULT '#2196F3',
    is_primary BOOLEAN DEFAULT 0,
    sync_status TEXT DEFAULT 'local',
    last_sync_version INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME
);
```

#### Sync Metadata Table
```sql
CREATE TABLE sync_metadata (
    key TEXT PRIMARY KEY,
    value TEXT,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Store sync timestamps and user info
INSERT INTO sync_metadata VALUES ('user_id', NULL, CURRENT_TIMESTAMP);
INSERT INTO sync_metadata VALUES ('last_sync_events', NULL, CURRENT_TIMESTAMP);
INSERT INTO sync_metadata VALUES ('last_sync_calendars', NULL, CURRENT_TIMESTAMP);
INSERT INTO sync_metadata VALUES ('auth_token', NULL, CURRENT_TIMESTAMP);
INSERT INTO sync_metadata VALUES ('device_id', 'device_unique_id', CURRENT_TIMESTAMP);
```

#### Pending Operations Table (Sync Queue)
```sql
CREATE TABLE pending_operations (
    id TEXT PRIMARY KEY,
    entity_type TEXT NOT NULL, -- 'event', 'calendar'
    entity_id TEXT NOT NULL, -- Local entity ID
    operation_type TEXT NOT NULL, -- 'create', 'update', 'delete'
    operation_data TEXT, -- JSON payload for operation
    retry_count INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

### Sync Status Values

- **`local`**: Created offline, not uploaded to cloud
- **`synced`**: Successfully synchronized with cloud
- **`pending`**: Modified locally, queued for upload
- **`conflict`**: Conflict detected during sync, requires resolution

## Relationships

```
users (1) ──────────── (*) calendars
  │                        │
  │                        │
  └── (*) calendar_shares   │
                           │
calendars (1) ──────────── (*) events
                             │
                             │
events (1) ─────────────── (*) event_reminders
```

## Indexing Strategy

### Performance Indexes
- **Time-based queries**: Indexes on start_time, end_time for efficient range queries
- **User data**: Indexes on user_id foreign keys for fast user-specific queries
- **Soft deletes**: Indexes on deleted_at columns for filtering active records

### Composite Indexes
```sql
-- For event queries with time ranges
CREATE INDEX idx_events_calendar_time ON events(calendar_id, start_time, end_time) 
    WHERE deleted_at IS NULL;

-- For user calendar access
CREATE INDEX idx_calendars_user_active ON calendars(user_id, is_primary) 
    WHERE deleted_at IS NULL;
```

## Data Types and Constraints

### Timezone Handling
- All timestamps stored as `TIMESTAMP WITH TIME ZONE`
- User timezone preference stored separately
- Client applications handle timezone conversion

### Soft Deletes
- All main entities use `deleted_at` for soft deletion
- Queries filter by `deleted_at IS NULL` for active records
- Cleanup jobs periodically hard delete old soft-deleted records

### Data Validation
- Email format validation at application level
- Color codes must be valid hex format
- Event times must be logical (end > start)

## Migration Strategy

### Version Control
- Database migrations using tools like Flyway or Liquibase
- Sequential versioning for all schema changes
- Rollback scripts for each migration

### Sample Migration
```sql
-- V001__initial_schema.sql
-- Creates initial tables and indexes

-- V002__add_event_categories.sql
ALTER TABLE events ADD COLUMN category_id UUID REFERENCES categories(id);
CREATE INDEX idx_events_category ON events(category_id);
```

## Performance Considerations

### Query Optimization
- Prepared statements for common queries
- Connection pooling for database connections
- Query result caching for frequently accessed data

### Monitoring
- Slow query logging enabled
- Database performance metrics collection
- Regular EXPLAIN ANALYZE on critical queries

### Backup Strategy
- Daily full backups
- Continuous WAL archiving
- Point-in-time recovery capability
- Cross-region backup replication