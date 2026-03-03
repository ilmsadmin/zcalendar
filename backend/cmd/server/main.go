package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"

	"zplus-calendar/internal/config"
	"zplus-calendar/internal/database"
	"zplus-calendar/internal/middleware"
	"zplus-calendar/internal/router"
	"zplus-calendar/pkg/logger"
)

// @title           Zplus Calendar API
// @version         1.0.0
// @description     Vietnamese Lunar Calendar with Notes & Tasks
// @host            localhost:8080
// @BasePath        /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	zapLog, err := logger.NewLogger(cfg.App.Env)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer zapLog.Sync()

	zapLog.Info("Starting Zplus Calendar API",
		zap.String("env", cfg.App.Env),
		zap.String("port", cfg.App.Port),
	)

	// Connect databases
	pgDB, err := database.NewPostgres(cfg.Database.Postgres)
	if err != nil {
		zapLog.Fatal("Failed to connect PostgreSQL", zap.Error(err))
	}
	zapLog.Info("PostgreSQL connected")

	// Auto migrate database schema
	if err := database.AutoMigrate(pgDB, zapLog); err != nil {
		zapLog.Fatal("Failed to run auto migrations", zap.Error(err))
	}

	mongoDB, err := database.NewMongo(cfg.Database.Mongo)
	if err != nil {
		zapLog.Fatal("Failed to connect MongoDB", zap.Error(err))
	}
	defer mongoDB.Disconnect()
	zapLog.Info("MongoDB connected")

	redisClient, err := database.NewRedis(cfg.Database.Redis)
	if err != nil {
		zapLog.Fatal("Failed to connect Redis", zap.Error(err))
	}
	defer redisClient.Close()
	zapLog.Info("Redis connected")

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      cfg.App.Name,
		ErrorHandler: middleware.ErrorHandler(zapLog),
		BodyLimit:    cfg.Upload.MaxSize,
	})

	// Global middleware
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORS.AllowOrigins,
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))
	app.Use(middleware.RequestLogger(zapLog))

	// Setup routes
	router.Setup(app, &router.Dependencies{
		Config: cfg,
		Logger: zapLog,
		PgDB:   pgDB,
		Mongo:  mongoDB,
		Redis:  redisClient,
	})

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": cfg.App.Name,
			"version": "1.0.0",
		})
	})

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		addr := fmt.Sprintf(":%s", cfg.App.Port)
		if err := app.Listen(addr); err != nil {
			zapLog.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	zapLog.Info("Server started", zap.String("port", cfg.App.Port))

	<-quit
	zapLog.Info("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		zapLog.Fatal("Server forced to shutdown", zap.Error(err))
	}

	zapLog.Info("Server exited")
}
