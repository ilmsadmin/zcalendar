package database

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"zplus-calendar/internal/config"
	"zplus-calendar/internal/model"
)

func NewPostgres(cfg config.PostgresConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect postgres: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.MaxConnections)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func AutoMigrate(db *gorm.DB, logger *zap.Logger) error {
	logger.Info("Running auto migrations...")

	err := db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Tag{},
		&model.Task{},
		&model.TaskTag{},
		&model.Reminder{},
		&model.VietnameseHoliday{},
		&model.UserSession{},
	)
	if err != nil {
		return fmt.Errorf("auto migration failed: %w", err)
	}

	logger.Info("Auto migrations completed successfully")
	return nil
}
