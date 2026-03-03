package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"zplus-calendar/internal/config"
	"zplus-calendar/internal/database"
	"zplus-calendar/internal/handler"
	"zplus-calendar/internal/middleware"
	"zplus-calendar/internal/repository"
	"zplus-calendar/internal/service"
	"zplus-calendar/pkg/utils"
)

// Dependencies holds all dependencies for route handlers
type Dependencies struct {
	Config *config.Config
	Logger *zap.Logger
	PgDB   *gorm.DB
	Mongo  *database.MongoDB
	Redis  *redis.Client
}

// Setup configures all routes for the application
func Setup(app *fiber.App, deps *Dependencies) {
	// ── Middleware ──────────────────────────────────────────────────
	authMw := middleware.NewAuthMiddleware(deps.Config.JWT, deps.Redis, deps.Logger)
	rateLimiter := middleware.RateLimit(deps.Redis, 60, 60)

	// ── Repositories ───────────────────────────────────────────────
	userRepo := repository.NewUserRepository(deps.PgDB)
	sessionRepo := repository.NewSessionRepository(deps.Redis)
	taskRepo := repository.NewTaskRepository(deps.PgDB)
	categoryRepo := repository.NewCategoryRepository(deps.PgDB)
	tagRepo := repository.NewTagRepository(deps.PgDB)
	noteRepo := repository.NewNoteRepository(deps.Mongo)

	// ── Utilities ──────────────────────────────────────────────────
	jwtManager := utils.NewJWTManager(deps.Config.JWT)

	// ── Services ───────────────────────────────────────────────────
	authService := service.NewAuthService(userRepo, sessionRepo, jwtManager, deps.Logger)
	calendarService := service.NewCalendarService(deps.Logger)
	taskService := service.NewTaskService(taskRepo, deps.Logger)
	categoryService := service.NewCategoryService(categoryRepo, deps.Logger)
	tagService := service.NewTagService(tagRepo, deps.Logger)
	noteService := service.NewNoteService(noteRepo, deps.Logger)

	// ── Handlers ───────────────────────────────────────────────────
	authHandler := handler.NewAuthHandler(authService, deps.Logger)
	calendarHandler := handler.NewCalendarHandler(calendarService, deps.Logger)
	taskHandler := handler.NewTaskHandler(taskService, deps.Logger)
	categoryHandler := handler.NewCategoryHandler(categoryService, deps.Logger)
	tagHandler := handler.NewTagHandler(tagService, deps.Logger)
	noteHandler := handler.NewNoteHandler(noteService, deps.Logger)

	// ── API v1 ─────────────────────────────────────────────────────
	api := app.Group("/api/v1")
	api.Use(rateLimiter)

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/refresh", authHandler.RefreshToken)
	auth.Post("/logout", authMw.Protected(), authHandler.Logout)
	auth.Post("/change-password", authMw.Protected(), authHandler.ChangePassword)
	auth.Get("/profile", authMw.Protected(), authHandler.GetProfile)
	auth.Put("/profile", authMw.Protected(), authHandler.UpdateProfile)
	auth.Put("/preferences", authMw.Protected(), authHandler.UpdatePreferences)
	auth.Delete("/account", authMw.Protected(), authHandler.DeleteAccount)

	// Calendar routes (public)
	calendar := api.Group("/calendar")
	calendar.Get("/today", calendarHandler.GetTodayInfo)
	calendar.Get("/day", calendarHandler.GetDayInfo)
	calendar.Get("/month/:year/:month", calendarHandler.GetMonthInfo)
	calendar.Get("/convert/solar-to-lunar", calendarHandler.SolarToLunar)
	calendar.Get("/convert/lunar-to-solar", calendarHandler.LunarToSolar)
	calendar.Get("/holidays/:year", calendarHandler.GetHolidays)

	// Task routes (protected)
	tasks := api.Group("/tasks", authMw.Protected())
	tasks.Post("/", taskHandler.Create)
	tasks.Get("/", taskHandler.List)
	tasks.Get("/search", taskHandler.Search)
	tasks.Get("/by-date/:date", taskHandler.GetByDate)
	tasks.Get("/:id", taskHandler.Get)
	tasks.Put("/:id", taskHandler.Update)
	tasks.Delete("/:id", taskHandler.Delete)
	tasks.Patch("/:id/status", taskHandler.UpdateStatus)

	// Note routes (protected)
	notes := api.Group("/notes", authMw.Protected())
	notes.Post("/", noteHandler.Create)
	notes.Get("/", noteHandler.List)
	notes.Get("/search", noteHandler.Search)
	notes.Get("/by-date/:date", noteHandler.GetByDate)
	notes.Get("/:id", noteHandler.Get)
	notes.Put("/:id", noteHandler.Update)
	notes.Delete("/:id", noteHandler.Delete)

	// Category routes (protected)
	categories := api.Group("/categories", authMw.Protected())
	categories.Post("/", categoryHandler.Create)
	categories.Get("/", categoryHandler.List)
	categories.Put("/:id", categoryHandler.Update)
	categories.Delete("/:id", categoryHandler.Delete)

	// Tag routes (protected)
	tags := api.Group("/tags", authMw.Protected())
	tags.Post("/", tagHandler.Create)
	tags.Get("/", tagHandler.List)
	tags.Put("/:id", tagHandler.Update)
	tags.Delete("/:id", tagHandler.Delete)
}
