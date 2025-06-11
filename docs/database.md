# ZCalendar Database Design

## Database Technology
- **Primary Database**: PostgreSQL 14+
- **Caching Layer**: Redis
- **Connection Pooling**: PgBouncer

## Schema Overview

### Core Tables

#### users
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    timezone VARCHAR(50) DEFAULT 'UTC',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX idx_users_email ON users(email);
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
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX idx_calendars_user_id ON calendars(user_id);
CREATE INDEX idx_calendars_deleted_at ON calendars(deleted_at);
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
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    
    CONSTRAINT events_time_check CHECK (end_time > start_time)
);

CREATE INDEX idx_events_calendar_id ON events(calendar_id);
CREATE INDEX idx_events_start_time ON events(start_time);
CREATE INDEX idx_events_end_time ON events(end_time);
CREATE INDEX idx_events_deleted_at ON events(deleted_at);
CREATE INDEX idx_events_time_range ON events(start_time, end_time);
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