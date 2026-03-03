# ✅ Zplus Calendar - Task Management

> **Ngày tạo**: 03/03/2026  
> **Cập nhật lần cuối**: 03/03/2026  
> **Tham chiếu**: [SYSTEM_DESIGN.md](./docs/SYSTEM_DESIGN.md)

## Tiến độ tổng quan

| Module | Trạng thái | Ghi chú |
|--------|-----------|---------|
| Infrastructure | ✅ | Config, DB, Logger, Mi| 6.1.1 | Viết `useCalendar` hook | 🔴 | ✅ | — | Fetch month data |
| 6.1.2 | Viết `useLunar` hook | 🔴 | ✅ | — | Lunar utils cho frontend |
| 6.1.3 | Viết `calendarStore` (Zustand) | 🔴 | ✅ | — | Selected date, view mode |
| 6.1.4 | Component `CalendarHeader` - Tháng/Năm navigation + view switcher | 🔴 | ✅ | — | — |
| 6.1.5 | Component `CalendarGrid` - Lưới lịch tháng | 🔴 | ✅ | — | 6 tuần × 7 ngày |
| 6.1.6 | Component `CalendarDay` - Ô ngày (solar + lunar + indicators) | 🔴 | ✅ | — | Hiển thị âm lịch |
| 6.1.7 | Component `LunarInfo` - Panel thông tin âm lịch chi tiết | 🔴 | ✅ | — | Can chi, tiết khí |e, Docker, Makefile |
| Auth Module | ✅ | Register, Login, Logout, Refresh, ChangePassword |
| Lunar Engine | ✅ | 13 tests pass, all algorithms implemented |
| Calendar API | ✅ | Today, Day, Month, Convert, Holidays |
| Task Module | ✅ | CRUD, search, by-date, status, tags, categories |
| Note Module | ✅ | MongoDB CRUD, search, by-date, pin |
| Seed Data | ✅ | Admin (role: admin), demo user, categories, tags, holidays |
| Frontend | ✅ | Auth, Calendar (public), Tasks, Notes, Settings pages |

## Quy ước

- ⬜ Chưa bắt đầu
- 🔄 Đang thực hiện
- ✅ Hoàn thành
- ❌ Hủy bỏ
- 🚫 Blocked
- ⏸️ Tạm hoãn

**Độ ưu tiên**: 🔴 Critical | 🟠 High | 🟡 Medium | 🟢 Low

---

## 📌 Phase 1 — MVP (Tuần 1 → Tuần 10)

---

### 🏗️ 1. Project Setup & Infrastructure (Tuần 1-2)

#### 1.1 Khởi tạo Repository & Cấu trúc

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 1.1.1 | Khởi tạo Git repo, `.gitignore`, branch strategy (main/develop/feature) | 🔴 | ✅ | — | — |
| 1.1.2 | Tạo cấu trúc thư mục Backend (Go Fiber - Clean Architecture) | 🔴 | ✅ | — | `cmd/`, `internal/`, `pkg/` |
| 1.1.3 | Tạo cấu trúc thư mục Frontend (Next.js 14 App Router) | 🔴 | ✅ | — | `app/`, `components/`, `lib/` |
| 1.1.4 | Tạo `Makefile` cho dev, build, test, migrate | 🟠 | ✅ | — | — |
| 1.1.5 | Tạo file `.env.example` cho Backend & Frontend | 🟠 | ✅ | — | — |
| 1.1.6 | Setup `.editorconfig`, `prettier`, `eslint` cho Frontend | 🟡 | ⬜ | — | — |
| 1.1.7 | Setup `golangci-lint` cho Backend | 🟡 | ⬜ | — | — |

#### 1.2 Docker & DevOps

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 1.2.1 | Viết `docker-compose.yml` (PostgreSQL, MongoDB, Redis, Backend, Frontend) | 🔴 | ✅ | — | — |
| 1.2.2 | Viết `Dockerfile` cho Backend (multi-stage build) | 🔴 | ✅ | — | golang:1.22-alpine |
| 1.2.3 | Viết `Dockerfile` cho Frontend (Next.js standalone) | 🔴 | ✅ | — | node:20-alpine |
| 1.2.4 | Config Nginx reverse proxy | 🟠 | ⬜ | — | Port 80/443 |
| 1.2.5 | Setup `Air` hot reload cho Go development | 🟡 | ✅ | — | `.air.toml` |
| 1.2.6 | Verify docker-compose up chạy toàn bộ stack thành công | 🔴 | ⬜ | — | Smoke test |

