# ZCalendar - Data Synchronization

## Overview

ZCalendar implements an **offline-first** architecture where all core functionality works without an internet connection. Users can optionally register accounts to enable cloud synchronization and access premium features.

## Operating Modes

### Guest Mode (Default)
- **Full offline functionality**: Create, edit, delete events and calendars
- **Local storage only**: All data stored on device using Core Data (iOS) or Room (Android)
- **No network dependency**: App works completely offline
- **Privacy focused**: No data leaves the device
- **Instant responsiveness**: All operations are immediate

### Registered User Mode
- **Hybrid storage**: Local storage + cloud backup
- **Background synchronization**: Automatic sync when online
- **Multi-device access**: Data available across all registered devices
- **Conflict resolution**: Automatic handling of conflicting changes
- **Premium features**: Sharing, collaboration, integrations

## Account Registration & Data Migration

### Migration Process

When a user registers an account from guest mode:

1. **Account Creation**
   - User provides email, password, and display name
   - Backend creates new user account with unique ID
   - Server generates JWT tokens for authentication

2. **Data Upload**
   - Client bundles all local events and calendars
   - Data sent to `/auth/register` endpoint with migration payload
   - Server validates and imports all local data
   - Server assigns new server IDs to imported entities

3. **Local Database Update**
   - Client receives mapping of local IDs to server IDs
   - Local database updated with server identifiers
   - Sync metadata initialized for future operations
   - User automatically logged in with new account

4. **Sync Activation**
   - Background sync service activated
   - Periodic sync jobs scheduled
   - Real-time sync for new operations enabled

### Migration Data Format

```json
{
  "events": [
    {
      "local_id": "local_event_123",
      "title": "Team Meeting",
      "start_time": "2024-01-15T10:00:00Z",
      "end_time": "2024-01-15T11:00:00Z",
      "description": "Weekly team sync",
      "calendar_local_id": "local_cal_456"
    }
  ],
  "calendars": [
    {
      "local_id": "local_cal_456", 
      "name": "Work Calendar",
      "color": "#FF5722",
      "description": "Work-related events"
    }
  ]
}
```

## Synchronization Strategy

### Conflict Resolution

When the same event is modified on multiple devices:

1. **Timestamp Comparison**
   - Use `updated_at` timestamp to determine latest version
   - Last write wins for most conflicts

2. **Field-Level Merging**
   - For complex conflicts, merge non-conflicting fields
   - User-initiated changes take priority over automatic updates

3. **Manual Resolution**
   - For critical conflicts, prompt user to choose preferred version
   - Provide diff view showing changes from each device

### Sync Operations

#### Initial Sync (After Registration)
```http
GET /sync/initial
Authorization: Bearer <jwt_token>
```

Downloads complete cloud state for new device.

#### Incremental Sync
```http
GET /sync?since=2024-01-15T10:00:00Z
Authorization: Bearer <jwt_token>
```

Downloads only changes since last sync timestamp.

#### Upload Pending Changes
```http
POST /sync/upload
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "events": {
    "created": [...],
    "updated": [...],
    "deleted": ["event_123"]
  },
  "calendars": {
    "created": [...],
    "updated": [...], 
    "deleted": ["cal_456"]
  }
}
```

### Sync Triggers

1. **App Launch**: Quick sync check when app starts
2. **Background Refresh**: Periodic sync (every 15 minutes when active)
3. **User Action**: Manual pull-to-refresh
4. **Network Change**: Sync when connection restored
5. **Real-time**: Immediate sync for new user actions (when online)

## Offline Capabilities

### Core Features (No Internet Required)

- ✅ Create/edit/delete events and calendars
- ✅ Set reminders and notifications
- ✅ View calendar in all modes (day/week/month/year)
- ✅ Search events and calendars
- ✅ Event recurrence patterns
- ✅ Local notifications
- ✅ Data export to local files

### Online-Only Features

- ❌ **Calendar Sharing**: Share calendars with other users
- ❌ **Event Invitations**: Invite others to events
- ❌ **Real-time Collaboration**: Live collaborative editing
- ❌ **External Integrations**: Google Calendar, Outlook sync
- ❌ **Cloud Backup/Restore**: Access data from new devices
- ❌ **Team Calendars**: Shared team/organization calendars
- ❌ **Analytics**: Usage insights and productivity metrics
- ❌ **Advanced Search**: Search across shared calendars
- ❌ **File Attachments**: Cloud storage for event attachments

