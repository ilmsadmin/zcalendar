package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"zplus-calendar/internal/config"
	"zplus-calendar/internal/database"
	"zplus-calendar/internal/model"
	pkglogger "zplus-calendar/pkg/logger"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Init logger
	zapLogger, err := pkglogger.NewLogger(cfg.App.Env)
	if err != nil {
		log.Fatal("Failed to init logger:", err)
	}
	defer zapLogger.Sync()

	// Connect to Postgres
	db, err := database.NewPostgres(cfg.Database.Postgres)
	if err != nil {
		log.Fatal("Failed to connect to Postgres:", err)
	}

	// Auto migrate
	if err := database.AutoMigrate(db, zapLogger); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	fmt.Println("🌱 Seeding database...")

	// ── Seed Admin User ────────────────────────────────────────────
	hashedPw, _ := bcrypt.GenerateFromPassword([]byte("admin123456"), 12)
	admin := model.User{
		Email:        "admin@zplus-calendar.com",
		Username:     "admin",
		PasswordHash: string(hashedPw),
		FullName:     "Admin User",
		Role:         "admin",
		IsActive:     true,
		IsVerified:   true,
	}

	result := db.Where("email = ?", admin.Email).FirstOrCreate(&admin)
	if result.Error != nil {
		log.Fatal("Failed to seed admin user:", result.Error)
	}
	fmt.Printf("  ✅ Admin user: %s (ID: %s)\n", admin.Email, admin.ID)

	// ── Seed Demo User ─────────────────────────────────────────────
	demoHashedPw, _ := bcrypt.GenerateFromPassword([]byte("demo123456"), 12)
	demo := model.User{
		Email:        "demo@zplus-calendar.com",
		Username:     "demo",
		PasswordHash: string(demoHashedPw),
		FullName:     "Demo User",
		Role:         "user",
		IsActive:     true,
	}

	result = db.Where("email = ?", demo.Email).FirstOrCreate(&demo)
	if result.Error != nil {
		log.Fatal("Failed to seed demo user:", result.Error)
	}
	fmt.Printf("  ✅ Demo user: %s (ID: %s)\n", demo.Email, demo.ID)

	// ── Seed Categories ────────────────────────────────────────────
	categories := []model.Category{
		{UserID: demo.ID, Name: "Công việc", Color: "#3B82F6", Icon: "briefcase", Type: "task", SortOrder: 1},
		{UserID: demo.ID, Name: "Cá nhân", Color: "#10B981", Icon: "user", Type: "both", SortOrder: 2},
		{UserID: demo.ID, Name: "Gia đình", Color: "#F59E0B", Icon: "home", Type: "both", SortOrder: 3},
		{UserID: demo.ID, Name: "Sức khoẻ", Color: "#EF4444", Icon: "heart", Type: "task", SortOrder: 4},
		{UserID: demo.ID, Name: "Học tập", Color: "#8B5CF6", Icon: "book", Type: "both", SortOrder: 5},
		{UserID: demo.ID, Name: "Nhật ký", Color: "#EC4899", Icon: "edit", Type: "note", SortOrder: 6},
	}

	for _, cat := range categories {
		db.Where("user_id = ? AND name = ?", cat.UserID, cat.Name).FirstOrCreate(&cat)
	}
	fmt.Printf("  ✅ Seeded %d categories for demo user\n", len(categories))

	// ── Seed Tags ──────────────────────────────────────────────────
	tags := []model.Tag{
		{UserID: demo.ID, Name: "Quan trọng", Color: "#EF4444"},
		{UserID: demo.ID, Name: "Gấp", Color: "#F59E0B"},
		{UserID: demo.ID, Name: "Âm lịch", Color: "#8B5CF6"},
		{UserID: demo.ID, Name: "Lễ tết", Color: "#EC4899"},
		{UserID: demo.ID, Name: "Họp", Color: "#3B82F6"},
		{UserID: demo.ID, Name: "Deadline", Color: "#F97316"},
	}

	for _, tag := range tags {
		db.Where("user_id = ? AND name = ?", tag.UserID, tag.Name).FirstOrCreate(&tag)
	}
	fmt.Printf("  ✅ Seeded %d tags for demo user\n", len(tags))

	// ── Seed Vietnamese Holidays ───────────────────────────────────
	holidays := []model.VietnameseHoliday{
		{Name: "Tết Dương lịch", Day: 1, Month: 1, IsLunar: false, Type: "public_holiday", IsDayOff: true, DaysOffCount: 1},
		{Name: "Tết Nguyên Đán", Day: 1, Month: 1, IsLunar: true, Type: "public_holiday", IsDayOff: true, DaysOffCount: 5},
		{Name: "Giỗ Tổ Hùng Vương", Day: 10, Month: 3, IsLunar: true, Type: "public_holiday", IsDayOff: true, DaysOffCount: 1},
		{Name: "Ngày Giải phóng miền Nam", Day: 30, Month: 4, IsLunar: false, Type: "public_holiday", IsDayOff: true, DaysOffCount: 1},
		{Name: "Ngày Quốc tế Lao động", Day: 1, Month: 5, IsLunar: false, Type: "public_holiday", IsDayOff: true, DaysOffCount: 1},
		{Name: "Quốc khánh", Day: 2, Month: 9, IsLunar: false, Type: "public_holiday", IsDayOff: true, DaysOffCount: 2},
		{Name: "Rằm tháng Giêng", Day: 15, Month: 1, IsLunar: true, Type: "cultural", IsDayOff: false},
		{Name: "Tết Đoan Ngọ", Day: 5, Month: 5, IsLunar: true, Type: "traditional", IsDayOff: false},
		{Name: "Vu Lan", Day: 15, Month: 7, IsLunar: true, Type: "cultural", IsDayOff: false},
		{Name: "Tết Trung Thu", Day: 15, Month: 8, IsLunar: true, Type: "traditional", IsDayOff: false},
		{Name: "Ngày Phụ nữ Việt Nam", Day: 20, Month: 10, IsLunar: false, Type: "cultural", IsDayOff: false},
		{Name: "Ngày Nhà giáo Việt Nam", Day: 20, Month: 11, IsLunar: false, Type: "cultural", IsDayOff: false},
	}

	for _, h := range holidays {
		db.Where("name = ? AND day = ? AND month = ?", h.Name, h.Day, h.Month).FirstOrCreate(&h)
	}
	fmt.Printf("  ✅ Seeded %d Vietnamese holidays\n", len(holidays))

	fmt.Println("\n🎉 Seeding completed!")
}