#### 1.3 Database Setup

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 1.3.1 | Setup PostgreSQL connection (GORM) | 🔴 | ✅ | — | Connection pooling |
| 1.3.2 | Setup MongoDB connection (Go Driver v2) | 🔴 | ✅ | — | — |
| 1.3.3 | Setup Redis connection (go-redis v9) | 🔴 | ✅ | — | — |
| 1.3.4 | Viết migration `001_create_users` | 🔴 | ✅ | — | 001_initial_schema.sql |
| 1.3.5 | Viết migration `002_create_categories` | 🔴 | ✅ | — | Combined in 001 |
| 1.3.6 | Viết migration `003_create_tags` | 🔴 | ✅ | — | Combined in 001 |
| 1.3.7 | Viết migration `004_create_tasks` + `task_tags` | 🔴 | ✅ | — | Combined in 001 |
| 1.3.8 | Viết migration `005_create_reminders` | 🟠 | ✅ | — | Combined in 001 |
| 1.3.9 | Viết migration `006_create_vietnamese_holidays` | 🟠 | ✅ | — | Combined in 001 |
| 1.3.10 | Viết migration `007_create_user_sessions` | 🟠 | ✅ | — | Combined in 001 |
| 1.3.11 | Tạo MongoDB indexes cho `notes` collection | 🔴 | ✅ | — | EnsureIndexes() |
| 1.3.12 | Tạo MongoDB indexes cho `calendar_events` collection | 🟠 | ⬜ | — | — |
| 1.3.13 | Tạo MongoDB indexes cho `activity_logs` (TTL 90 days) | 🟡 | ⬜ | — | — |
| 1.3.14 | Viết seed data cho `vietnamese_holidays` | 🟠 | ✅ | — | cmd/seed |
| 1.3.15 | Viết migration runner tool (`cmd/migrate/main.go`) | 🔴 | ✅ | — | GORM AutoMigrate |

#### 1.4 Backend Foundation

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 1.4.1 | Setup Go Fiber app với config (Viper) | 🔴 | ✅ | — | `config.yaml` |
| 1.4.2 | Setup structured logging (Zap) | 🔴 | ✅ | — | — |
| 1.4.3 | Viết common response helper (`Success`, `Error`, `Paginated`) | 🔴 | ✅ | — | — |
| 1.4.4 | Viết custom error types | 🔴 | ✅ | — | — |
| 1.4.5 | Viết pagination helper | 🟠 | ✅ | — | — |
| 1.4.6 | Setup request validator (go-playground/validator v10) | 🔴 | ✅ | — | Custom validators |
| 1.4.7 | Viết CORS middleware | 🔴 | ✅ | — | In main.go |
| 1.4.8 | Viết Recovery middleware (panic handler) | 🔴 | ✅ | — | In main.go |
| 1.4.9 | Viết Request Logger middleware | 🟠 | ✅ | — | — |
| 1.4.10 | Viết Rate Limiting middleware (Redis-based) | 🟠 | ✅ | — | 100 req/min |
| 1.4.11 | Setup Swagger documentation | 🟡 | ⬜ | — | swaggo/swag |
| 1.4.12 | Setup Dependency Injection container | 🟠 | ⬜ | — | Manual DI hoặc Wire |

---

### 🔐 2. Authentication & User Module (Tuần 2-3)

#### 2.1 Backend - Auth

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 2.1.1 | Viết `User` model (GORM) | 🔴 | ✅ | — | — |
| 2.1.2 | Viết `UserRepository` (PostgreSQL) | 🔴 | ✅ | — | — |
| 2.1.3 | Viết `SessionRepository` (Redis) | 🔴 | ✅ | — | — |
| 2.1.4 | Viết JWT utility (access + refresh token) | 🔴 | ✅ | — | golang-jwt/jwt v5 |
| 2.1.5 | Viết Password hashing utility (bcrypt) | 🔴 | ✅ | — | Cost 12 |
| 2.1.6 | Viết `AuthService` - Register | 🔴 | ✅ | — | Email validation, duplicate check |
| 2.1.7 | Viết `AuthService` - Login | 🔴 | ✅ | — | Issue JWT pair |
| 2.1.8 | Viết `AuthService` - Logout | 🔴 | ✅ | — | Invalidate session |
| 2.1.9 | Viết `AuthService` - Refresh Token | 🔴 | ✅ | — | Token rotation |
| 2.1.10 | Viết `AuthService` - Forgot Password | 🟠 | ⬜ | — | Send reset email |
| 2.1.11 | Viết `AuthService` - Reset Password | 🟠 | ⬜ | — | Token-based reset |
| 2.1.12 | Viết `AuthService` - Email Verification | 🟡 | ⬜ | — | — |
| 2.1.13 | Viết `AuthHandler` (routes + controllers) | 🔴 | ✅ | — | — |
| 2.1.14 | Viết Auth Middleware (JWT verification) | 🔴 | ✅ | — | — |
| 2.1.15 | Viết Auth DTOs (Request/Response) | 🔴 | ✅ | — | — |
| 2.1.16 | Unit tests cho AuthService | 🟠 | ⬜ | — | — |

