# 📅 Zplus Calendar

> Ứng dụng lịch thông minh hỗ trợ **Dương lịch** và **Âm lịch Việt Nam**, tích hợp quản lý công việc (Task) và ghi chú (Note).

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)
![Next.js](https://img.shields.io/badge/Next.js-16-black?logo=next.js)
![React](https://img.shields.io/badge/React-19-61DAFB?logo=react&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?logo=postgresql&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-7-47A248?logo=mongodb&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-7-DC382D?logo=redis&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)

---

## ✨ Tính năng chính

### 📆 Lịch Âm - Dương
- Xem lịch theo **tháng / tuần / ngày** với thông tin âm lịch song song
- Chuyển đổi **Dương lịch ↔ Âm lịch** chính xác
- Nhảy đến ngày bất kỳ (hỗ trợ cả nhập ngày âm lịch)
- Hiển thị **Can Chi** (ngày, tháng, năm), **tiết khí**, **tuổi**

### 🔮 Ngày tốt / xấu
- Đánh giá chất lượng ngày: Hoàng Đạo, Hắc Đạo
- Danh sách việc **nên làm** và **không nên làm**
- **12 giờ hoàng đạo** trong ngày

### 🎉 Ngày lễ Việt Nam
- Ngày lễ Dương lịch (Quốc khánh, Quốc tế Lao động…)
- Ngày lễ Âm lịch (Tết Nguyên Đán, Rằm tháng Giêng…)
- Phân biệt ngày nghỉ chính thức và ngày văn hoá

### ✅ Quản lý công việc (Task)
- Tạo, sửa, xoá task với độ ưu tiên (Thấp / TB / Cao / Gấp)
- Trạng thái: Chờ → Đang làm → Hoàn thành
- **Tạo nhanh inline** ngay từ chế độ xem ngày
- Toggle hoàn thành bằng 1 click
- Lọc & tìm kiếm task

### 📝 Ghi chú (Note)
- Ghi chú theo ngày, hỗ trợ chọn màu
- **Tạo nhanh inline** ngay từ chế độ xem ngày
- Ghim ghi chú quan trọng
- Tìm kiếm ghi chú

### ⚙️ Cài đặt
- Quản lý hồ sơ (họ tên, username)
- Đổi mật khẩu
- Chế độ giao diện: Sáng / Tối / Hệ thống
- Chế độ xem mặc định (tháng / tuần / ngày)
- Xoá tài khoản

---

## 🏗️ Kiến trúc

```
┌──────────────────────────────┐
│     Frontend (Next.js 16)    │
│  React 19 · Tailwind · SWR  │
└──────────────┬───────────────┘
               │ REST API
               ▼
┌──────────────────────────────┐
│     Backend (Go / Fiber)     │
│  JWT Auth · Lunar Engine     │
└──┬──────────┬──────────┬─────┘
   │          │          │
   ▼          ▼          ▼
┌──────┐ ┌──────┐ ┌──────────┐
│Postgres│ │MongoDB│ │  Redis   │
│Users,  │ │Notes  │ │Sessions, │
│Tasks   │ │       │ │Cache     │
└────────┘ └───────┘ └──────────┘
```

---

## 🛠️ Tech Stack

| Layer | Technology |
|-------|-----------|
| **Frontend** | Next.js 16 (App Router), React 19, TypeScript, Tailwind CSS v4, shadcn/ui |
| **State** | Zustand, TanStack React Query |
| **Backend** | Go 1.24, Fiber v2, GORM |
| **Auth** | JWT (access + refresh token), bcrypt |
| **Database** | PostgreSQL 16 (users, tasks, categories, tags) |
| **NoSQL** | MongoDB 7 (notes) |
| **Cache** | Redis 7 (sessions, rate limiting, token blacklist) |
| **Infra** | Docker Compose, Makefile |

---

## 📂 Cấu trúc dự án

```
Zplus_Calendar/
├── backend/                    # Go backend
│   ├── cmd/
│   │   ├── server/main.go     # Entry point
│   │   └── seed/main.go       # Database seeder
│   ├── configs/config.yaml    # App configuration
│   ├── internal/
│   │   ├── common/            # Shared response, pagination, errors
│   │   ├── config/            # Config loader (Viper)
│   │   ├── database/          # DB connections (Postgres, Mongo, Redis)
│   │   ├── dto/               # Request/Response DTOs
│   │   ├── handler/           # HTTP handlers
│   │   ├── lunar/             # ⭐ Lunar calendar engine
│   │   ├── middleware/        # Auth, Logger, Rate limit
│   │   ├── model/             # GORM models
│   │   ├── repository/        # Data access layer
│   │   ├── router/            # Route registration
│   │   └── service/           # Business logic
│   ├── migrations/sql/        # SQL migrations
│   └── pkg/                   # Shared packages (JWT, password)
├── frontend/                   # Next.js frontend
│   └── src/
│       ├── app/               # App Router pages
│       │   ├── (auth)/        # Login, Register
│       │   └── (dashboard)/   # Calendar, Tasks, Notes, Settings
│       ├── components/        # UI components
│       │   ├── calendar/      # Calendar views, date picker
│       │   ├── tasks/         # Task list, form, filters
│       │   ├── notes/         # Note card, form
│       │   ├── layout/        # Header, Sidebar, Mobile nav
│       │   └── ui/            # shadcn/ui primitives
│       ├── hooks/             # React Query hooks
│       ├── lib/api/           # API client (Axios)
│       ├── stores/            # Zustand stores
│       └── types/             # TypeScript types
├── docs/                       # Documentation
│   └── SYSTEM_DESIGN.md       # System design document
├── docker-compose.yml          # PostgreSQL, MongoDB, Redis
├── Makefile                    # Build, test, Docker commands
└── .env.example                # Environment variables template
```

---

## 🚀 Bắt đầu

### Yêu cầu

- **Go** ≥ 1.24
- **Node.js** ≥ 20
- **Docker** & **Docker Compose** (cho databases)

### 1. Clone repo

```bash
git clone https://github.com/ilmsadmin/zcalendar.git
cd zcalendar
```

### 2. Khởi động databases

```bash
make docker-db
# Hoặc:
docker compose up -d postgres mongo redis
```

Ports mặc định:
| Service | Port |
|---------|------|
| PostgreSQL | `5433` |
| MongoDB | `27018` |
| Redis | `6380` |

### 3. Cấu hình

```bash
cp .env.example .env
# Chỉnh sửa .env nếu cần
```

Backend đọc config từ `backend/configs/config.yaml`. Các giá trị mặc định phù hợp cho development.

### 4. Chạy Backend

```bash
# Chạy trực tiếp
cd backend && go run cmd/server/main.go

# Hoặc dùng hot-reload (cần Air)
make dev

# Hoặc build rồi chạy
make run
```

Backend sẽ chạy tại **http://localhost:8080**

### 5. Seed dữ liệu (tuỳ chọn)

```bash
make db-seed
# Tạo: admin user, demo user, categories, tags, holidays
```

### 6. Chạy Frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend sẽ chạy tại **http://localhost:3000**

### 7. Mở trình duyệt

Truy cập **http://localhost:3000** → Đăng ký tài khoản mới hoặc đăng nhập.

---

## 📡 API Endpoints

Base URL: `http://localhost:8080/api/v1`

### 🔓 Public

| Method | Endpoint | Mô tả |
|--------|----------|-------|
| `POST` | `/auth/register` | Đăng ký |
| `POST` | `/auth/login` | Đăng nhập |
| `POST` | `/auth/refresh` | Refresh token |
| `GET` | `/calendar/today` | Thông tin hôm nay |
| `GET` | `/calendar/day?day=&month=&year=` | Thông tin 1 ngày |
| `GET` | `/calendar/month/:year/:month` | Lịch tháng |
| `GET` | `/calendar/convert/solar-to-lunar` | Dương → Âm |
| `GET` | `/calendar/convert/lunar-to-solar` | Âm → Dương |
| `GET` | `/calendar/holidays/:year` | Ngày lễ theo năm |

### 🔐 Protected (Bearer Token)

| Method | Endpoint | Mô tả |
|--------|----------|-------|
| `POST` | `/auth/logout` | Đăng xuất |
| `POST` | `/auth/change-password` | Đổi mật khẩu |
| `GET` | `/auth/profile` | Lấy hồ sơ |
| `PUT` | `/auth/profile` | Cập nhật hồ sơ |
| `PUT` | `/auth/preferences` | Cập nhật tuỳ chọn |
| `DELETE` | `/auth/account` | Xoá tài khoản |
| | | |
| `POST` | `/tasks` | Tạo task |
| `GET` | `/tasks` | Danh sách task |
| `GET` | `/tasks/search` | Tìm kiếm task |
| `GET` | `/tasks/by-date/:date` | Task theo ngày |
| `GET` | `/tasks/:id` | Chi tiết task |
| `PUT` | `/tasks/:id` | Cập nhật task |
| `DELETE` | `/tasks/:id` | Xoá task |
| `PATCH` | `/tasks/:id/status` | Đổi trạng thái |
| | | |
| `POST` | `/notes` | Tạo ghi chú |
| `GET` | `/notes` | Danh sách ghi chú |
| `GET` | `/notes/search` | Tìm kiếm ghi chú |
| `GET` | `/notes/by-date/:date` | Ghi chú theo ngày |
| `GET` | `/notes/:id` | Chi tiết ghi chú |
| `PUT` | `/notes/:id` | Cập nhật ghi chú |
| `DELETE` | `/notes/:id` | Xoá ghi chú |
| | | |
| `POST` | `/categories` | Tạo danh mục |
| `GET` | `/categories` | Danh sách danh mục |
| `PUT` | `/categories/:id` | Cập nhật danh mục |
| `DELETE` | `/categories/:id` | Xoá danh mục |
| | | |
| `POST` | `/tags` | Tạo tag |
| `GET` | `/tags` | Danh sách tag |
| `PUT` | `/tags/:id` | Cập nhật tag |
| `DELETE` | `/tags/:id` | Xoá tag |

---

## 🧪 Testing

```bash
# Chạy tất cả tests
make test

# Chỉ test Lunar engine
make test-lunar

# Test với coverage report
make test-cover

# Test race condition
make test-race

# Benchmark lunar engine
make bench
```

---

## 🔧 Makefile Commands

```bash
make help            # Hiện tất cả commands

# Development
make dev             # Backend hot-reload (Air)
make dev-frontend    # Frontend dev server
make dev-all         # Chạy cả hai

# Build
make build           # Build Go binary
make build-frontend  # Build Next.js
make build-all       # Build tất cả

# Docker
make docker-db       # Chỉ start databases
make docker-up       # Start all services
make docker-down     # Stop all
make docker-clean    # Stop + xoá volumes

# Database
make db-seed         # Seed data
make db-reset        # Reset database
make db-shell        # PostgreSQL shell
make mongo-shell     # MongoDB shell
make redis-cli       # Redis CLI

# Quality
make fmt             # Format code
make vet             # Go vet
make lint            # Lint
make check           # Tất cả checks
```

---

## 🌙 Lunar Engine

Thuật toán chuyển đổi Âm - Dương lịch được implement trong `backend/internal/lunar/`:

| File | Chức năng |
|------|-----------|
| `converter.go` | Chuyển đổi Solar ↔ Lunar |
| `astronomy.go` | Tính toán thiên văn (New Moon, Sun Longitude) |
| `can_chi.go` | Can Chi ngày / tháng / năm, con giáp |
| `good_bad_day.go` | Ngày tốt/xấu, giờ hoàng đạo |
| `holiday.go` | Ngày lễ Việt Nam (Dương lịch + Âm lịch) |
| `solar_term.go` | 24 tiết khí |

Thuật toán dựa trên **Astronomical Algorithms** của Jean Meeus, múi giờ GMT+7 (Việt Nam).

---

## 📸 Screenshots

> _Coming soon_

---

## 📄 Tài liệu

- [System Design Document](./docs/SYSTEM_DESIGN.md) — Thiết kế hệ thống chi tiết
- [TODO & Roadmap](./TODO.md) — Tiến độ và kế hoạch phát triển

---

## 🤝 Đóng góp

1. Fork repo
2. Tạo branch: `git checkout -b feature/ten-tinh-nang`
3. Commit: `git commit -m "feat: Mô tả tính năng"`
4. Push: `git push origin feature/ten-tinh-nang`
5. Tạo Pull Request

---

## 📜 License

[MIT](LICENSE)

---

<p align="center">
  Made with ❤️ by <strong>Zplus Team</strong>
</p>