## Technical Implementation

### Local Storage Schema

#### Events Table
```sql
CREATE TABLE events (
    id TEXT PRIMARY KEY,
    server_id TEXT,
    title TEXT NOT NULL,
    description TEXT,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    all_day BOOLEAN DEFAULT 0,
    location TEXT,
    calendar_id TEXT NOT NULL,
    sync_status TEXT DEFAULT 'local', -- 'local', 'synced', 'pending'
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    FOREIGN KEY (calendar_id) REFERENCES calendars(id)
);
```

#### Sync Metadata Table
```sql
CREATE TABLE sync_metadata (
    key TEXT PRIMARY KEY,
    value TEXT,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Store last sync timestamps
INSERT INTO sync_metadata VALUES ('last_sync_events', '2024-01-15T10:00:00Z', CURRENT_TIMESTAMP);
INSERT INTO sync_metadata VALUES ('last_sync_calendars', '2024-01-15T10:00:00Z', CURRENT_TIMESTAMP);
```

### Sync Status Management

Each local entity tracks its synchronization status:

- **`local`**: Created offline, not yet uploaded
- **`synced`**: Synchronized with server
- **`pending`**: Modified locally, awaiting upload
- **`conflict`**: Conflicting changes detected

### Background Sync Service

#### iOS Implementation
```swift
import BackgroundTasks

class SyncService {
    func scheduleBackgroundSync() {
        let request = BGAppRefreshTaskRequest(identifier: "com.zcalendar.sync")
        request.earliestBeginDate = Date(timeIntervalSinceNow: 15 * 60) // 15 minutes
        
        try? BGTaskScheduler.shared.submit(request)
    }
    
    func handleBackgroundSync(task: BGAppRefreshTask) {
        task.expirationHandler = {
            task.setTaskCompleted(success: false)
        }
        
        performIncrementalSync { success in
            task.setTaskCompleted(success: success)
            self.scheduleBackgroundSync() // Schedule next sync
        }
    }
}
```

#### Android Implementation
```kotlin
@HiltWorker
class SyncWorker @AssistedInject constructor(
    @Assisted context: Context,
    @Assisted workerParams: WorkerParameters,
    private val syncRepository: SyncRepository
) : CoroutineWorker(context, workerParams) {

    override suspend fun doWork(): Result {
        return try {
            syncRepository.performIncrementalSync()
            Result.success()
        } catch (e: Exception) {
            Result.retry()
        }
    }
}

// Schedule periodic sync
val syncRequest = PeriodicWorkRequestBuilder<SyncWorker>(15, TimeUnit.MINUTES)
    .setConstraints(
        Constraints.Builder()
            .setRequiredNetworkType(NetworkType.CONNECTED)
            .build()
    )
    .build()

WorkManager.getInstance(context).enqueueUniquePeriodicWork(
    "calendar_sync",
    ExistingPeriodicWorkPolicy.KEEP,
    syncRequest
)
```

## Error Handling & Recovery

### Network Failures
- Operations queued locally for retry when connection restored
- User notified of sync status in UI
- Graceful degradation to offline mode

### Server Errors
- Exponential backoff for retry attempts
- Conflict resolution UI for unresolvable conflicts
- Local data preserved as backup

### Data Corruption
- Local database integrity checks
- Automatic repair for minor corruption
- Full re-sync option for major issues

## Security & Privacy

### Guest Mode Security
- All data stored locally with platform encryption
- No network transmissions
- No user tracking or analytics
- Complete privacy protection

### Registered User Security
- End-to-end encryption for sensitive data
- JWT token authentication with refresh tokens
- Secure HTTPS communication
- Server-side data encryption at rest
- Regular security audits

## Performance Optimization

### Local Performance
- Database indexing for common queries
- Efficient Core Data/Room relationships
- Lazy loading for large date ranges
- Memory management for recurring events

### Sync Performance
- Delta sync with minimal data transfer
- Compression for large payloads
- Batch operations for multiple changes
- Intelligent sync scheduling based on usage patterns

## Monitoring & Analytics

### Sync Metrics (Registered Users Only)
- Sync success/failure rates
- Sync duration and data volume
- Conflict frequency and resolution
- Network usage optimization

### Error Tracking
- Sync failure categorization
- Performance bottleneck identification
- User impact assessment
- Automated alerting for critical issues