#### 2.2 Backend - User Profile

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 2.2.1 | Viết `UserService` - Get Profile | 🔴 | ⬜ | — | — |
| 2.2.2 | Viết `UserService` - Update Profile | 🔴 | ⬜ | — | — |
| 2.2.3 | Viết `UserService` - Change Password | 🟠 | ⬜ | — | — |
| 2.2.4 | Viết `UserService` - Update Preferences | 🟠 | ⬜ | — | Theme, locale, calendar view |
| 2.2.5 | Viết `UserService` - Upload Avatar | 🟡 | ⬜ | — | — |
| 2.2.6 | Viết `UserHandler` (routes + controllers) | 🔴 | ⬜ | — | — |
| 2.2.7 | Cache user profile trong Redis (TTL 1h) | 🟠 | ⬜ | — | Invalidate on update |

#### 2.3 Frontend - Auth

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 2.3.1 | Setup Next.js project với TypeScript, Tailwind, Shadcn/UI | 🔴 | ✅ | — | — |
| 2.3.2 | Setup API client (axios/fetch wrapper) | 🔴 | ✅ | — | Interceptors, token refresh |
| 2.3.3 | Setup TanStack Query | 🔴 | ✅ | — | — |
| 2.3.4 | Setup Zustand stores | 🔴 | ✅ | — | Auth, UI stores |
| 2.3.5 | Viết `useAuth` hook | 🔴 | ✅ | — | — |
| 2.3.6 | Trang Login (UI + logic) | 🔴 | ✅ | — | — |
| 2.3.7 | Trang Register (UI + logic) | 🔴 | ✅ | — | — |
| 2.3.8 | Trang Forgot Password | 🟠 | ⬜ | — | — |
| 2.3.9 | Auth Layout (centered card) | 🔴 | ✅ | — | — |
| 2.3.10 | Protected route middleware (Next.js middleware.ts) | 🔴 | ✅ | — | Redirect to login |
| 2.3.11 | Form validation với Zod | 🟠 | ✅ | — | — |

---

### 🌙 3. Lunar Calendar Engine (Tuần 3-4)

#### 3.1 Core Algorithm

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 3.1.1 | Implement `JulianDayFromDate()` - Date → Julian Day | 🔴 | ✅ | — | — |
| 3.1.2 | Implement `DateFromJulianDay()` - Julian Day → Date | 🔴 | ✅ | — | — |
| 3.1.3 | Implement `NewMoon()` - Tính thời điểm Sóc (Jean Meeus) | 🔴 | ✅ | — | — |
| 3.1.4 | Implement `SunLongitude()` - Kinh độ Mặt trời | 🔴 | ✅ | — | — |
| 3.1.5 | Implement `SolarToLunar()` - Dương lịch → Âm lịch | 🔴 | ✅ | — | Timezone UTC+7 |
| 3.1.6 | Implement `LunarToSolar()` - Âm lịch → Dương lịch | 🔴 | ✅ | — | — |
| 3.1.7 | Xử lý **tháng nhuận** (Leap month) | 🔴 | ✅ | — | Bug fixed in getLeapMonthOffset |
| 3.1.8 | Implement `CanChiOfDay()` - Can Chi ngày | 🔴 | ✅ | — | — |
| 3.1.9 | Implement `CanChiOfMonth()` - Can Chi tháng | 🔴 | ✅ | — | — |
| 3.1.10 | Implement `CanChiOfYear()` - Can Chi năm | 🔴 | ✅ | — | — |
| 3.1.11 | Implement `CanChiOfHour()` - Can Chi giờ | 🟠 | ✅ | — | 12 giờ/ngày |
| 3.1.12 | Implement 24 Tiết khí (Solar Terms) | 🟠 | ✅ | — | — |
| 3.1.13 | Implement tính giờ Hoàng đạo / Hắc đạo | 🟠 | ✅ | — | — |
| 3.1.14 | Implement ngày tốt/xấu calculator | 🟡 | ✅ | — | — |
| 3.1.15 | Unit tests cho toàn bộ lunar engine (>= 95% coverage) | 🔴 | ✅ | — | 13 tests pass |
| 3.1.16 | Benchmark performance (< 1ms/conversion) | 🟠 | ⬜ | — | — |

