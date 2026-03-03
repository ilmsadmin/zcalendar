# 📅 Zplus Calendar - Hệ Thống Lịch Âm Dương Việt Nam

## Phiên bản: 1.0.0
## Ngày tạo: 03/03/2026
## Tác giả: Zplus Team

---

## 📋 Mục Lục

1. [Tổng Quan Dự Án](#1-tổng-quan-dự-án)
2. [Kiến Trúc Hệ Thống](#2-kiến-trúc-hệ-thống)
3. [Tech Stack](#3-tech-stack)
4. [Database Design](#4-database-design)
5. [API Design](#5-api-design)
6. [Frontend Design](#6-frontend-design)
7. [Thuật Toán Chuyển Đổi Âm Dương Lịch](#7-thuật-toán-chuyển-đổi-âm-dương-lịch)
8. [Chức Năng Note & Task](#8-chức-năng-note--task)
9. [Caching Strategy](#9-caching-strategy)
10. [Authentication & Authorization](#10-authentication--authorization)
11. [Deployment](#11-deployment)
12. [Đề Xuất Mở Rộng](#12-đề-xuất-mở-rộng)

---

## 1. Tổng Quan Dự Án

### 1.1 Mô Tả
Zplus Calendar là ứng dụng lịch thông minh hỗ trợ hiển thị **Dương lịch (Solar Calendar)** và **Âm lịch (Lunar Calendar)** theo chuẩn lịch Việt Nam. Ứng dụng tích hợp chức năng ghi chú (Note) và quản lý công việc (Task), phục vụ người dùng Việt Nam trong việc tra cứu ngày lễ, ngày tốt/xấu, và quản lý công việc hàng ngày.

### 1.2 Mục Tiêu
- Hiển thị lịch Dương lịch và Âm lịch chính xác
- Hỗ trợ các ngày lễ Việt Nam (cả Dương lịch & Âm lịch)
- Tính toán ngày tốt/xấu, giờ hoàng đạo
- Quản lý ghi chú và công việc theo ngày
- Giao diện đẹp, dễ sử dụng, responsive
- Hiệu năng cao, thời gian phản hồi < 200ms

### 1.3 Đối Tượng Sử Dụng
- Người dùng Việt Nam cần tra cứu âm lịch
- Người dùng cần quản lý công việc và ghi chú hàng ngày
- Doanh nghiệp cần lịch phù hợp với văn hóa Việt Nam

### 1.4 Phạm Vi Chức Năng (Scope)

| Module | Mô Tả | Ưu Tiên |
|--------|--------|---------|
| Calendar View | Hiển thị lịch tháng/tuần/ngày với Âm - Dương lịch | P0 |
| Lunar Conversion | Chuyển đổi Dương ↔ Âm lịch | P0 |
| Vietnamese Holidays | Ngày lễ, ngày nghỉ Việt Nam | P0 |
| Notes | Ghi chú theo ngày | P0 |
| Tasks | Quản lý công việc | P0 |
| Good/Bad Days | Ngày tốt/xấu, giờ hoàng đạo | P1 |
| Notifications | Nhắc nhở task/event | P1 |
| User Auth | Đăng ký, đăng nhập | P0 |
| Sharing | Chia sẻ lịch, task | P2 |
| Calendar Sync | Đồng bộ Google Calendar | P2 |

---

## 2. Kiến Trúc Hệ Thống

### 2.1 Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                        CLIENT LAYER                              │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │              Next.js 14 (App Router)                     │    │
│  │  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐  │    │
│  │  │ Calendar  │ │  Notes   │ │  Tasks   │ │ Settings │  │    │
│  │  │  Module   │ │  Module  │ │  Module  │ │  Module  │  │    │
│  │  └──────────┘ └──────────┘ └──────────┘ └──────────┘  │    │
│  │  ┌─────────────────────────────────────────────────┐    │    │
│  │  │         State Management (Zustand)               │    │    │
│  │  └─────────────────────────────────────────────────┘    │    │
│  └─────────────────────────────────────────────────────────┘    │
└────────────────────────────┬────────────────────────────────────┘
                             │ HTTPS/REST API + WebSocket
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                     API GATEWAY / LOAD BALANCER                  │
│                        (Nginx / Traefik)                         │
└────────────────────────────┬────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                      BACKEND LAYER                               │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │              Go Fiber v2 API Server                      │    │
│  │                                                          │    │
│  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  │    │
│  │  │   Calendar    │  │    Notes     │  │    Tasks     │  │    │
│  │  │   Handler     │  │   Handler    │  │   Handler    │  │    │
│  │  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘  │    │
│  │         │                  │                  │          │    │
│  │  ┌──────▼───────┐  ┌──────▼───────┐  ┌──────▼───────┐  │    │
│  │  │   Calendar    │  │    Notes     │  │    Tasks     │  │    │
│  │  │   Service     │  │   Service    │  │   Service    │  │    │
│  │  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘  │    │
│  │         │                  │                  │          │    │
│  │  ┌──────▼───────┐  ┌──────▼───────┐  ┌──────▼───────┐  │    │
│  │  │   Calendar    │  │    Notes     │  │    Tasks     │  │    │
│  │  │  Repository   │  │  Repository  │  │  Repository  │  │    │
│  │  └──────────────┘  └──────────────┘  └──────────────┘  │    │
│  │                                                          │    │
│  │  ┌──────────────────────────────────────────────────┐   │    │
│  │  │          Lunar Calendar Engine (Core)              │   │    │
│  │  │  - Solar ↔ Lunar Conversion                       │   │    │
│  │  │  - Can Chi Calculation                            │   │    │
│  │  │  - Good/Bad Day Calculator                        │   │    │
│  │  │  - Holiday Database                               │   │    │
│  │  └──────────────────────────────────────────────────┘   │    │
│  └─────────────────────────────────────────────────────────┘    │
└──────────┬──────────────────┬──────────────────┬────────────────┘
           │                  │                  │
           ▼                  ▼                  ▼
┌──────────────────┐ ┌───────────────┐ ┌─────────────────┐
│   PostgreSQL     │ │   MongoDB     │ │     Redis       │
│                  │ │               │ │                 │
│ - Users          │ │ - Notes       │ │ - Session       │
│ - Tasks          │ │ - Calendar    │ │ - Cache lunar   │
│ - Categories     │ │   Metadata    │ │ - Rate limiting │
│ - Tags           │ │ - Activity    │ │ - Real-time     │
│ - Reminders      │ │   Logs        │ │   pub/sub       │
└──────────────────┘ └───────────────┘ └─────────────────┘
```

### 2.2 Design Pattern
- **Clean Architecture** (Handler → Service → Repository)
- **Repository Pattern** cho Data Access
- **Dependency Injection** cho quản lý dependencies
- **Observer Pattern** cho notifications
- **Strategy Pattern** cho lunar calculation algorithms

---

## 3. Tech Stack

### 3.1 Backend

| Technology | Version | Mục Đích |
|-----------|---------|----------|
| **Go** | 1.22+ | Ngôn ngữ lập trình chính |
| **Go Fiber** | v2.52+ | Web framework (hiệu năng cao) |
| **GORM** | v2 | ORM cho PostgreSQL |
| **MongoDB Go Driver** | v2 | Driver cho MongoDB |
| **go-redis** | v9 | Redis client |
| **JWT-Go** | v5 | Authentication |
| **Validator** | v10 | Request validation |
| **Zap** | v1.27 | Structured logging |
| **Viper** | v1.18 | Configuration management |
| **Swagger** | v2 | API documentation |
| **Air** | v1.52 | Hot reload cho development |

### 3.2 Frontend

| Technology | Version | Mục Đích |
|-----------|---------|----------|
| **Next.js** | 14+ | React framework (App Router) |
| **TypeScript** | 5.x | Type safety |
| **Tailwind CSS** | 3.x | Styling |
| **Shadcn/UI** | latest | UI Component library |
| **Zustand** | 4.x | State management |
| **TanStack Query** | v5 | Data fetching & caching |
| **Day.js** | 1.x | Date manipulation |
| **Framer Motion** | 11.x | Animations |
| **React DnD** | 16.x | Drag & Drop cho tasks |
| **Lucide React** | latest | Icons |

### 3.3 Database

| Database | Mục Đích | Lý Do Chọn |
|----------|----------|-------------|
| **PostgreSQL 16** | Users, Tasks, Categories, Reminders | Dữ liệu quan hệ, ACID compliance, complex queries |
| **MongoDB 7** | Notes, Calendar metadata, Activity logs | Flexible schema, rich text, document-oriented |
| **Redis 7** | Cache, Sessions, Rate limiting | In-memory, cực nhanh, TTL support |

### 3.4 Infrastructure

| Technology | Mục Đích |
|-----------|----------|
| **Docker** | Containerization |
| **Docker Compose** | Local development |
| **Nginx** | Reverse proxy, Load balancing |
| **GitHub Actions** | CI/CD |
| **Prometheus** | Metrics collection |
| **Grafana** | Monitoring dashboard |

---

## 4. Database Design

### 4.1 PostgreSQL Schema

```sql
-- ============================================
-- DATABASE: zplus_calendar
-- ============================================

-- Extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- ============================================
-- TABLE: users
-- ============================================
CREATE TABLE users (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email           VARCHAR(255) UNIQUE NOT NULL,
    username        VARCHAR(100) UNIQUE NOT NULL,
    password_hash   VARCHAR(255) NOT NULL,
    full_name       VARCHAR(255),
    avatar_url      VARCHAR(500),
    timezone        VARCHAR(50) DEFAULT 'Asia/Ho_Chi_Minh',
    locale          VARCHAR(10) DEFAULT 'vi-VN',
    
    -- Preferences
    default_calendar_view   VARCHAR(20) DEFAULT 'month',  -- month, week, day
    show_lunar              BOOLEAN DEFAULT true,
    show_good_bad_days      BOOLEAN DEFAULT true,
    theme                   VARCHAR(20) DEFAULT 'light',   -- light, dark, auto
    
    -- OAuth
    google_id       VARCHAR(255),
    facebook_id     VARCHAR(255),
    
    -- Status
    is_active       BOOLEAN DEFAULT true,
    is_verified     BOOLEAN DEFAULT false,
    last_login_at   TIMESTAMP WITH TIME ZONE,
    
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at      TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);

-- ============================================
-- TABLE: categories (cho Tasks & Notes)
-- ============================================
CREATE TABLE categories (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name        VARCHAR(100) NOT NULL,
    color       VARCHAR(7) DEFAULT '#3B82F6',   -- Hex color
    icon        VARCHAR(50),                      -- Lucide icon name
    type        VARCHAR(20) NOT NULL,             -- 'task', 'note', 'both'
    sort_order  INTEGER DEFAULT 0,
    
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(user_id, name, type)
);

CREATE INDEX idx_categories_user ON categories(user_id);

-- ============================================
-- TABLE: tags
-- ============================================
CREATE TABLE tags (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name        VARCHAR(50) NOT NULL,
    color       VARCHAR(7) DEFAULT '#6B7280',
    
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(user_id, name)
);

CREATE INDEX idx_tags_user ON tags(user_id);

-- ============================================
-- TABLE: tasks
-- ============================================
CREATE TABLE tasks (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category_id     UUID REFERENCES categories(id) ON DELETE SET NULL,
    parent_task_id  UUID REFERENCES tasks(id) ON DELETE CASCADE,  -- Subtask support
    
    title           VARCHAR(500) NOT NULL,
    description     TEXT,
    
    -- Scheduling
    due_date        DATE,
    due_time        TIME,
    start_date      DATE,
    start_time      TIME,
    
    -- Lunar date reference (for recurring lunar events)
    lunar_day       INTEGER,          -- 1-30
    lunar_month     INTEGER,          -- 1-12
    lunar_year      INTEGER,
    is_lunar_based  BOOLEAN DEFAULT false,
    
    -- Status & Priority
    status          VARCHAR(20) DEFAULT 'todo',     -- todo, in_progress, done, cancelled
    priority        VARCHAR(10) DEFAULT 'medium',   -- low, medium, high, urgent
    
    -- Recurrence
    is_recurring        BOOLEAN DEFAULT false,
    recurrence_rule     VARCHAR(255),   -- RRULE format (RFC 5545)
    recurrence_end_date DATE,
    
    -- Progress
    progress        INTEGER DEFAULT 0 CHECK (progress >= 0 AND progress <= 100),
    completed_at    TIMESTAMP WITH TIME ZONE,
    
    -- Meta
    sort_order      INTEGER DEFAULT 0,
    estimated_hours DECIMAL(5,2),
    actual_hours    DECIMAL(5,2),
    
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at      TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_tasks_user ON tasks(user_id);
CREATE INDEX idx_tasks_due_date ON tasks(due_date);
CREATE INDEX idx_tasks_status ON tasks(status);
CREATE INDEX idx_tasks_category ON tasks(category_id);
CREATE INDEX idx_tasks_parent ON tasks(parent_task_id);
CREATE INDEX idx_tasks_user_date ON tasks(user_id, due_date);

-- ============================================
-- TABLE: task_tags (Many-to-Many)
-- ============================================
CREATE TABLE task_tags (
    task_id     UUID NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    tag_id      UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (task_id, tag_id)
);

-- ============================================
-- TABLE: reminders
-- ============================================
CREATE TABLE reminders (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    task_id         UUID REFERENCES tasks(id) ON DELETE CASCADE,
    
    title           VARCHAR(255) NOT NULL,
    message         TEXT,
    remind_at       TIMESTAMP WITH TIME ZONE NOT NULL,
    
    -- Reminder type
    type            VARCHAR(20) DEFAULT 'push',  -- push, email, both
    is_sent         BOOLEAN DEFAULT false,
    sent_at         TIMESTAMP WITH TIME ZONE,
    
    -- Recurrence  
    is_recurring    BOOLEAN DEFAULT false,
    recurrence_rule VARCHAR(255),
    
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_reminders_user ON reminders(user_id);
CREATE INDEX idx_reminders_remind_at ON reminders(remind_at);
CREATE INDEX idx_reminders_unsent ON reminders(is_sent, remind_at);

-- ============================================
-- TABLE: vietnamese_holidays
-- ============================================
CREATE TABLE vietnamese_holidays (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name            VARCHAR(255) NOT NULL,
    name_en         VARCHAR(255),
    description     TEXT,
    
    -- Date
    day             INTEGER NOT NULL,
    month           INTEGER NOT NULL,
    is_lunar        BOOLEAN DEFAULT false,      -- true = Âm lịch, false = Dương lịch
    
    -- Type
    type            VARCHAR(30) NOT NULL,        -- public_holiday, cultural, memorial, traditional
    is_day_off      BOOLEAN DEFAULT false,       -- Có được nghỉ không
    days_off_count  INTEGER DEFAULT 0,           -- Số ngày nghỉ
    
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_holidays_date ON vietnamese_holidays(month, day);
CREATE INDEX idx_holidays_lunar ON vietnamese_holidays(is_lunar);

-- ============================================
-- TABLE: user_sessions (backup cho Redis)
-- ============================================
CREATE TABLE user_sessions (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    refresh_token   VARCHAR(500) NOT NULL,
    device_info     JSONB,
    ip_address      VARCHAR(45),
    expires_at      TIMESTAMP WITH TIME ZONE NOT NULL,
    
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_sessions_user ON user_sessions(user_id);
CREATE INDEX idx_sessions_token ON user_sessions(refresh_token);
CREATE INDEX idx_sessions_expires ON user_sessions(expires_at);
```

### 4.2 MongoDB Schema

```javascript
// ============================================
// COLLECTION: notes
// ============================================
{
    _id: ObjectId,
    user_id: "uuid-string",                    // Reference to PostgreSQL users
    
    // Content
    title: "Ghi chú ngày rằm",
    content: {
        // Rich text content (Tiptap/ProseMirror format)
        type: "doc",
        content: [
            {
                type: "heading",
                attrs: { level: 2 },
                content: [{ type: "text", text: "Chuẩn bị lễ rằm tháng 7" }]
            },
            {
                type: "bulletList",
                content: [
                    {
                        type: "listItem",
                        content: [
                            {
                                type: "paragraph",
                                content: [{ type: "text", text: "Mua hoa quả" }]
                            }
                        ]
                    }
                ]
            },
            {
                type: "paragraph",
                content: [
                    { type: "text", text: "Ghi nhớ: " },
                    { type: "text", marks: [{ type: "bold" }], text: "cúng lúc 12h trưa" }
                ]
            }
        ]
    },
    plain_text: "Chuẩn bị lễ rằm tháng 7\nMua hoa quả\nGhi nhớ: cúng lúc 12h trưa",
    
    // Date association
    date: ISODate("2026-03-03"),               // Dương lịch
    lunar_date: {                               // Âm lịch
        day: 15,
        month: 1,
        year: 2026,
        leap: false                             // Tháng nhuận
    },
    
    // Organization
    category_id: "uuid-string",
    tags: ["uuid-string-1", "uuid-string-2"],
    color: "#FEF3C7",                           // Note background color
    is_pinned: false,
    is_archived: false,
    
    // Attachments
    attachments: [
        {
            id: "uuid-string",
            file_name: "photo.jpg",
            file_url: "/uploads/2026/03/photo.jpg",
            file_size: 1024000,                 // bytes
            mime_type: "image/jpeg",
            uploaded_at: ISODate("2026-03-03T10:00:00Z")
        }
    ],
    
    // Checklist (embedded)
    checklist: [
        { id: "uuid", text: "Mua hoa", is_checked: true, sort_order: 0 },
        { id: "uuid", text: "Mua quả", is_checked: false, sort_order: 1 }
    ],
    
    // Sharing
    shared_with: [
        {
            user_id: "uuid-string",
            permission: "read"                  // read, write
        }
    ],
    
    // Meta
    word_count: 42,
    created_at: ISODate("2026-03-03T08:00:00Z"),
    updated_at: ISODate("2026-03-03T10:30:00Z"),
    deleted_at: null
}

// Indexes cho notes collection
db.notes.createIndex({ user_id: 1, date: -1 });
db.notes.createIndex({ user_id: 1, "lunar_date.month": 1, "lunar_date.day": 1 });
db.notes.createIndex({ user_id: 1, is_pinned: -1, updated_at: -1 });
db.notes.createIndex({ user_id: 1, category_id: 1 });
db.notes.createIndex({ user_id: 1, tags: 1 });
db.notes.createIndex({ plain_text: "text", title: "text" });  // Full-text search
db.notes.createIndex({ deleted_at: 1 }, { expireAfterSeconds: 2592000 }); // TTL 30 days

// ============================================
// COLLECTION: calendar_events
// ============================================
{
    _id: ObjectId,
    user_id: "uuid-string",
    
    title: "Giỗ Ông Nội",
    description: "Giỗ ông nội năm thứ 5",
    
    // Time
    start_date: ISODate("2026-03-15"),
    end_date: ISODate("2026-03-15"),
    all_day: true,
    
    // Lunar reference
    is_lunar_event: true,
    lunar_date: {
        day: 20,
        month: 2,
        year: 2026,
        leap: false
    },
    
    // Recurrence (yearly lunar)
    recurrence: {
        type: "yearly_lunar",                   // none, daily, weekly, monthly, yearly, yearly_lunar
        interval: 1,
        end_date: null
    },
    
    // Display
    color: "#EF4444",
    icon: "🙏",
    
    // Reminder
    reminders: [
        { type: "notification", minutes_before: 1440 },  // 1 day before
        { type: "notification", minutes_before: 60 }      // 1 hour before
    ],
    
    // Meta
    event_type: "family",                       // family, work, personal, health, education
    location: "Nhà thờ họ",
    
    created_at: ISODate("2026-03-01T08:00:00Z"),
    updated_at: ISODate("2026-03-01T08:00:00Z"),
    deleted_at: null
}

// ============================================
// COLLECTION: activity_logs
// ============================================
{
    _id: ObjectId,
    user_id: "uuid-string",
    
    action: "task.created",                     // task.created, note.updated, etc.
    entity_type: "task",
    entity_id: "uuid-string",
    
    changes: {
        before: { status: "todo" },
        after: { status: "done" }
    },
    
    metadata: {
        ip_address: "192.168.1.1",
        user_agent: "Mozilla/5.0...",
        device: "desktop"
    },
    
    created_at: ISODate("2026-03-03T10:00:00Z")
}

// TTL index - auto delete after 90 days
db.activity_logs.createIndex({ created_at: 1 }, { expireAfterSeconds: 7776000 });
db.activity_logs.createIndex({ user_id: 1, created_at: -1 });
db.activity_logs.createIndex({ entity_type: 1, entity_id: 1 });
```

### 4.3 Redis Data Structure

```
# ============================================
# SESSION MANAGEMENT
# ============================================
# User session
SET     session:{session_id}        {user_id, device_info, created_at}
EXPIRE  session:{session_id}        86400       # 24 hours

# Refresh token
SET     refresh:{user_id}:{device}  {refresh_token}
EXPIRE  refresh:{user_id}:{device}  604800      # 7 days

# ============================================
# CALENDAR CACHE
# ============================================
# Lunar calendar data for a month (pre-computed)
HSET    lunar:2026:3    "1" "{lunar_day:13, lunar_month:1, can_chi:'Giáp Tý', ...}"
                        "2" "{lunar_day:14, lunar_month:1, can_chi:'Ất Sửu', ...}"
                        ...
                        "31" "{...}"
EXPIRE  lunar:2026:3    2592000     # 30 days

# Holidays for a year
SET     holidays:2026               "[{name, date, type, ...}]"
EXPIRE  holidays:2026               31536000    # 1 year

# Good/Bad days for a month
SET     good_days:2026:3            "[{date, quality, hours, ...}]"
EXPIRE  good_days:2026:3            2592000     # 30 days

# ============================================
# USER DATA CACHE
# ============================================
# User profile
SET     user:{user_id}:profile      {serialized_user_data}
EXPIRE  user:{user_id}:profile      3600        # 1 hour

# User tasks for a date
SET     user:{user_id}:tasks:2026-03-03    [{task1}, {task2}]
EXPIRE  user:{user_id}:tasks:2026-03-03    300         # 5 minutes

# User notes for a date
SET     user:{user_id}:notes:2026-03-03    [{note1}, {note2}]
EXPIRE  user:{user_id}:notes:2026-03-03    300         # 5 minutes

# ============================================
# RATE LIMITING
# ============================================
# API rate limit per user
SET     ratelimit:{user_id}:{endpoint}    {count}
EXPIRE  ratelimit:{user_id}:{endpoint}    60    # 1 minute window

# ============================================
# REAL-TIME
# ============================================
# Pub/Sub channels for live updates
PUBLISH calendar:{user_id}    {event_type, data}
```

---

## 5. API Design

### 5.1 Base URL & Versioning
```
Base URL:       https://api.zpluscalendar.com
API Version:    /api/v1
```

### 5.2 Authentication Endpoints

```
POST    /api/v1/auth/register           # Đăng ký
POST    /api/v1/auth/login              # Đăng nhập
POST    /api/v1/auth/logout             # Đăng xuất
POST    /api/v1/auth/refresh            # Refresh token
POST    /api/v1/auth/forgot-password    # Quên mật khẩu
POST    /api/v1/auth/reset-password     # Đặt lại mật khẩu
POST    /api/v1/auth/verify-email       # Xác thực email
POST    /api/v1/auth/google             # Login with Google
POST    /api/v1/auth/facebook           # Login with Facebook
```

### 5.3 Calendar Endpoints

```
GET     /api/v1/calendar/month/:year/:month         # Lấy dữ liệu lịch tháng
GET     /api/v1/calendar/week/:year/:week            # Lấy dữ liệu lịch tuần
GET     /api/v1/calendar/day/:year/:month/:day       # Lấy dữ liệu lịch ngày

GET     /api/v1/calendar/lunar/:year/:month/:day     # Chuyển đổi Dương → Âm
GET     /api/v1/calendar/solar/:year/:month/:day     # Chuyển đổi Âm → Dương

GET     /api/v1/calendar/holidays/:year              # Danh sách ngày lễ năm
GET     /api/v1/calendar/good-days/:year/:month      # Ngày tốt/xấu trong tháng
GET     /api/v1/calendar/today                       # Thông tin ngày hôm nay
```

### 5.4 Calendar Response Example

```json
// GET /api/v1/calendar/day/2026/3/3
{
    "status": "success",
    "data": {
        "solar": {
            "date": "2026-03-03",
            "day": 3,
            "month": 3,
            "year": 2026,
            "day_of_week": 2,                   // Thứ Ba
            "day_of_week_name": "Thứ Ba",
            "week_of_year": 10
        },
        "lunar": {
            "day": 15,
            "month": 1,
            "year": 2026,
            "leap": false,
            "day_name": "Rằm",
            "month_name": "Tháng Giêng"
        },
        "can_chi": {
            "day": "Bính Thân",
            "month": "Canh Dần",
            "year": "Bính Ngọ",
            "hour": "Mậu Tý"
        },
        "good_bad": {
            "quality": "good",                  // good, neutral, bad
            "quality_name": "Ngày Tốt",
            "good_for": [
                "Cưới hỏi",
                "Khai trương",
                "Xuất hành"
            ],
            "bad_for": [
                "Phá thổ",
                "An táng"
            ],
            "good_hours": [
                { "hour": "Tý", "time": "23:00-01:00", "quality": "hoàng đạo" },
                { "hour": "Sửu", "time": "01:00-03:00", "quality": "hắc đạo" },
                { "hour": "Dần", "time": "03:00-05:00", "quality": "hoàng đạo" }
            ]
        },
        "holidays": [
            {
                "name": "Tết Nguyên tiêu",
                "type": "traditional",
                "is_lunar": true,
                "description": "Rằm tháng Giêng - Tết Nguyên tiêu"
            }
        ],
        "solar_term": {
            "name": "Kinh trập",
            "name_en": "Jingzhe",
            "description": "Sâu bọ bắt đầu hoạt động"
        }
    }
}
```

### 5.5 Task Endpoints

```
# CRUD
POST    /api/v1/tasks                           # Tạo task mới
GET     /api/v1/tasks                           # Danh sách tasks (with filters)
GET     /api/v1/tasks/:id                       # Chi tiết task
PUT     /api/v1/tasks/:id                       # Cập nhật task
DELETE  /api/v1/tasks/:id                       # Xóa task (soft delete)

# Actions
PATCH   /api/v1/tasks/:id/status                # Cập nhật trạng thái
PATCH   /api/v1/tasks/:id/priority              # Cập nhật độ ưu tiên
POST    /api/v1/tasks/:id/subtasks              # Tạo subtask
GET     /api/v1/tasks/:id/subtasks              # Lấy subtasks

# Queries
GET     /api/v1/tasks/date/:date                # Tasks theo ngày
GET     /api/v1/tasks/range?start=&end=         # Tasks theo khoảng thời gian
GET     /api/v1/tasks/overdue                   # Tasks quá hạn
GET     /api/v1/tasks/stats                     # Thống kê tasks
GET     /api/v1/tasks/search?q=                 # Tìm kiếm tasks

# Bulk operations
POST    /api/v1/tasks/bulk/status               # Cập nhật hàng loạt trạng thái
DELETE  /api/v1/tasks/bulk                      # Xóa hàng loạt
```

### 5.6 Task Request/Response Example

```json
// POST /api/v1/tasks
// Request:
{
    "title": "Chuẩn bị lễ rằm tháng Giêng",
    "description": "Mua hoa quả, nhang đèn cho lễ rằm",
    "due_date": "2026-03-03",
    "due_time": "10:00",
    "is_lunar_based": true,
    "lunar_day": 15,
    "lunar_month": 1,
    "priority": "high",
    "category_id": "uuid-string",
    "tags": ["uuid-tag-1", "uuid-tag-2"],
    "subtasks": [
        { "title": "Mua hoa cúc" },
        { "title": "Mua trái cây" },
        { "title": "Mua nhang đèn" }
    ],
    "reminders": [
        { "type": "push", "minutes_before": 1440 },
        { "type": "push", "minutes_before": 60 }
    ]
}

// Response:
{
    "status": "success",
    "data": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "title": "Chuẩn bị lễ rằm tháng Giêng",
        "description": "Mua hoa quả, nhang đèn cho lễ rằm",
        "due_date": "2026-03-03",
        "due_time": "10:00:00",
        "is_lunar_based": true,
        "lunar_day": 15,
        "lunar_month": 1,
        "lunar_year": 2026,
        "status": "todo",
        "priority": "high",
        "progress": 0,
        "category": {
            "id": "uuid-string",
            "name": "Gia đình",
            "color": "#EF4444"
        },
        "tags": [
            { "id": "uuid-tag-1", "name": "Lễ tết", "color": "#F59E0B" }
        ],
        "subtasks": [
            { "id": "uuid", "title": "Mua hoa cúc", "status": "todo" },
            { "id": "uuid", "title": "Mua trái cây", "status": "todo" },
            { "id": "uuid", "title": "Mua nhang đèn", "status": "todo" }
        ],
        "subtask_progress": {
            "total": 3,
            "completed": 0,
            "percentage": 0
        },
        "reminders_count": 2,
        "created_at": "2026-03-01T08:00:00Z",
        "updated_at": "2026-03-01T08:00:00Z"
    }
}
```

### 5.7 Note Endpoints

```
# CRUD
POST    /api/v1/notes                           # Tạo note mới
GET     /api/v1/notes                           # Danh sách notes (with filters)
GET     /api/v1/notes/:id                       # Chi tiết note
PUT     /api/v1/notes/:id                       # Cập nhật note
DELETE  /api/v1/notes/:id                       # Xóa note (soft delete)

# Actions
PATCH   /api/v1/notes/:id/pin                   # Pin/Unpin note
PATCH   /api/v1/notes/:id/archive               # Archive/Unarchive
PATCH   /api/v1/notes/:id/color                 # Đổi màu note
POST    /api/v1/notes/:id/share                 # Chia sẻ note
DELETE  /api/v1/notes/:id/share/:userId          # Hủy chia sẻ

# Queries
GET     /api/v1/notes/date/:date                # Notes theo ngày
GET     /api/v1/notes/pinned                    # Notes đã pin
GET     /api/v1/notes/archived                  # Notes đã archive
GET     /api/v1/notes/search?q=                 # Full-text search
GET     /api/v1/notes/tags/:tagId               # Notes theo tag

# Attachments
POST    /api/v1/notes/:id/attachments           # Upload attachment
DELETE  /api/v1/notes/:id/attachments/:attachId  # Xóa attachment
```

### 5.8 Other Endpoints

```
# User Profile
GET     /api/v1/users/me                        # Thông tin user
PUT     /api/v1/users/me                        # Cập nhật profile
PUT     /api/v1/users/me/password               # Đổi mật khẩu
PUT     /api/v1/users/me/preferences            # Cập nhật preferences
POST    /api/v1/users/me/avatar                 # Upload avatar

# Categories
GET     /api/v1/categories                      # Danh sách categories
POST    /api/v1/categories                      # Tạo category
PUT     /api/v1/categories/:id                  # Cập nhật
DELETE  /api/v1/categories/:id                  # Xóa

# Tags
GET     /api/v1/tags                            # Danh sách tags
POST    /api/v1/tags                            # Tạo tag
PUT     /api/v1/tags/:id                        # Cập nhật
DELETE  /api/v1/tags/:id                        # Xóa

# Reminders
GET     /api/v1/reminders                       # Danh sách reminders
POST    /api/v1/reminders                       # Tạo reminder
PUT     /api/v1/reminders/:id                   # Cập nhật
DELETE  /api/v1/reminders/:id                   # Xóa

# Activity
GET     /api/v1/activity                        # Activity log
```

### 5.9 Common Query Parameters

```
# Pagination
?page=1&limit=20

# Sorting
?sort_by=created_at&sort_order=desc

# Filtering
?status=todo,in_progress
?priority=high,urgent
?category_id=uuid
?tag_id=uuid
?date_from=2026-03-01&date_to=2026-03-31

# Search
?q=search_term
```

### 5.10 Error Response Format

```json
{
    "status": "error",
    "error": {
        "code": "VALIDATION_ERROR",
        "message": "Dữ liệu không hợp lệ",
        "details": [
            {
                "field": "title",
                "message": "Tiêu đề không được để trống"
            },
            {
                "field": "due_date",
                "message": "Ngày phải có định dạng YYYY-MM-DD"
            }
        ]
    }
}
```

---

## 6. Frontend Design

### 6.1 Project Structure

```
frontend/
├── app/                            # Next.js App Router
│   ├── (auth)/                     # Auth layout group
│   │   ├── login/
│   │   │   └── page.tsx
│   │   ├── register/
│   │   │   └── page.tsx
│   │   └── layout.tsx
│   ├── (dashboard)/                # Dashboard layout group
│   │   ├── calendar/
│   │   │   ├── page.tsx            # Calendar view
│   │   │   ├── [date]/
│   │   │   │   └── page.tsx        # Day detail view
│   │   │   └── loading.tsx
│   │   ├── tasks/
│   │   │   ├── page.tsx            # Tasks list
│   │   │   ├── [id]/
│   │   │   │   └── page.tsx        # Task detail
│   │   │   └── loading.tsx
│   │   ├── notes/
│   │   │   ├── page.tsx            # Notes list
│   │   │   ├── [id]/
│   │   │   │   └── page.tsx        # Note editor
│   │   │   └── loading.tsx
│   │   ├── settings/
│   │   │   └── page.tsx
│   │   └── layout.tsx              # Dashboard layout with sidebar
│   ├── api/                        # API routes (BFF)
│   │   └── auth/
│   │       └── [...nextauth]/
│   │           └── route.ts
│   ├── layout.tsx                  # Root layout
│   ├── page.tsx                    # Landing page
│   └── globals.css
│
├── components/
│   ├── ui/                         # Shadcn UI components
│   │   ├── button.tsx
│   │   ├── dialog.tsx
│   │   ├── dropdown-menu.tsx
│   │   ├── input.tsx
│   │   ├── calendar.tsx
│   │   └── ...
│   ├── calendar/
│   │   ├── CalendarGrid.tsx        # Lưới lịch tháng
│   │   ├── CalendarHeader.tsx      # Header với navigation
│   │   ├── CalendarDay.tsx         # Ô ngày với âm/dương
│   │   ├── LunarInfo.tsx           # Thông tin âm lịch
│   │   ├── WeekView.tsx            # Giao diện tuần
│   │   ├── DayView.tsx             # Giao diện ngày chi tiết
│   │   ├── MiniCalendar.tsx        # Lịch nhỏ sidebar
│   │   ├── HolidayBadge.tsx        # Badge ngày lễ
│   │   ├── GoodDayIndicator.tsx    # Chỉ báo ngày tốt/xấu
│   │   └── DatePicker.tsx          # Date picker with lunar
│   ├── tasks/
│   │   ├── TaskList.tsx            # Danh sách tasks
│   │   ├── TaskItem.tsx            # Item task
│   │   ├── TaskForm.tsx            # Form tạo/sửa task
│   │   ├── TaskDetail.tsx          # Chi tiết task
│   │   ├── TaskFilters.tsx         # Bộ lọc
│   │   ├── TaskKanban.tsx          # Kanban board view
│   │   ├── SubtaskList.tsx         # Danh sách subtask
│   │   └── TaskStats.tsx           # Thống kê
│   ├── notes/
│   │   ├── NoteList.tsx            # Danh sách notes
│   │   ├── NoteCard.tsx            # Card note
│   │   ├── NoteEditor.tsx          # Rich text editor
│   │   ├── NoteToolbar.tsx         # Toolbar editor
│   │   └── NoteSearch.tsx          # Tìm kiếm
│   ├── layout/
│   │   ├── Sidebar.tsx             # Sidebar navigation
│   │   ├── Header.tsx              # Top header
│   │   ├── Footer.tsx
│   │   └── MobileNav.tsx           # Mobile navigation
│   └── shared/
│       ├── CategorySelect.tsx
│       ├── TagInput.tsx
│       ├── ColorPicker.tsx
│       ├── EmptyState.tsx
│       ├── LoadingSpinner.tsx
│       └── ConfirmDialog.tsx
│
├── hooks/
│   ├── useCalendar.ts              # Calendar data hook
│   ├── useLunar.ts                 # Lunar conversion hook
│   ├── useTasks.ts                 # Tasks CRUD hook
│   ├── useNotes.ts                 # Notes CRUD hook
│   ├── useAuth.ts                  # Auth hook
│   └── useDebounce.ts             # Debounce hook
│
├── lib/
│   ├── api.ts                      # API client (axios/fetch)
│   ├── lunar.ts                    # Lunar calendar utilities
│   ├── utils.ts                    # Helper functions
│   ├── constants.ts                # Constants
│   ├── validations.ts              # Zod schemas
│   └── date-utils.ts              # Date formatting
│
├── stores/
│   ├── calendarStore.ts            # Zustand calendar store
│   ├── taskStore.ts                # Zustand task store
│   ├── noteStore.ts                # Zustand note store
│   └── uiStore.ts                  # UI state store
│
├── types/
│   ├── calendar.ts                 # Calendar types
│   ├── task.ts                     # Task types
│   ├── note.ts                     # Note types
│   ├── user.ts                     # User types
│   └── api.ts                      # API response types
│
├── styles/
│   └── calendar.css                # Calendar-specific styles
│
├── public/
│   ├── images/
│   └── fonts/
│
├── next.config.js
├── tailwind.config.ts
├── tsconfig.json
└── package.json
```

### 6.2 UI/UX Design

#### Calendar Month View
```
┌─────────────────────────────────────────────────────────────────┐
│  ☰  Zplus Calendar          🔍   🔔 3   👤 Nguyễn Văn A       │
├─────────┬───────────────────────────────────────────────────────┤
│         │  ◄  Tháng 3, 2026 (Bính Ngọ)  ►    📅 Hôm nay       │
│ 📅 Lịch │  [Tháng] [Tuần] [Ngày]                               │
│         │                                                        │
│ ✅ Tasks│  T2     T3      T4      T5      T6      T7     CN    │
│         │ ┌──────┬──────┬──────┬──────┬──────┬──────┬──────┐   │
│ 📝 Notes│ │      │      │      │      │      │      │  1   │   │
│         │ │      │      │      │      │      │      │ 11/1 │   │
│ 🏷️ Tags │ │      │      │      │      │      │      │      │   │
│         │ ├──────┼──────┼──────┼──────┼──────┼──────┼──────┤   │
│ ⚙️ Cài  │ │  2   │ *3*  │  4   │  5   │  6   │  7   │  8   │   │
│   đặt   │ │ 12/1 │ 13/1 │ 14/1 │ 15/1 │ 16/1 │ 17/1 │ 18/1 │   │
│         │ │      │🔵2📝 │      │🟡Rằm │      │      │      │   │
│         │ ├──────┼──────┼──────┼──────┼──────┼──────┼──────┤   │
│ ──────  │ │  9   │  10  │  11  │  12  │  13  │  14  │  15  │   │
│         │ │ 19/1 │ 20/1 │ 21/1 │ 22/1 │ 23/1 │ 24/1 │ 25/1 │   │
│ Tháng 3 │ │      │      │      │      │      │      │      │   │
│ ┌─────┐ │ ├──────┼──────┼──────┼──────┼──────┼──────┼──────┤   │
│ │mini │ │ │  16  │  17  │  18  │  19  │  20  │  21  │  22  │   │
│ │cal  │ │ │ 26/1 │ 27/1 │ 28/1 │ 29/1 │ 1/2  │  2/2 │  3/2 │   │
│ └─────┘ │ │      │      │      │      │      │      │      │   │
│         │ ├──────┼──────┼──────┼──────┼──────┼──────┼──────┤   │
│ Sắp tới │ │  23  │  24  │  25  │  26  │  27  │  28  │  29  │   │
│ • Task1 │ │  4/2 │  5/2 │  6/2 │  7/2 │  8/2 │  9/2 │ 10/2 │   │
│ • Task2 │ │      │      │      │      │      │      │      │   │
│         │ ├──────┼──────┼──────┼──────┼──────┼──────┼──────┤   │
│         │ │  30  │  31  │      │      │      │      │      │   │
│         │ │ 11/2 │ 12/2 │      │      │      │      │      │   │
│         │ └──────┴──────┴──────┴──────┴──────┴──────┴──────┘   │
├─────────┴───────────────────────────────────────────────────────┤
│  Chú thích: 🔴 Ngày lễ  🟡 Rằm/Mùng 1  🔵 Có task  📝 Có note │
└─────────────────────────────────────────────────────────────────┘
```

#### Day Detail View
```
┌─────────────────────────────────────────────────────────────────┐
│  ← Quay lại     Thứ Ba, 03/03/2026                             │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  ┌─────────────────────────────────────────┐                    │
│  │  📅 DƯƠNG LỊCH        📅 ÂM LỊCH        │                    │
│  │  03 Tháng 03, 2026    13 Tháng 01, 2026 │                    │
│  │  Thứ Ba               Năm Bính Ngọ       │                    │
│  │                                           │                    │
│  │  🏛️ CAN CHI                               │                    │
│  │  Ngày: Bính Thân                          │                    │
│  │  Tháng: Canh Dần                          │                    │
│  │  Năm: Bính Ngọ                            │                    │
│  │                                           │                    │
│  │  ⭐ ĐÁNH GIÁ: Ngày Tốt                    │                    │
│  │  ✅ Nên: Cưới hỏi, Khai trương           │                    │
│  │  ❌ Không nên: Phá thổ, An táng           │                    │
│  │                                           │                    │
│  │  🕐 GIỜ HOÀNG ĐẠO                        │                    │
│  │  Tý (23-1) | Dần (3-5) | Mão (5-7)      │                    │
│  │  Ngọ (11-13) | Mùi (13-15) | Dậu (17-19)│                    │
│  └─────────────────────────────────────────┘                    │
│                                                                  │
│  ── TASKS (3) ──────────────────── [+ Thêm Task]               │
│  ┌─────────────────────────────────────────┐                    │
│  │  ☐ 🔴 Chuẩn bị lễ rằm          10:00   │                    │
│  │  ☐ 🟡 Họp team buổi sáng        09:00   │                    │
│  │  ☑️ 🟢 Review code PR #123       14:00   │                    │
│  └─────────────────────────────────────────┘                    │
│                                                                  │
│  ── NOTES (2) ──────────────────── [+ Thêm Note]               │
│  ┌──────────────────┐  ┌──────────────────┐                     │
│  │ 📝 Mua sắm       │  │ 📝 Ý tưởng       │                     │
│  │ - Hoa cúc         │  │ Thiết kế UI      │                     │
│  │ - Trái cây        │  │ cho trang...     │                     │
│  │ 📌 Pinned         │  │                  │                     │
│  └──────────────────┘  └──────────────────┘                     │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

### 6.3 Color Scheme & Design Tokens

```typescript
// Design Tokens
const theme = {
    colors: {
        primary: {
            50:  '#FFF7ED',
            100: '#FFEDD5',
            200: '#FED7AA',
            300: '#FDBA74',
            400: '#FB923C',
            500: '#F97316',     // Primary Orange (Warm Vietnamese feel)
            600: '#EA580C',
            700: '#C2410C',
            800: '#9A3412',
            900: '#7C2D12',
        },
        accent: {
            500: '#DC2626',     // Red for Lunar events
        },
        lunar: {
            fullMoon: '#FBBF24',    // Vàng cho ngày rằm
            newMoon: '#6B7280',     // Xám cho mùng 1
            holiday: '#DC2626',     // Đỏ cho ngày lễ
            goodDay: '#22C55E',     // Xanh cho ngày tốt
            badDay: '#EF4444',      // Đỏ cho ngày xấu
        },
        task: {
            urgent: '#DC2626',
            high: '#F97316',
            medium: '#3B82F6',
            low: '#6B7280',
        }
    },
    fonts: {
        heading: "'Be Vietnam Pro', sans-serif",
        body: "'Inter', sans-serif",
        mono: "'JetBrains Mono', monospace"
    }
};
```

---

## 7. Thuật Toán Chuyển Đổi Âm Dương Lịch

### 7.1 Tổng Quan

Thuật toán chuyển đổi âm dương lịch Việt Nam dựa trên:
- **Astronomical Algorithms** của Jean Meeus
- **Múi giờ UTC+7** (Việt Nam)
- **Tính toán Sóc (New Moon)** và **Trung khí (Major Solar Term)**

### 7.2 Core Algorithm (Go)

```go
package lunar

import (
    "math"
    "time"
)

const (
    PI          = math.Pi
    TimeZoneVN  = 7.0 // UTC+7
)

// LunarDate represents a Vietnamese lunar calendar date
type LunarDate struct {
    Day    int  `json:"day"`
    Month  int  `json:"month"`
    Year   int  `json:"year"`
    Leap   bool `json:"leap"`     // Tháng nhuận
}

// SolarDate represents a solar calendar date
type SolarDate struct {
    Day   int `json:"day"`
    Month int `json:"month"`
    Year  int `json:"year"`
}

// DayInfo contains full information for a calendar day
type DayInfo struct {
    Solar       SolarDate       `json:"solar"`
    Lunar       LunarDate       `json:"lunar"`
    CanChi      CanChiInfo      `json:"can_chi"`
    GoodBad     GoodBadInfo     `json:"good_bad"`
    Holidays    []HolidayInfo   `json:"holidays"`
    SolarTerm   *SolarTermInfo  `json:"solar_term,omitempty"`
}

// CanChiInfo - Thông tin Can Chi
type CanChiInfo struct {
    DayCan    string `json:"day_can"`
    DayChi    string `json:"day_chi"`
    DayFull   string `json:"day_full"`
    MonthCan  string `json:"month_can"`
    MonthChi  string `json:"month_chi"`
    MonthFull string `json:"month_full"`
    YearCan   string `json:"year_can"`
    YearChi   string `json:"year_chi"`
    YearFull  string `json:"year_full"`
}

// 10 Thiên Can
var ThienCan = [10]string{
    "Giáp", "Ất", "Bính", "Đinh", "Mậu",
    "Kỷ", "Canh", "Tân", "Nhâm", "Quý",
}

// 12 Địa Chi
var DiaChi = [12]string{
    "Tý", "Sửu", "Dần", "Mão", "Thìn", "Tỵ",
    "Ngọ", "Mùi", "Thân", "Dậu", "Tuất", "Hợi",
}

// JulianDayFromDate - Chuyển đổi ngày sang Julian Day Number
func JulianDayFromDate(day, month, year int) float64 {
    a := float64((14 - month) / 12)
    y := float64(year) + 4800 - a
    m := float64(month) + 12*a - 3
    
    jd := float64(day) + math.Floor((153*m+2)/5) + 365*y +
        math.Floor(y/4) - math.Floor(y/100) + math.Floor(y/400) - 32045
    
    return jd
}

// DateFromJulianDay - Chuyển Julian Day Number sang ngày
func DateFromJulianDay(jd float64) (int, int, int) {
    z := int(jd + 0.5)
    a := z
    if z >= 2299161 {
        alpha := int((float64(z) - 1867216.25) / 36524.25)
        a = z + 1 + alpha - alpha/4
    }
    b := a + 1524
    c := int((float64(b) - 122.1) / 365.25)
    d := int(365.25 * float64(c))
    e := int(float64(b-d) / 30.6001)
    
    day := b - d - int(30.6001*float64(e))
    month := e - 1
    if e > 13 {
        month = e - 13
    }
    year := c - 4716
    if month > 2 {
        year = c - 4715
    }
    // Adjust for year 0
    if year <= 0 {
        year--
    }
    return day, month, year
}

// NewMoon - Tính thời điểm Sóc (New Moon) thứ k
// Dựa trên Astronomical Algorithms - Jean Meeus
func NewMoon(k int) float64 {
    T := float64(k) / 1236.85
    T2 := T * T
    T3 := T2 * T
    
    // Mean phase
    jd := 2415020.75933 + 29.53058868*float64(k) +
        0.0001178*T2 - 0.000000155*T3 +
        0.00033*math.Sin((166.56+132.87*T-0.009173*T2)*PI/180)
    
    // Sun's mean anomaly
    M := 359.2242 + 29.10535608*float64(k) - 0.0000333*T2 - 0.00000347*T3
    M = M * PI / 180
    
    // Moon's mean anomaly
    Mprime := 306.0253 + 385.81691806*float64(k) + 0.0107306*T2 + 0.00001236*T3
    Mprime = Mprime * PI / 180
    
    // Moon's argument of latitude
    F := 21.2964 + 390.67050284*float64(k) - 0.0016528*T2 - 0.00000239*T3
    F = F * PI / 180
    
    // Corrections
    C1 := (0.1734-0.000393*T)*math.Sin(M) +
        0.0021*math.Sin(2*M) -
        0.4068*math.Sin(Mprime) +
        0.0161*math.Sin(2*Mprime) -
        0.0004*math.Sin(3*Mprime) +
        0.0104*math.Sin(2*F) -
        0.0051*math.Sin(M+Mprime) -
        0.0074*math.Sin(M-Mprime) +
        0.0004*math.Sin(2*F+M) -
        0.0004*math.Sin(2*F-M) -
        0.0006*math.Sin(2*F+Mprime) +
        0.0010*math.Sin(2*F-Mprime) +
        0.0005*math.Sin(M+2*Mprime)
    
    return jd + C1
}

// SunLongitude - Tính kinh độ mặt trời tại thời điểm JD
func SunLongitude(jd float64) float64 {
    T := (jd - 2451545.0) / 36525.0
    T2 := T * T
    
    // Geometric mean longitude
    L0 := 280.46646 + 36000.76983*T + 0.0003032*T2
    L0 = math.Mod(L0, 360)
    if L0 < 0 {
        L0 += 360
    }
    
    // Mean anomaly
    M := 357.52911 + 35999.05029*T - 0.0001537*T2
    M = M * PI / 180
    
    // Equation of center
    C := (1.914602 - 0.004817*T - 0.000014*T2) * math.Sin(M)
    C += (0.019993 - 0.000101*T) * math.Sin(2*M)
    C += 0.000289 * math.Sin(3*M)
    
    // Sun's true longitude
    theta := L0 + C
    
    // Apparent longitude
    omega := 125.04 - 1934.136*T
    lambda := theta - 0.00569 - 0.00478*math.Sin(omega*PI/180)
    
    lambda = math.Mod(lambda, 360)
    if lambda < 0 {
        lambda += 360
    }
    
    return lambda
}

// getSolarTermIndex - Xác định Trung khí
// Trung khí chia năm thành 12 phần, mỗi phần 30°
func getSolarTermIndex(sunLong float64) int {
    return int(math.Floor(sunLong / 30))
}

// SolarToLunar - Chuyển đổi Dương lịch sang Âm lịch Việt Nam
func SolarToLunar(solarDay, solarMonth, solarYear int) LunarDate {
    // Implementation using New Moon and Solar Term calculations
    // with Vietnam timezone (UTC+7)
    
    jd := JulianDayFromDate(solarDay, solarMonth, solarYear)
    
    // Find the lunar month containing this date
    k := int(math.Floor((float64(solarYear)-2000)*12.3685))
    
    // Search for the new moon before this date
    for NewMoon(k)+TimeZoneVN/24.0 > jd {
        k--
    }
    
    // ... (full implementation with leap month detection)
    
    return LunarDate{
        Day:   0, // calculated
        Month: 0, // calculated
        Year:  0, // calculated
        Leap:  false,
    }
}

// CanChiOfDay - Tính Can Chi của ngày
func CanChiOfDay(jd float64) CanChiInfo {
    jdn := int(jd + 0.5)
    canIndex := (jdn + 9) % 10
    chiIndex := (jdn + 1) % 12
    
    return CanChiInfo{
        DayCan:  ThienCan[canIndex],
        DayChi:  DiaChi[chiIndex],
        DayFull: ThienCan[canIndex] + " " + DiaChi[chiIndex],
    }
}

// CanChiOfYear - Tính Can Chi của năm
func CanChiOfYear(lunarYear int) (string, string) {
    canIndex := (lunarYear + 6) % 10
    chiIndex := (lunarYear + 8) % 12
    return ThienCan[canIndex], DiaChi[chiIndex]
}

// CanChiOfMonth - Tính Can Chi của tháng
func CanChiOfMonth(lunarMonth, lunarYear int) (string, string) {
    yearCan := (lunarYear + 6) % 10
    monthCanStart := (yearCan % 5) * 2 + 2
    canIndex := (monthCanStart + lunarMonth - 1) % 10
    chiIndex := (lunarMonth + 1) % 12
    return ThienCan[canIndex], DiaChi[chiIndex]
}
```

### 7.3 24 Tiết Khí (Solar Terms)

```go
// 24 Tiết khí
var SolarTerms = [24]string{
    "Xuân phân",    // 0°   - Vernal Equinox
    "Thanh minh",   // 15°
    "Cốc vũ",      // 30°
    "Lập hạ",       // 45°
    "Tiểu mãn",     // 60°
    "Mang chủng",   // 75°
    "Hạ chí",       // 90°  - Summer Solstice
    "Tiểu thử",     // 105°
    "Đại thử",      // 120°
    "Lập thu",       // 135°
    "Xử thử",       // 150°
    "Bạch lộ",      // 165°
    "Thu phân",      // 180° - Autumnal Equinox
    "Hàn lộ",       // 195°
    "Sương giáng",  // 210°
    "Lập đông",      // 225°
    "Tiểu tuyết",   // 240°
    "Đại tuyết",    // 255°
    "Đông chí",      // 270° - Winter Solstice
    "Tiểu hàn",     // 285°
    "Đại hàn",      // 300°
    "Lập xuân",      // 315°
    "Vũ thủy",      // 330°
    "Kinh trập",    // 345°
}
```

---

## 8. Chức Năng Note & Task

### 8.1 Notes Feature

#### Chức năng chính:
- **Rich Text Editor** (Tiptap/ProseMirror):
  - Heading (H1-H3)
  - Bold, Italic, Underline, Strikethrough
  - Bullet list, Numbered list
  - Checklist (todo items)
  - Code block
  - Quote block
  - Image embed
  - Link
  - Table
  - Highlight colors
  
- **Organization**:
  - Pin/Unpin notes
  - Archive
  - Color coding (8 colors)
  - Categories & Tags
  - Drag & Drop sorting
  
- **Association with Calendar**:
  - Gắn note vào ngày cụ thể (Dương/Âm lịch)
  - Hiển thị note trên calendar view
  - Quick note từ calendar day click
  
- **Search**:
  - Full-text search (MongoDB text index)
  - Filter by date range, category, tag
  - Search highlighting

- **Sharing** (P2):
  - Share note với user khác (read/write)
  - Generate shareable link

#### Note Views:
1. **Grid View** - Cards layout (kiểu Google Keep)
2. **List View** - Danh sách compact
3. **Calendar View** - Notes trên lịch

### 8.2 Tasks Feature

#### Chức năng chính:
- **CRUD Operations**:
  - Tạo task với title, description, due date/time
  - Subtasks (nested tasks)
  - Priority levels: Low, Medium, High, Urgent
  - Status: Todo → In Progress → Done / Cancelled

- **Scheduling**:
  - Due date & time
  - Start date & time
  - **Lunar-based scheduling** (nhắc theo Âm lịch)
  - Recurring tasks (daily, weekly, monthly, yearly, yearly_lunar)
  - RRULE format support

- **Organization**:
  - Categories with colors & icons
  - Tags (multi-select)
  - Drag & Drop reordering
  - Estimated & Actual hours tracking

- **Views**:
  1. **List View** - Danh sách tasks
  2. **Kanban Board** - Columns: Todo | In Progress | Done
  3. **Calendar View** - Tasks trên lịch
  4. **Timeline View** - Gantt-style timeline

- **Reminders**:
  - Push notification
  - Email notification
  - Multiple reminders per task
  - Smart reminders (before due date)

- **Statistics**:
  - Tasks completed today/week/month
  - Completion rate
  - Overdue tasks count
  - Time spent analysis
  - Category distribution chart

#### Task Workflow:
```
[Create] → [Todo] → [In Progress] → [Done]
                  ↘ [Cancelled]
```

### 8.3 Note & Task Integration

```
┌──────────────────────────────────────────┐
│           Calendar Day View               │
│                                           │
│  ┌─────────────────────────────────────┐  │
│  │  📅 03/03/2026 | 13/01 Âm lịch     │  │
│  └─────────────────────────────────────┘  │
│                                           │
│  ┌── Tasks ────────────┐                  │
│  │ ☐ Task 1            │ ←── Linked      │
│  │ ☑ Task 2            │                  │
│  │ + Add Task          │                  │
│  └─────────────────────┘                  │
│           ↕ Cross-reference               │
│  ┌── Notes ────────────┐                  │
│  │ 📝 Note 1           │ ←── Linked      │
│  │ 📝 Note 2           │                  │
│  │ + Add Note          │                  │
│  └─────────────────────┘                  │
│                                           │
│  • Notes có thể chứa task references     │
│  • Tasks có thể link tới notes           │
│  • Cả hai đều gắn vào ngày âm/dương     │
└──────────────────────────────────────────┘
```

---

## 9. Caching Strategy

### 9.1 Cache Layers

```
Request → Redis Cache → Application Cache → Database
            (L1)            (L2)              (L3)
```

### 9.2 Caching Rules

| Data | TTL | Strategy | Invalidation |
|------|-----|----------|-------------|
| Lunar calendar data | 30 days | Pre-compute monthly | Không cần (static) |
| Holidays | 1 year | Pre-compute yearly | Manual on update |
| Good/Bad days | 30 days | Pre-compute monthly | Không cần (static) |
| User profile | 1 hour | Cache-aside | On profile update |
| User tasks (by date) | 5 min | Cache-aside | On task CRUD |
| User notes (by date) | 5 min | Cache-aside | On note CRUD |
| API responses | Varies | ETag + Cache-Control | On data change |

### 9.3 Cache Warming

```go
// Pre-compute lunar data cho năm hiện tại + 1 năm tiếp theo
func WarmLunarCache(redisClient *redis.Client) {
    currentYear := time.Now().Year()
    for year := currentYear; year <= currentYear+1; year++ {
        for month := 1; month <= 12; month++ {
            data := computeLunarMonth(year, month)
            key := fmt.Sprintf("lunar:%d:%d", year, month)
            redisClient.HSet(ctx, key, data)
            redisClient.Expire(ctx, key, 30*24*time.Hour)
        }
    }
}
```

---

## 10. Authentication & Authorization

### 10.1 Auth Flow

```
┌──────────┐     ┌──────────┐     ┌──────────┐     ┌──────────┐
│  Client   │────▶│  Next.js  │────▶│ Go Fiber │────▶│PostgreSQL│
│ (Browser) │     │   BFF     │     │   API    │     │  + Redis │
└──────────┘     └──────────┘     └──────────┘     └──────────┘

1. User đăng nhập → Next.js BFF
2. BFF gửi credentials → Go Fiber API
3. API verify → issue JWT (access + refresh token)
4. Access Token → Redis (session) + Cookie (httpOnly)
5. Refresh Token → PostgreSQL + Cookie (httpOnly)
```

### 10.2 JWT Structure

```json
// Access Token (short-lived: 15 minutes)
{
    "sub": "user-uuid",
    "email": "user@email.com",
    "username": "nguyenvana",
    "role": "user",
    "iat": 1709424000,
    "exp": 1709424900
}

// Refresh Token (long-lived: 7 days)
{
    "sub": "user-uuid",
    "device": "device-fingerprint",
    "jti": "unique-token-id",
    "iat": 1709424000,
    "exp": 1710028800
}
```

### 10.3 Security Measures

- **HTTPS** everywhere
- **HttpOnly, Secure, SameSite** cookies
- **CORS** whitelist
- **Rate Limiting**: 100 requests/minute per user
- **Input Validation** (Go Validator + Zod on frontend)
- **SQL Injection Protection** (GORM parameterized queries)
- **XSS Protection** (Content Security Policy)
- **CSRF Token** for mutations
- **Password Hashing** (bcrypt with cost 12)
- **Account lockout** after 5 failed attempts

---

## 11. Deployment

### 11.1 Docker Compose (Development)

```yaml
# docker-compose.yml
version: '3.8'

services:
  # ========================
  # BACKEND - Go Fiber
  # ========================
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=development
      - APP_PORT=8080
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_NAME=zplus_calendar
      - DB_USER=zplus
      - DB_PASSWORD=zplus_secret
      - MONGO_URI=mongodb://mongo:27017/zplus_calendar
      - REDIS_URL=redis://redis:6379
      - JWT_SECRET=your-super-secret-key
      - JWT_ACCESS_EXPIRY=15m
      - JWT_REFRESH_EXPIRY=7d
    depends_on:
      - postgres
      - mongo
      - redis
    volumes:
      - ./backend:/app
    networks:
      - zplus-network
    restart: unless-stopped

  # ========================
  # FRONTEND - Next.js
  # ========================
  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    ports:
      - "3000:3000"
    environment:
      - NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
      - NEXTAUTH_SECRET=nextauth-secret
      - NEXTAUTH_URL=http://localhost:3000
    depends_on:
      - backend
    volumes:
      - ./frontend:/app
      - /app/node_modules
      - /app/.next
    networks:
      - zplus-network
    restart: unless-stopped

  # ========================
  # PostgreSQL
  # ========================
  postgres:
    image: postgres:16-alpine
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_DB=zplus_calendar
      - POSTGRES_USER=zplus
      - POSTGRES_PASSWORD=zplus_secret
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./database/init.sql:/docker-entrypoint-initdb.d/init.sql
    networks:
      - zplus-network
    restart: unless-stopped

  # ========================
  # MongoDB
  # ========================
  mongo:
    image: mongo:7
    ports:
      - "27017:27017"
    volumes:
      - mongo_data:/data/db
    networks:
      - zplus-network
    restart: unless-stopped

  # ========================
  # Redis
  # ========================
  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    command: redis-server --appendonly yes --maxmemory 256mb --maxmemory-policy allkeys-lru
    volumes:
      - redis_data:/data
    networks:
      - zplus-network
    restart: unless-stopped

  # ========================
  # Nginx (Reverse Proxy)
  # ========================
  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx/nginx.conf:/etc/nginx/nginx.conf
      - ./nginx/ssl:/etc/nginx/ssl
    depends_on:
      - backend
      - frontend
    networks:
      - zplus-network
    restart: unless-stopped

volumes:
  postgres_data:
  mongo_data:
  redis_data:

networks:
  zplus-network:
    driver: bridge
```

### 11.2 Backend Dockerfile

```dockerfile
# backend/Dockerfile
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/server ./cmd/server

# Production image
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Ho_Chi_Minh

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/configs ./configs

EXPOSE 8080

CMD ["./server"]
```

### 11.3 Backend Project Structure

```
backend/
├── cmd/
│   └── server/
│       └── main.go                 # Entry point
├── internal/
│   ├── config/
│   │   └── config.go               # Configuration
│   ├── middleware/
│   │   ├── auth.go                  # JWT middleware
│   │   ├── cors.go                  # CORS
│   │   ├── ratelimit.go            # Rate limiting
│   │   ├── logger.go               # Request logging
│   │   └── recovery.go             # Panic recovery
│   ├── handler/                     # HTTP Handlers (Controllers)
│   │   ├── auth_handler.go
│   │   ├── calendar_handler.go
│   │   ├── task_handler.go
│   │   ├── note_handler.go
│   │   ├── category_handler.go
│   │   ├── tag_handler.go
│   │   ├── user_handler.go
│   │   └── reminder_handler.go
│   ├── service/                     # Business Logic
│   │   ├── auth_service.go
│   │   ├── calendar_service.go
│   │   ├── task_service.go
│   │   ├── note_service.go
│   │   ├── category_service.go
│   │   ├── tag_service.go
│   │   ├── user_service.go
│   │   ├── reminder_service.go
│   │   └── notification_service.go
│   ├── repository/                  # Data Access
│   │   ├── postgres/
│   │   │   ├── user_repo.go
│   │   │   ├── task_repo.go
│   │   │   ├── category_repo.go
│   │   │   ├── tag_repo.go
│   │   │   └── reminder_repo.go
│   │   ├── mongo/
│   │   │   ├── note_repo.go
│   │   │   ├── calendar_event_repo.go
│   │   │   └── activity_repo.go
│   │   └── redis/
│   │       ├── cache_repo.go
│   │       └── session_repo.go
│   ├── model/                       # Data Models
│   │   ├── user.go
│   │   ├── task.go
│   │   ├── note.go
│   │   ├── category.go
│   │   ├── tag.go
│   │   ├── reminder.go
│   │   ├── calendar_event.go
│   │   └── holiday.go
│   ├── dto/                         # Data Transfer Objects
│   │   ├── request/
│   │   │   ├── auth_request.go
│   │   │   ├── task_request.go
│   │   │   └── note_request.go
│   │   └── response/
│   │       ├── auth_response.go
│   │       ├── calendar_response.go
│   │       ├── task_response.go
│   │       └── note_response.go
│   ├── lunar/                       # Lunar Calendar Engine
│   │   ├── converter.go            # Solar ↔ Lunar conversion
│   │   ├── can_chi.go              # Can Chi calculation
│   │   ├── good_bad_day.go         # Good/Bad day calculator
│   │   ├── solar_term.go           # 24 Solar terms
│   │   ├── holiday.go              # Vietnamese holidays
│   │   └── lunar_test.go           # Unit tests
│   ├── database/
│   │   ├── postgres.go             # PostgreSQL connection
│   │   ├── mongo.go                # MongoDB connection
│   │   ├── redis.go                # Redis connection
│   │   └── migrations/             # SQL migrations
│   │       ├── 001_create_users.up.sql
│   │       ├── 001_create_users.down.sql
│   │       ├── 002_create_tasks.up.sql
│   │       └── ...
│   ├── router/
│   │   └── router.go               # Route definitions
│   └── common/
│       ├── errors.go               # Custom errors
│       ├── response.go             # Standard response
│       ├── pagination.go           # Pagination helper
│       └── validator.go            # Custom validators
├── pkg/                             # Shared packages
│   ├── logger/
│   │   └── logger.go
│   └── utils/
│       ├── hash.go
│       ├── jwt.go
│       └── string.go
├── configs/
│   ├── config.yaml
│   └── config.example.yaml
├── docs/
│   └── swagger/
│       └── swagger.json
├── go.mod
├── go.sum
├── Dockerfile
├── Makefile
└── .air.toml                       # Hot reload config
```

### 11.4 Production Deployment Architecture

```
                    ┌─────────────┐
                    │  CloudFlare  │
                    │     CDN      │
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │   Nginx /   │
                    │  Traefik    │
                    │  (LB + SSL) │
                    └──────┬──────┘
                           │
              ┌────────────┼────────────┐
              │            │            │
       ┌──────▼───┐ ┌─────▼────┐ ┌────▼─────┐
       │ Frontend  │ │ Frontend │ │ Frontend │
       │ Next.js   │ │ Next.js  │ │ Next.js  │
       │ Instance 1│ │Instance 2│ │Instance 3│
       └──────┬───┘ └─────┬────┘ └────┬─────┘
              │            │            │
              └────────────┼────────────┘
                           │
              ┌────────────┼────────────┐
              │            │            │
       ┌──────▼───┐ ┌─────▼────┐ ┌────▼─────┐
       │ Backend   │ │ Backend  │ │ Backend  │
       │ Go Fiber  │ │ Go Fiber │ │ Go Fiber │
       │ Instance 1│ │Instance 2│ │Instance 3│
       └──────┬───┘ └─────┬────┘ └────┬─────┘
              │            │            │
              └────────────┼────────────┘
                           │
         ┌─────────────────┼─────────────────┐
         │                 │                  │
  ┌──────▼──────┐  ┌──────▼──────┐  ┌───────▼──────┐
  │ PostgreSQL  │  │   MongoDB   │  │    Redis     │
  │  Primary    │  │  Primary    │  │  Primary     │
  │  + Replica  │  │  + Replica  │  │  + Sentinel  │
  └─────────────┘  └─────────────┘  └──────────────┘
```

---

## 12. Đề Xuất Mở Rộng

### 12.1 Tính Năng Đề Xuất (Phase 2)

| # | Tính Năng | Mô Tả | Ưu Tiên |
|---|-----------|--------|---------|
| 1 | **📱 Mobile App (React Native/Flutter)** | Ứng dụng mobile với offline support, push notifications | Cao |
| 2 | **🔄 Google Calendar Sync** | Đồng bộ 2 chiều với Google Calendar | Cao |
| 3 | **📊 Dashboard & Analytics** | Thống kê productivity, task completion trends, time tracking | Trung bình |
| 4 | **👥 Team/Family Calendar** | Lịch chia sẻ cho nhóm/gia đình, assign tasks cho thành viên | Cao |
| 5 | **🤖 AI Smart Suggestions** | AI gợi ý ngày tốt cho sự kiện, tự động phân loại notes, smart scheduling | Trung bình |
| 6 | **🌐 Multi-language** | Hỗ trợ Tiếng Anh, Tiếng Trung, Tiếng Nhật (các nước dùng Âm lịch) | Thấp |
| 7 | **📧 Email Digest** | Tóm tắt hàng ngày/tuần qua email (tasks, events, lunar info) | Thấp |
| 8 | **🔌 Widget** | Widget cho iOS/Android, embed calendar widget cho website | Trung bình |
| 9 | **📆 iCal Export/Import** | Export/Import lịch định dạng .ics | Trung bình |
| 10 | **🎨 Themes & Customization** | Custom themes, calendar skins, font sizes | Thấp |
| 11 | **⏰ Pomodoro Timer** | Tích hợp timer Pomodoro vào task management | Thấp |
| 12 | **🔗 Third-party Integrations** | Slack, Telegram, Notion, Trello webhooks | Trung bình |

### 12.2 Đề Xuất Kỹ Thuật

#### A. WebSocket Real-time Updates
```
Sử dụng WebSocket cho:
- Real-time task updates khi shared calendar
- Live notification delivery
- Collaborative note editing
- Calendar event sync across devices
```

#### B. Offline-First Architecture
```
- Service Worker cho PWA support
- IndexedDB cho local storage
- Background sync khi có mạng lại
- Conflict resolution strategy (Last-Write-Wins hoặc CRDT)
```

#### C. GraphQL API (Bổ sung)
```
Thêm GraphQL endpoint cho:
- Flexible queries (client chỉ lấy data cần thiết)
- Subscription cho real-time
- Giảm over-fetching trên mobile
```

#### D. Event Sourcing cho Calendar Events
```
- Lưu trữ mọi thay đổi dưới dạng events
- Replay events để rebuild state
- Audit trail cho shared calendars
- Undo/Redo support
```

#### E. Microservices Evolution Path
```
Khi scale lớn, tách thành:
1. Auth Service
2. Calendar Service (Lunar Engine)
3. Task Service
4. Note Service
5. Notification Service
6. File/Media Service

Sử dụng:
- gRPC cho inter-service communication
- NATS/RabbitMQ cho async messaging
- API Gateway (Kong/Traefik)
```

#### F. Performance Optimizations
```
- Pre-render static calendar pages (ISR với Next.js)
- Lazy loading cho notes content
- Virtual scrolling cho long task lists
- Image optimization & CDN
- Database query optimization (prepared statements, connection pooling)
- Redis pipeline cho batch operations
```

#### G. Monitoring & Observability
```
- Prometheus + Grafana cho metrics
- Jaeger cho distributed tracing
- ELK Stack cho centralized logging
- Sentry cho error tracking
- Uptime monitoring (UptimeRobot/Pingdom)
```

### 12.3 Ngày Lễ Việt Nam (Seed Data)

```json
[
    {"name": "Tết Dương lịch", "day": 1, "month": 1, "is_lunar": false, "type": "public_holiday", "is_day_off": true},
    {"name": "Tết Nguyên Đán", "day": 1, "month": 1, "is_lunar": true, "type": "public_holiday", "is_day_off": true, "days_off_count": 5},
    {"name": "Tết Nguyên tiêu (Rằm tháng Giêng)", "day": 15, "month": 1, "is_lunar": true, "type": "traditional"},
    {"name": "Ngày Valentine", "day": 14, "month": 2, "is_lunar": false, "type": "cultural"},
    {"name": "Ngày Quốc tế Phụ nữ", "day": 8, "month": 3, "is_lunar": false, "type": "cultural"},
    {"name": "Tết Hàn thực", "day": 3, "month": 3, "is_lunar": true, "type": "traditional"},
    {"name": "Giỗ Tổ Hùng Vương", "day": 10, "month": 3, "is_lunar": true, "type": "public_holiday", "is_day_off": true},
    {"name": "Ngày Giải phóng miền Nam", "day": 30, "month": 4, "is_lunar": false, "type": "public_holiday", "is_day_off": true},
    {"name": "Ngày Quốc tế Lao động", "day": 1, "month": 5, "is_lunar": false, "type": "public_holiday", "is_day_off": true},
    {"name": "Ngày sinh Chủ tịch Hồ Chí Minh", "day": 19, "month": 5, "is_lunar": false, "type": "memorial"},
    {"name": "Tết Đoan Ngọ", "day": 5, "month": 5, "is_lunar": true, "type": "traditional"},
    {"name": "Ngày Gia đình Việt Nam", "day": 28, "month": 6, "is_lunar": false, "type": "cultural"},
    {"name": "Lễ Vu Lan (Rằm tháng 7)", "day": 15, "month": 7, "is_lunar": true, "type": "traditional"},
    {"name": "Tết Trung Thu", "day": 15, "month": 8, "is_lunar": true, "type": "traditional"},
    {"name": "Ngày Quốc khánh", "day": 2, "month": 9, "is_lunar": false, "type": "public_holiday", "is_day_off": true, "days_off_count": 2},
    {"name": "Tết Trùng Cửu", "day": 9, "month": 9, "is_lunar": true, "type": "traditional"},
    {"name": "Ngày Phụ nữ Việt Nam", "day": 20, "month": 10, "is_lunar": false, "type": "cultural"},
    {"name": "Ngày Halloween", "day": 31, "month": 10, "is_lunar": false, "type": "cultural"},
    {"name": "Ngày Nhà giáo Việt Nam", "day": 20, "month": 11, "is_lunar": false, "type": "cultural"},
    {"name": "Ngày Quân đội Nhân dân", "day": 22, "month": 12, "is_lunar": false, "type": "memorial"},
    {"name": "Giáng sinh", "day": 25, "month": 12, "is_lunar": false, "type": "cultural"},
    {"name": "Tết Ông Công Ông Táo", "day": 23, "month": 12, "is_lunar": true, "type": "traditional"}
]
```

### 12.4 Timeline Phát Triển (Đề Xuất)

```
Phase 1 - MVP (8-10 tuần)
├── Tuần 1-2:  Setup project, Database, Auth
├── Tuần 3-4:  Lunar Calendar Engine + Calendar API
├── Tuần 5-6:  Task Module (Backend + Frontend)
├── Tuần 7-8:  Note Module (Backend + Frontend)
├── Tuần 9:    Calendar UI + Integration
└── Tuần 10:   Testing, Bug fixes, Deploy

Phase 2 - Enhancement (6-8 tuần)
├── Tuần 1-2:  Good/Bad Days, Solar Terms
├── Tuần 3-4:  Notifications, Reminders
├── Tuần 5-6:  Google Calendar Sync
└── Tuần 7-8:  Analytics Dashboard, Performance

Phase 3 - Scale (8-10 tuần)
├── Tuần 1-3:  Mobile App (React Native)
├── Tuần 4-6:  Team/Family Calendar, Sharing
├── Tuần 7-8:  AI Features, Smart Suggestions
└── Tuần 9-10: Monitoring, Security Audit
```

---

## Phụ Lục

### A. Environment Variables

```env
# Backend (.env)
APP_ENV=development
APP_PORT=8080
APP_NAME=zplus-calendar

# PostgreSQL
DB_HOST=localhost
DB_PORT=5432
DB_NAME=zplus_calendar
DB_USER=zplus
DB_PASSWORD=zplus_secret
DB_SSL_MODE=disable
DB_MAX_CONNECTIONS=25
DB_MAX_IDLE_CONNECTIONS=5

# MongoDB
MONGO_URI=mongodb://localhost:27017/zplus_calendar
MONGO_DB_NAME=zplus_calendar

# Redis
REDIS_URL=redis://localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT
JWT_SECRET=your-256-bit-secret-key-here
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=168h

# OAuth
GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
FACEBOOK_APP_ID=
FACEBOOK_APP_SECRET=

# Email (SMTP)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=noreply@zpluscalendar.com
SMTP_PASSWORD=

# File Upload
UPLOAD_MAX_SIZE=10485760    # 10MB
UPLOAD_PATH=./uploads
ALLOWED_EXTENSIONS=jpg,jpeg,png,gif,pdf,doc,docx

# Rate Limiting
RATE_LIMIT_MAX=100
RATE_LIMIT_WINDOW=60s
```

### B. Makefile

```makefile
.PHONY: dev build test migrate

# Development
dev:
	docker-compose up -d postgres mongo redis
	cd backend && air

dev-frontend:
	cd frontend && npm run dev

dev-all:
	docker-compose up

# Build
build-backend:
	cd backend && go build -o bin/server ./cmd/server

build-frontend:
	cd frontend && npm run build

# Test
test-backend:
	cd backend && go test ./... -v -cover

test-frontend:
	cd frontend && npm test

# Database
migrate-up:
	cd backend && go run cmd/migrate/main.go up

migrate-down:
	cd backend && go run cmd/migrate/main.go down

seed:
	cd backend && go run cmd/seed/main.go

# Docker
docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

# Swagger
swagger:
	cd backend && swag init -g cmd/server/main.go -o docs/swagger
```

---

> **📌 Ghi chú**: Document này là bản thiết kế ban đầu và có thể được cập nhật trong quá trình phát triển. Mọi thay đổi lớn cần được review và approve bởi team lead.
