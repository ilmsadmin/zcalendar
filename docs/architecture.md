# ZCalendar - System Architecture

## Architecture Overview

ZCalendar follows an **offline-first architecture** with optional cloud synchronization for registered users. The application operates in two modes:

### Guest Mode (Offline-First)
```
┌─────────────────┐    ┌─────────────────┐
│   iOS App       │    │  Android App    │
│   (Swift)       │    │   (Kotlin)      │
│                 │    │                 │
│  ┌───────────┐  │    │  ┌───────────┐  │
│  │Local Store│  │    │  │Local Store│  │
│  │(Core Data)│  │    │  │  (Room)   │  │
│  └───────────┘  │    │  └───────────┘  │
└─────────────────┘    └─────────────────┘
```

### Registered User Mode (Hybrid)
```
┌─────────────────┐    ┌─────────────────┐
│   iOS App       │    │  Android App    │
│   (Swift)       │    │   (Kotlin)      │
│                 │    │                 │
│  ┌───────────┐  │    │  ┌───────────┐  │
│  │Local Store│  │    │  │Local Store│  │
│  │(Core Data)│  │    │  │  (Room)   │  │
│  └─────┬─────┘  │    │  └─────┬─────┘  │
└────────┼────────┘    └────────┼────────┘
         │                      │
         │      HTTPS/REST      │
         │    (Background Sync) │
   ┌─────┴──────────────────────┴─────┐
   │         Backend API              │
   │       (Node.js/Express)          │
   └─────────────┬───────────────────┘
                 │
         ┌───────┴────────┐
         │   Database     │
         │ (PostgreSQL)   │
         └────────────────┘
```

## Backend Architecture

### Authentication & User Management
- **Guest Mode**: No authentication required, anonymous usage
- **Account Registration**: Optional user registration with email/password
- **Data Migration**: Automated transfer of offline data to cloud upon registration
- **JWT Tokens**: Stateless authentication for registered users
- **Refresh Tokens**: Long-term session management

### API Layer
- **Framework**: Express.js/Node.js
- **Dual Mode Support**: Both anonymous and authenticated endpoints
- **Rate Limiting**: Different limits for guest vs registered users
- **Request Validation**: Input validation and sanitization
- **Error Handling**: Centralized error handling middleware

### Business Logic Layer
- **Services**: Modular business logic services
- **Controllers**: Request/response handling
- **Middleware**: Authentication, logging, validation
- **Data Access Layer**: Repository pattern for database operations

### Data Synchronization
- **Primary Database**: PostgreSQL for registered user data
- **Local-First**: All operations work offline with local storage
- **Background Sync**: Automatic synchronization when online
- **Conflict Resolution**: Timestamp-based and last-write-wins strategies
- **Caching**: Redis for session management and sync optimization
- **File Storage**: Cloud storage for attachments (AWS S3/Azure Blob)

### Security
- **Guest Mode Privacy**: No data leaves the device in guest mode
- **Optional Authentication**: JWT-based stateless authentication for registered users
- **Authorization**: Role-based access control (RBAC) for registered features
- **Data Encryption**: Encryption at rest and in transit for cloud data
- **API Security**: CORS, HTTPS, rate limiting with different tiers

## Mobile Architecture

### iOS Application Architecture
```
┌──────────────────────────────────────┐
│              UI Layer                │
│    (Views, ViewControllers)          │
├──────────────────────────────────────┤
│           Presentation Layer         │
│         (ViewModels, MVVM)           │
├──────────────────────────────────────┤
│            Business Layer            │
│      (Use Cases, Services)           │
├──────────────────────────────────────┤
│           Sync Management            │
│     (Offline/Online Coordinator)     │
├──────────────────────────────────────┤
│             Data Layer               │
│  (Local: Core Data, Remote: API)     │
└──────────────────────────────────────┘
```

### Android Application Architecture
```
┌──────────────────────────────────────┐
│              UI Layer                │
│    (Activities, Fragments)           │
├──────────────────────────────────────┤
│           ViewModel Layer            │
│         (MVVM, LiveData)             │
├──────────────────────────────────────┤
│            Repository Layer          │
│    (Offline-First Data Management)   │
├──────────────────────────────────────┤
│           Sync Coordinator           │
│     (Background Sync Service)        │
├──────────────────────────────────────┤
│             Data Layer               │
│   (Local: Room, Remote: Retrofit)    │
└──────────────────────────────────────┘
```

## Data Flow

### Guest Mode Event Creation
1. User creates event in mobile app
2. App validates input locally
3. Event stored immediately in local database
4. UI updated instantly
5. No network requests required

### Registered User Event Creation
1. User creates event in mobile app
2. App validates input locally
3. Event stored immediately in local database
4. UI updated instantly
5. Background sync queues the event for upload
6. When online, app sends POST request to backend API
7. Backend validates and stores event in cloud database
8. Backend returns success response with server event ID
9. App updates local event with server ID for future sync

### Account Registration & Data Migration
1. User chooses to create account from guest mode
2. App collects registration information (email, password)
3. App sends registration request with local data payload
4. Backend creates user account and processes data migration
5. Server validates and imports all local events/calendars
6. Server returns authentication tokens and migrated data IDs
7. App updates local storage with server IDs
8. Future operations use hybrid local/cloud approach

### Synchronization Flow (Registered Users)
1. App periodically checks for updates when online
2. App sends last sync timestamp to server
3. Backend returns all changes since last sync
4. App processes incremental updates:
   - New events from other devices
   - Updated events with conflict resolution
   - Deleted events marked for removal
5. Local database updated with remote changes
6. UI refreshed to reflect synchronized state

## Design Patterns

### Backend Patterns
- **Repository Pattern**: Data access abstraction
- **Service Layer**: Business logic encapsulation
- **Dependency Injection**: Loose coupling of components
- **Observer Pattern**: Event notifications

### Mobile Patterns
- **Offline-First Repository**: Data layer abstracts local/remote storage
- **MVVM**: Separation of concerns with reactive UI updates
- **Sync Coordinator**: Manages background synchronization
- **Observer Pattern**: UI updates for data changes
- **Command Pattern**: Queuing operations for background sync
- **Singleton**: Shared services and sync managers

## Scalability Considerations

### Horizontal Scaling
- Stateless API design enables load balancing
- Database sharding strategies for large datasets
- CDN for static assets and file storage

### Performance Optimization
- **Database indexing** for common queries (both local and cloud)
- **Intelligent sync** with delta updates and compression
- **Offline-first design** ensures instant UI responsiveness  
- **Background sync** prevents blocking user interactions
- **Lazy loading** for large datasets and historical data
- **Conflict resolution** algorithms for seamless multi-device experience

## Security Architecture

### Data Protection
- End-to-end encryption for sensitive data
- Regular security audits and penetration testing
- Secure credential storage on mobile devices
- API rate limiting and DDoS protection

### Privacy Compliance
- GDPR compliance for European users
- Data retention policies
- User consent management
- Right to be forgotten implementation