#### 3.2 Calendar API

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 3.2.1 | Viết `CalendarService` - Get Month Data | 🔴 | ✅ | — | Full calendar grid |
| 3.2.2 | Viết `CalendarService` - Get Week Data | 🟠 | ⬜ | — | — |
| 3.2.3 | Viết `CalendarService` - Get Day Detail | 🔴 | ✅ | — | Full lunar + can chi + good/bad |
| 3.2.4 | Viết `CalendarService` - Solar to Lunar API | 🔴 | ✅ | — | — |
| 3.2.5 | Viết `CalendarService` - Lunar to Solar API | 🔴 | ✅ | — | — |
| 3.2.6 | Viết `CalendarService` - Get Holidays by Year | 🟠 | ✅ | — | — |
| 3.2.7 | Viết `CalendarService` - Get Today Info | 🔴 | ✅ | — | — |
| 3.2.8 | Viết `CalendarHandler` (routes + controllers) | 🔴 | ✅ | — | — |
| 3.2.9 | Viết Calendar DTOs (Response) | 🔴 | ✅ | — | — |
| 3.2.10 | Cache lunar data theo tháng trong Redis (TTL 30 days) | 🟠 | ⬜ | — | Pre-compute |
| 3.2.11 | Cache holidays theo năm trong Redis (TTL 1 year) | 🟡 | ⬜ | — | — |
| 3.2.12 | Viết cache warming job (năm hiện tại + 1) | 🟡 | ⬜ | — | — |

---

### ✅ 4. Task Module (Tuần 5-6)

#### 4.1 Backend - Tasks

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 4.1.1 | Viết `Task` model (GORM) | 🔴 | ✅ | — | — |
| 4.1.2 | Viết `Category` model (GORM) | 🔴 | ✅ | — | — |
| 4.1.3 | Viết `Tag` model (GORM) | 🔴 | ✅ | — | — |
| 4.1.4 | Viết `TaskRepository` - CRUD | 🔴 | ✅ | — | — |
| 4.1.5 | Viết `TaskRepository` - Query by date, range, status | 🔴 | ✅ | — | — |
| 4.1.6 | Viết `TaskRepository` - Search (full-text) | 🟠 | ✅ | — | ILIKE search |
| 4.1.7 | Viết `TaskRepository` - Subtasks | 🟠 | ✅ | — | Parent-child relation |
| 4.1.8 | Viết `TaskRepository` - Bulk operations | 🟡 | ⬜ | — | — |
| 4.1.9 | Viết `CategoryRepository` - CRUD | 🔴 | ✅ | — | — |
| 4.1.10 | Viết `TagRepository` - CRUD | 🔴 | ✅ | — | — |
| 4.1.11 | Viết `TaskService` - Create Task (with tags, subtasks) | 🔴 | ✅ | — | Transaction |
| 4.1.12 | Viết `TaskService` - Update Task | 🔴 | ✅ | — | — |
| 4.1.13 | Viết `TaskService` - Delete Task (soft delete) | 🔴 | ✅ | — | — |
| 4.1.14 | Viết `TaskService` - Update Status | 🔴 | ✅ | — | State machine |
| 4.1.15 | Viết `TaskService` - Get Tasks by Date | 🔴 | ✅ | — | — |
| 4.1.16 | Viết `TaskService` - Get Tasks by Range | 🟠 | ✅ | — | — |
| 4.1.17 | Viết `TaskService` - Get Overdue Tasks | 🟠 | ⬜ | — | — |
| 4.1.18 | Viết `TaskService` - Task Statistics | 🟡 | ⬜ | — | Count, completion rate |
| 4.1.19 | Viết `TaskService` - Lunar-based task scheduling | 🟠 | ⬜ | — | Convert lunar → solar date |
| 4.1.20 | Viết `TaskHandler` (routes + controllers) | 🔴 | ✅ | — | — |
| 4.1.21 | Viết `CategoryHandler` + `TagHandler` | 🔴 | ✅ | — | — |
| 4.1.22 | Viết Task DTOs (Request/Response) | 🔴 | ✅ | — | — |
| 4.1.23 | Cache tasks by date trong Redis (TTL 5 min) | 🟡 | ⬜ | — | Invalidate on CRUD |
| 4.1.24 | Unit tests cho TaskService | 🟠 | ⬜ | — | — |

#### 4.2 Frontend - Tasks

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 4.2.1 | Viết TypeScript types cho Task, Category, Tag | 🔴 | ✅ | — | `types/task.ts` |
| 4.2.2 | Viết `useTasks` hook (TanStack Query) | 🔴 | ✅ | — | CRUD + filters |
| 4.2.3 | Viết `taskStore` (Zustand) | 🟠 | ⬜ | — | UI state |
| 4.2.4 | Component `TaskList` - Danh sách tasks | 🔴 | ✅ | — | — |
| 4.2.5 | Component `TaskItem` - Task row/card | 🔴 | ✅ | — | Checkbox, priority badge |
| 4.2.6 | Component `TaskForm` - Tạo/sửa task (Dialog) | 🔴 | ✅ | — | Date picker, priority, category |
| 4.2.7 | Component `TaskDetail` - Chi tiết task | 🟠 | ⬜ | — | Subtasks, tags, progress |
| 4.2.8 | Component `TaskFilters` - Bộ lọc (status, priority, category, date) | 🟠 | ✅ | — | — |
| 4.2.9 | Component `SubtaskList` - Subtasks với checkbox | 🟠 | ⬜ | — | — |
| 4.2.10 | Component `TaskKanban` - Kanban board view | 🟡 | ⬜ | — | Drag & Drop (React DnD) |
| 4.2.11 | Component `TaskStats` - Thống kê | 🟡 | ⬜ | — | Charts |
| 4.2.12 | Trang `/tasks` - Task management page | 🔴 | ✅ | — | — |
| 4.2.13 | Trang `/tasks/[id]` - Task detail page | 🟠 | ⬜ | — | — |
| 4.2.14 | Quick add task từ calendar day | 🟠 | ⬜ | — | — |
| 4.2.15 | Bulk actions (select multiple, update status, delete) | 🟡 | ⬜ | — | — |

