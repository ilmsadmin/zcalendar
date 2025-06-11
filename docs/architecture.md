# ZCalendar - System Architecture

## Architecture Overview

ZCalendar follows a client-server architecture with native mobile applications communicating with a centralized backend API.

```
┌─────────────────┐    ┌─────────────────┐
│   iOS App       │    │  Android App    │
│   (Swift)       │    │   (Kotlin)      │
└─────────┬───────┘    └─────────┬───────┘
          │                      │
          │      HTTPS/REST      │
          │                      │
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

### API Layer
- **Framework**: Express.js/Node.js
- **Authentication**: JWT tokens with refresh token strategy
- **Rate Limiting**: Implemented to prevent abuse
- **Request Validation**: Input validation and sanitization
- **Error Handling**: Centralized error handling middleware

### Business Logic Layer
- **Services**: Modular business logic services
- **Controllers**: Request/response handling
- **Middleware**: Authentication, logging, validation
- **Data Access Layer**: Repository pattern for database operations

### Database Design
- **Primary Database**: PostgreSQL for structured data
- **Caching**: Redis for session management and caching
- **File Storage**: Cloud storage for attachments (AWS S3/Azure Blob)

### Security
- **Authentication**: JWT-based stateless authentication
- **Authorization**: Role-based access control (RBAC)
- **Data Encryption**: Encryption at rest and in transit
- **API Security**: CORS, HTTPS, rate limiting

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
│             Data Layer               │
│    (Repositories, Core Data)         │
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
│       (Data Management)              │
├──────────────────────────────────────┤
│             Data Layer               │
│      (Room, Retrofit, Local)         │
└──────────────────────────────────────┘
```

## Data Flow

### Event Creation Flow
1. User creates event in mobile app
2. App validates input locally
3. App sends POST request to backend API
4. Backend validates and stores event in database
5. Backend returns success response with event ID
6. App updates local storage and UI

### Synchronization Flow
1. App periodically checks for updates
2. Backend returns modified events since last sync
3. App merges remote changes with local data
4. Conflicts resolved using timestamp-based strategy
5. UI updated to reflect changes

## Design Patterns

### Backend Patterns
- **Repository Pattern**: Data access abstraction
- **Service Layer**: Business logic encapsulation
- **Dependency Injection**: Loose coupling of components
- **Observer Pattern**: Event notifications

### Mobile Patterns
- **MVVM**: Separation of concerns
- **Repository Pattern**: Data layer abstraction
- **Observer Pattern**: UI updates
- **Singleton**: Shared services and managers

## Scalability Considerations

### Horizontal Scaling
- Stateless API design enables load balancing
- Database sharding strategies for large datasets
- CDN for static assets and file storage

### Performance Optimization
- Database indexing for common queries
- API response caching with Redis
- Mobile app offline-first design
- Lazy loading for large datasets

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