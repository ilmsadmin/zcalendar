# ZCalendar API Documentation

## Base URL
```
Production: https://api.zcalendar.com/v1
Development: http://localhost:3000/api/v1
```

## Operating Modes

ZCalendar API supports two operating modes:

### Guest Mode (No Authentication)
- Limited API access for basic functionality
- No persistent server-side storage
- Rate limited to prevent abuse
- Suitable for offline-first usage

### Registered User Mode (Authenticated)
- Full API access with authentication required
- Persistent cloud storage and synchronization
- Higher rate limits
- Advanced features and integrations

## Authentication

### JWT Token Authentication
All API requests require a valid JWT token in the Authorization header:
```
Authorization: Bearer <jwt_token>
```

### Account Registration
```http
POST /auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe",
  "migrate_data": {
    "events": [
      {
        "title": "Local Event",
        "start_time": "2024-01-15T10:00:00Z",
        "end_time": "2024-01-15T11:00:00Z",
        "local_id": "local_event_1"
      }
    ],
    "calendars": [
      {
        "name": "Personal",
        "color": "#FF5722",
        "local_id": "local_cal_1"
      }
    ]
  }
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refreshToken": "refresh_token_here",
    "user": {
      "id": "123",
      "email": "user@example.com",
      "name": "John Doe"
    },
    "migrated_data": {
      "events": [
        {
          "local_id": "local_event_1",
          "server_id": "event_456",
          "status": "migrated"
        }
      ],
      "calendars": [
        {
          "local_id": "local_cal_1", 
          "server_id": "cal_789",
          "status": "migrated"
        }
      ]
    }
  }
}
```
### Login
```http
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refreshToken": "refresh_token_here",
    "user": {
      "id": "123",
      "email": "user@example.com",
      "name": "John Doe"
    }
  }
}
```

## Guest Mode APIs

### Get System Information (No Auth Required)
```http
GET /system/info
```

**Response:**
```json
{
  "success": true,
  "data": {
    "app_version": "1.0.0",
    "api_version": "v1",
    "guest_features": [
      "local_events",
      "local_calendars", 
      "basic_reminders"
    ],
    "premium_features": [
      "sync",
      "sharing",
      "integrations",
      "collaboration",
      "analytics"
    ]
  }
}
```

### Validate Event Data (No Auth Required)
```http
POST /guest/validate/event
Content-Type: application/json

{
  "title": "Meeting",
  "start_time": "2024-01-15T10:00:00Z",
  "end_time": "2024-01-15T11:00:00Z"
}
```

## Events API (Authenticated Users Only)

> **Note**: Event management requires user authentication. Guest users manage events locally on their device.

### Get Events
```http
GET /events?start=2024-01-01&end=2024-01-31
Authorization: Bearer <jwt_token>
```

**Query Parameters:**
- `start` (required): Start date (YYYY-MM-DD)
- `end` (required): End date (YYYY-MM-DD)
- `calendar_id` (optional): Filter by calendar ID

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": "event_123",
      "title": "Team Meeting",
      "description": "Weekly team sync",
      "start_time": "2024-01-15T10:00:00Z",
      "end_time": "2024-01-15T11:00:00Z",
      "all_day": false,
      "location": "Conference Room A",
      "calendar_id": "cal_456",
      "reminders": [
        {
          "type": "notification",
          "minutes_before": 15
        }
      ],
      "recurrence": null,
      "created_at": "2024-01-10T08:00:00Z",
      "updated_at": "2024-01-10T08:00:00Z"
    }
  ]
}
```

### Create Event
```http
POST /events
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "title": "New Meeting",
  "description": "Project discussion",
  "start_time": "2024-01-20T14:00:00Z",
  "end_time": "2024-01-20T15:00:00Z",
  "all_day": false,
  "location": "Room B",
  "calendar_id": "cal_456",
  "reminders": [
    {
      "type": "notification",
      "minutes_before": 10
    }
  ]
}
```

### Update Event
```http
PUT /events/{event_id}
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "title": "Updated Meeting Title",
  "start_time": "2024-01-20T15:00:00Z",
  "end_time": "2024-01-20T16:00:00Z"
}
```

### Delete Event
```http
DELETE /events/{event_id}
Authorization: Bearer <jwt_token>
```

## Calendars API (Authenticated Users Only)

> **Note**: Calendar management requires user authentication. Guest users manage calendars locally on their device.

### Get User Calendars
```http
GET /calendars
Authorization: Bearer <jwt_token>
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": "cal_456",
      "name": "Work Calendar",
      "description": "Work-related events",
      "color": "#FF5722",
      "is_primary": true,
      "is_shared": false,
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