---

### 📝 5. Note Module (Tuần 7-8)

#### 5.1 Backend - Notes

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 5.1.1 | Viết `Note` model (MongoDB struct) | 🔴 | ✅ | — | NoteDocument DTO |
| 5.1.2 | Viết `NoteRepository` - CRUD (MongoDB) | 🔴 | ✅ | — | — |
| 5.1.3 | Viết `NoteRepository` - Query by date, category, tag | 🔴 | ✅ | — | — |
| 5.1.4 | Viết `NoteRepository` - Full-text search | 🟠 | ✅ | — | Regex search |
| 5.1.5 | Viết `NoteRepository` - Pin/Archive operations | 🟠 | ✅ | — | Pin filter |
| 5.1.6 | Viết `NoteService` - Create Note (with lunar date) | 🔴 | ✅ | — | — |
| 5.1.7 | Viết `NoteService` - Update Note | 🔴 | ✅ | — | — |
| 5.1.8 | Viết `NoteService` - Delete Note (soft delete, TTL 30 days) | 🔴 | ✅ | — | Hard delete for now |
| 5.1.9 | Viết `NoteService` - Pin/Unpin | 🟠 | ✅ | — | Via update |
| 5.1.10 | Viết `NoteService` - Archive/Unarchive | 🟠 | ⬜ | — | — |
| 5.1.11 | Viết `NoteService` - Change Color | 🟡 | ✅ | — | Via update |
| 5.1.12 | Viết `NoteService` - Search Notes | 🟠 | ✅ | — | — |
| 5.1.13 | Viết `NoteHandler` (routes + controllers) | 🔴 | ✅ | — | — |
| 5.1.14 | Viết Note DTOs (Request/Response) | 🔴 | ✅ | — | — |
| 5.1.15 | File upload service (attachments) | 🟡 | ⬜ | — | Max 10MB, image/pdf |
| 5.1.16 | Cache notes by date trong Redis (TTL 5 min) | 🟡 | ⬜ | — | — |
| 5.1.17 | Unit tests cho NoteService | 🟠 | ⬜ | — | — |

#### 5.2 Frontend - Notes

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 5.2.1 | Viết TypeScript types cho Note | 🔴 | ✅ | — | `types/note.ts` |
| 5.2.2 | Viết `useNotes` hook (TanStack Query) | 🔴 | ✅ | — | — |
| 5.2.3 | Viết `noteStore` (Zustand) | 🟠 | ⬜ | — | — |
| 5.2.4 | Setup Tiptap rich text editor | 🔴 | ⬜ | — | Extensions: bold, italic, list, etc. |
| 5.2.5 | Component `NoteEditor` - Rich text editor | 🔴 | ⬜ | — | Tiptap + toolbar |
| 5.2.6 | Component `NoteToolbar` - Formatting toolbar | 🔴 | ⬜ | — | — |
| 5.2.7 | Component `NoteList` - Danh sách notes | 🔴 | ✅ | — | — |
| 5.2.8 | Component `NoteCard` - Note card (Google Keep style) | 🔴 | ✅ | — | Color, pin badge |
| 5.2.9 | Component `NoteSearch` - Tìm kiếm notes | 🟠 | ⬜ | — | Debounce search |
| 5.2.10 | Grid View (masonry layout) | 🟠 | ⬜ | — | Responsive columns |
| 5.2.11 | List View | 🟡 | ⬜ | — | Compact list |
| 5.2.12 | Color picker cho notes | 🟡 | ✅ | — | 8 preset colors |
| 5.2.13 | Trang `/notes` - Notes management page | 🔴 | ✅ | — | — |
| 5.2.14 | Trang `/notes/[id]` - Note editor page | 🔴 | ⬜ | — | — |
| 5.2.15 | Quick add note từ calendar day | 🟠 | ⬜ | — | — |
| 5.2.16 | Checklist support trong note editor | 🟡 | ⬜ | — | Tiptap TaskList extension |
| 5.2.17 | Auto-save note (debounce 1s) | 🟠 | ⬜ | — | — |

