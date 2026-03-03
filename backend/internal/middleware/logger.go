package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// RequestLogger logs each incoming request
func RequestLogger(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		// Log after response
		duration := time.Since(start)
		statusCode := c.Response().StatusCode()

		fields := []zap.Field{
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", statusCode),
			zap.Duration("duration", duration),
			zap.String("ip", c.IP()),
			zap.String("user_agent", c.Get("User-Agent")),
		}

		if queryStr := c.Request().URI().QueryString(); len(queryStr) > 0 {
			fields = append(fields, zap.String("query", string(queryStr)))
		}

		if statusCode >= 500 {
			logger.Error("Request failed", fields...)
		} else if statusCode >= 400 {
			logger.Warn("Request warning", fields...)
		} else {
			logger.Info("Request completed", fields...)
		}

		return err
	}
}

// ErrorHandler is the global error handler for Fiber
func ErrorHandler(logger *zap.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		logger.Error("Unhandled error",
			zap.Error(err),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
		)

		return c.Status(code).JSON(fiber.Map{
			"status": "error",
			"error": fiber.Map{
				"code":    fmt.Sprintf("HTTP_%d", code),
				"message": err.Error(),
			},
		})
	}
}