### Create Calendar
```http
POST /calendars
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "name": "Personal Calendar",
  "description": "Personal events",
  "color": "#2196F3"
}
```

## Sharing & Collaboration API (Premium Features)

### Share Calendar
```http
POST /calendars/{calendar_id}/share
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "email": "colleague@example.com",
  "permission": "read",
  "message": "Sharing my work calendar with you"
}
```

### Get Shared Calendars
```http
GET /calendars/shared
Authorization: Bearer <jwt_token>
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": "cal_shared_123",
      "name": "Team Calendar",
      "owner": {
        "name": "John Doe",
        "email": "john@example.com"
      },
      "permission": "write",
      "shared_at": "2024-01-10T00:00:00Z"
    }
  ]
}
```

### Invite to Event
```http
POST /events/{event_id}/invite
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "emails": ["user1@example.com", "user2@example.com"],
  "message": "Please join our team meeting"
}
```

## Integration API (Premium Features)

### Import from External Calendar
```http
POST /integrations/import
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "provider": "google",
  "access_token": "external_service_token",
  "calendar_ids": ["primary", "work"]
}
```

### Setup Calendar Sync
```http
POST /integrations/sync
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "provider": "google",
  "sync_direction": "bidirectional",
  "sync_frequency": "realtime"
}
```

## Sync API (Authenticated Users Only)

### Get Changes Since Last Sync
```http
GET /sync?since=2024-01-15T10:00:00Z
Authorization: Bearer <jwt_token>
```

**Response:**
```json
{
  "success": true,
  "data": {
    "events": {
      "created": [...],
      "updated": [...],
      "deleted": ["event_123", "event_456"]
    },
    "calendars": {
      "created": [...],
      "updated": [...],
      "deleted": ["cal_789"]
    },
    "last_sync": "2024-01-15T12:00:00Z"
  }
}
```

## Error Responses

### Standard Error Format
```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Validation failed",
    "details": [
      {
        "field": "title",
        "message": "Title is required"
      }
    ]
  }
}
```

### Common Error Codes
- `UNAUTHORIZED` (401): Invalid or missing authentication
- `FORBIDDEN` (403): Insufficient permissions
- `NOT_FOUND` (404): Resource not found
- `VALIDATION_ERROR` (400): Request validation failed
- `RATE_LIMIT_EXCEEDED` (429): Too many requests
- `INTERNAL_ERROR` (500): Server error

## Rate Limiting

- **Guest users**: 100 requests per hour (validation and system info only)
- **Authenticated users**: 1000 requests per hour (full API access)
- **Premium users**: 5000 requests per hour (with integrations and sharing)
- **Rate limit headers** included in responses:
  ```
  X-RateLimit-Limit: 1000
  X-RateLimit-Remaining: 999
  X-RateLimit-Reset: 1642684800
  X-RateLimit-Tier: authenticated
  ```

## Pagination

For endpoints returning lists, pagination is supported:
```http
GET /events?page=1&limit=50
```

**Response includes pagination metadata:**
```json
{
  "success": true,
  "data": [...],
  "pagination": {
    "page": 1,
    "limit": 50,
    "total": 200,
    "pages": 4
  }
}
```