---

### 📅 6. Calendar UI (Tuần 9)

#### 6.1 Frontend - Calendar Views

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 6.1.1 | Viết `useCalendar` hook | 🔴 | ⬜ | — | Fetch month data |
| 6.1.2 | Viết `useLunar` hook | 🔴 | ⬜ | — | Lunar utils cho frontend |
| 6.1.3 | Viết `calendarStore` (Zustand) | 🔴 | ⬜ | — | Selected date, view mode |
| 6.1.4 | Component `CalendarHeader` - Tháng/Năm navigation + view switcher | 🔴 | ⬜ | — | — |
| 6.1.5 | Component `CalendarGrid` - Lưới lịch tháng | 🔴 | ⬜ | — | 6 tuần × 7 ngày |
| 6.1.6 | Component `CalendarDay` - Ô ngày (solar + lunar + indicators) | 🔴 | ⬜ | — | Hiển thị âm lịch |
| 6.1.7 | Component `LunarInfo` - Panel thông tin âm lịch chi tiết | 🔴 | ⬜ | — | Can chi, tiết khí |
| 6.1.8 | Component `WeekView` - Giao diện tuần | 🟠 | ⬜ | — | Timeline style |
| 6.1.9 | Component `DayView` - Giao diện ngày chi tiết | 🔴 | ⬜ | — | Lunar + tasks + notes |
| 6.1.10 | Component `MiniCalendar` - Lịch nhỏ cho sidebar | 🟠 | ⬜ | — | — |
| 6.1.11 | Component `HolidayBadge` - Badge ngày lễ | 🟠 | ⬜ | — | Màu đỏ cho ngày lễ |
| 6.1.12 | Component `GoodDayIndicator` - Chỉ báo ngày tốt/xấu | 🟡 | ⬜ | — | Dot indicator |
| 6.1.13 | Component `DatePicker` - Date picker kèm âm lịch | 🟠 | ⬜ | — | Dùng cho task/note forms |
| 6.1.14 | Trang `/calendar` - Calendar main page | 🔴 | ✅ | — | — |
| 6.1.15 | Trang `/calendar/[date]` - Day detail page | 🔴 | ⬜ | — | — |
| 6.1.16 | Highlight ngày hôm nay | 🔴 | ✅ | — | — |
| 6.1.17 | Highlight ngày Rằm (15) và Mùng 1 | 🟠 | ✅ | — | Màu vàng/xám |
| 6.1.18 | Hiển thị task count + note count trên mỗi ngày | 🟠 | ⬜ | — | Badge dots |
| 6.1.19 | Keyboard navigation (← → ↑ ↓) | 🟡 | ⬜ | — | — |
| 6.1.20 | Responsive design (mobile/tablet/desktop) | 🔴 | ⬜ | — | — |

#### 6.2 Frontend - Layout & Navigation

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 6.2.1 | Component `Sidebar` - Navigation sidebar | 🔴 | ✅ | — | Calendar, Tasks, Notes, Settings |
| 6.2.2 | Component `Header` - Top header bar | 🔴 | ✅ | — | Search, notifications, user menu |
| 6.2.3 | Component `MobileNav` - Bottom navigation (mobile) | 🟠 | ✅ | — | — |
| 6.2.4 | Dashboard Layout (`(dashboard)/layout.tsx`) | 🔴 | ✅ | — | Sidebar + content area |
| 6.2.5 | Dark mode support | 🟡 | ✅ | — | `next-themes` |
| 6.2.6 | Loading skeletons cho calendar, tasks, notes | 🟠 | ⬜ | — | — |
| 6.2.7 | Empty states cho lists | 🟡 | ⬜ | — | — |
| 6.2.8 | Toast notifications (Sonner) | 🟠 | ✅ | — | — |

---

### 🧪 7. Testing & Bug Fixes (Tuần 10)

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 7.1 | Integration tests - Auth flow | 🔴 | ⬜ | — | Register → Login → Protected route |
| 7.2 | Integration tests - Calendar API | 🔴 | ⬜ | — | Verify lunar conversion accuracy |
| 7.3 | Integration tests - Task CRUD | 🔴 | ⬜ | — | — |
| 7.4 | Integration tests - Note CRUD | 🔴 | ⬜ | — | — |
| 7.5 | E2E tests - Calendar navigation | 🟠 | ⬜ | — | Playwright |
| 7.6 | E2E tests - Create task from calendar | 🟠 | ⬜ | — | — |
| 7.7 | E2E tests - Create note from calendar | 🟠 | ⬜ | — | — |
| 7.8 | Performance testing - API response < 200ms | 🟠 | ⬜ | — | — |
| 7.9 | Cross-browser testing (Chrome, Firefox, Safari) | 🟡 | ⬜ | — | — |
| 7.10 | Mobile responsive testing | 🟠 | ⬜ | — | — |
| 7.11 | Security audit (OWASP Top 10) | 🟠 | ⬜ | — | — |
| 7.12 | Fix bugs từ QA/testing | 🔴 | ⬜ | — | — |
| 7.13 | Code review toàn bộ codebase | 🟠 | ⬜ | — | — |
| 7.14 | Setup CI/CD pipeline (GitHub Actions) | 🟠 | ⬜ | — | Lint → Test → Build → Deploy |
| 7.15 | Deploy MVP lên staging environment | 🔴 | ⬜ | — | — |
| 7.16 | Viết README.md cho Backend & Frontend | 🟡 | ⬜ | — | Setup guide |

---

## 📌 Phase 2 — Enhancement (Tuần 11 → Tuần 18)

---

### ⭐ 8. Good/Bad Days & Solar Terms (Tuần 11-12)

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 8.1 | Hoàn thiện thuật toán ngày tốt/xấu chi tiết | 🟠 | ⬜ | — | Theo quy tắc VN |
| 8.2 | API endpoint `GET /good-days/:year/:month` | 🟠 | ⬜ | — | — |
| 8.3 | Hiển thị "Nên" / "Không nên" cho mỗi ngày | 🟠 | ⬜ | — | — |
| 8.4 | Hiển thị 12 giờ hoàng đạo/hắc đạo | 🟠 | ⬜ | — | — |
| 8.5 | Hiển thị tiết khí hiện tại trên calendar | 🟡 | ⬜ | — | — |
| 8.6 | Filter ngày tốt theo loại việc (cưới, khai trương, ...) | 🟡 | ⬜ | — | — |

### 🔔 9. Notifications & Reminders (Tuần 13-14)

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 9.1 | Viết `Reminder` model + repository | 🟠 | ⬜ | — | — |
| 9.2 | Viết `ReminderService` - CRUD | 🟠 | ⬜ | — | — |
| 9.3 | Viết `NotificationService` - Push notifications | 🟠 | ⬜ | — | — |
| 9.4 | Background job: Scan & send reminders (cron) | 🟠 | ⬜ | — | Every minute |
| 9.5 | Email notification service (SMTP) | 🟡 | ⬜ | — | — |
| 9.6 | Frontend: Notification bell + dropdown | 🟠 | ⬜ | — | — |
| 9.7 | Frontend: Reminder settings trong task form | 🟠 | ⬜ | — | — |
| 9.8 | Browser push notification (Web Push API) | 🟡 | ⬜ | — | Service Worker |
| 9.9 | Recurring reminders (based on RRULE) | 🟡 | ⬜ | — | — |

### 🔄 10. Google Calendar Sync (Tuần 15-16)

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 10.1 | Google OAuth2 integration | 🟠 | ⬜ | — | — |
| 10.2 | Google Calendar API client | 🟠 | ⬜ | — | — |
| 10.3 | Import events từ Google Calendar | 🟠 | ⬜ | — | — |
| 10.4 | Export events sang Google Calendar | 🟠 | ⬜ | — | — |
| 10.5 | Two-way sync (webhook/polling) | 🟡 | ⬜ | — | — |
| 10.6 | Frontend: Google Calendar connect UI | 🟠 | ⬜ | — | Settings page |
| 10.7 | Conflict resolution strategy | 🟡 | ⬜ | — | — |

### 📊 11. Analytics Dashboard (Tuần 17-18)

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 11.1 | API: Task completion statistics | 🟡 | ⬜ | — | Daily/Weekly/Monthly |
| 11.2 | API: Productivity trends | 🟡 | ⬜ | — | — |
| 11.3 | API: Category distribution | 🟡 | ⬜ | — | — |
| 11.4 | Frontend: Dashboard page | 🟡 | ⬜ | — | — |
| 11.5 | Frontend: Charts (Recharts/Chart.js) | 🟡 | ⬜ | — | Bar, Pie, Line |
| 11.6 | Frontend: Weekly summary card | 🟡 | ⬜ | — | — |
| 11.7 | Performance optimization & caching | 🟡 | ⬜ | — | — |

---

## 📌 Phase 3 — Scale (Tuần 19 → Tuần 28)

---

### 📱 12. Mobile App (Tuần 19-21)

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 12.1 | Setup React Native / Flutter project | 🟡 | ⬜ | — | Đánh giá lựa chọn |
| 12.2 | Calendar view (mobile optimized) | 🟡 | ⬜ | — | — |
| 12.3 | Task management (mobile) | 🟡 | ⬜ | — | — |
| 12.4 | Note editor (mobile) | 🟡 | ⬜ | — | — |
| 12.5 | Push notifications (FCM/APNs) | 🟡 | ⬜ | — | — |
| 12.6 | Offline support (SQLite + sync) | 🟡 | ⬜ | — | — |
| 12.7 | Widget (iOS/Android) | 🟢 | ⬜ | — | — |

### 👥 13. Team/Family Calendar (Tuần 22-24)

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 13.1 | Database schema cho Teams/Groups | 🟡 | ⬜ | — | — |
| 13.2 | Team CRUD API | 🟡 | ⬜ | — | — |
| 13.3 | Shared calendars | 🟡 | ⬜ | — | — |
| 13.4 | Assign tasks cho team members | 🟡 | ⬜ | — | — |
| 13.5 | Share notes với team | 🟡 | ⬜ | — | — |
| 13.6 | Role-based permissions (owner, admin, member, viewer) | 🟡 | ⬜ | — | — |
| 13.7 | WebSocket real-time updates | 🟡 | ⬜ | — | — |
| 13.8 | Frontend: Team management UI | 🟡 | ⬜ | — | — |

### 🤖 14. AI Features (Tuần 25-26)

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 14.1 | AI: Smart scheduling suggestions | 🟢 | ⬜ | — | OpenAI/Gemini API |
| 14.2 | AI: Auto-categorize notes & tasks | 🟢 | ⬜ | — | — |
| 14.3 | AI: Summarize notes | 🟢 | ⬜ | — | — |
| 14.4 | AI: Suggest ngày tốt cho sự kiện | 🟢 | ⬜ | — | — |
| 14.5 | AI: Natural language task creation | 🟢 | ⬜ | — | "Nhắc tôi mua hoa ngày rằm" |

### 🔧 15. Monitoring & Production (Tuần 27-28)

| # | Task | Ưu tiên | Trạng thái | Assignee | Ghi chú |
|---|------|---------|------------|----------|---------|
| 15.1 | Setup Prometheus metrics | 🟡 | ⬜ | — | — |
| 15.2 | Setup Grafana dashboards | 🟡 | ⬜ | — | — |
| 15.3 | Setup Sentry error tracking | 🟡 | ⬜ | — | Backend + Frontend |
| 15.4 | Setup centralized logging (ELK/Loki) | 🟢 | ⬜ | — | — |
| 15.5 | Load testing (k6/Locust) | 🟡 | ⬜ | — | 1000 concurrent users |
| 15.6 | Security audit & penetration testing | 🟠 | ⬜ | — | — |
| 15.7 | Database backup strategy | 🟠 | ⬜ | — | Daily backup |
| 15.8 | Disaster recovery plan | 🟡 | ⬜ | — | — |
| 15.9 | Production deployment (VPS/Cloud) | 🟠 | ⬜ | — | — |
| 15.10 | SSL certificate (Let's Encrypt) | 🟠 | ⬜ | — | — |
| 15.11 | CDN setup (CloudFlare) | 🟡 | ⬜ | — | — |
| 15.12 | Domain & DNS configuration | 🟠 | ⬜ | — | — |

---

## 📊 Tổng Kết

### Thống kê tasks theo Phase

| Phase | Tổng Tasks | Critical 🔴 | High 🟠 | Medium 🟡 | Low 🟢 |
|-------|-----------|-------------|---------|-----------|--------|
| **Phase 1 - MVP** | ~120 | ~50 | ~40 | ~25 | ~5 |
| **Phase 2 - Enhancement** | ~35 | ~0 | ~18 | ~17 | ~0 |
| **Phase 3 - Scale** | ~30 | ~0 | ~5 | ~15 | ~10 |
| **Tổng** | **~185** | **~50** | **~63** | **~57** | **~15** |

### Thống kê tasks theo Module

| Module | Tasks | Phase |
|--------|-------|-------|
| Setup & Infrastructure | ~25 | 1 |
| Authentication & User | ~30 | 1 |
| Lunar Calendar Engine | ~25 | 1 |
| Task Module | ~40 | 1 |
| Note Module | ~35 | 1 |
| Calendar UI | ~30 | 1 |
| Testing & QA | ~16 | 1 |
| Good/Bad Days | ~6 | 2 |
| Notifications | ~9 | 2 |
| Google Calendar Sync | ~7 | 2 |
| Analytics | ~7 | 2 |
| Mobile App | ~7 | 3 |
| Team/Family | ~8 | 3 |
| AI Features | ~5 | 3 |
| Monitoring & Production | ~12 | 3 |

---

## 📝 Changelog

| Ngày | Thay đổi | Người thay đổi |
|------|----------|----------------|
| 03/03/2026 | Khởi tạo TODO.md | — |

---

> **💡 Lưu ý**: Cập nhật trạng thái task hàng ngày. Review và điều chỉnh ưu tiên hàng tuần trong sprint planning.
