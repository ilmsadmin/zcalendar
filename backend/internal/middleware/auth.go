package middleware

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"zplus-calendar/internal/common"
	"zplus-calendar/internal/config"
)

type AuthMiddleware struct {
	jwtConfig config.JWTConfig
	redis     *redis.Client
	logger    *zap.Logger
}

func NewAuthMiddleware(jwtCfg config.JWTConfig, redis *redis.Client, logger *zap.Logger) *AuthMiddleware {
	return &AuthMiddleware{
		jwtConfig: jwtCfg,
		redis:     redis,
		logger:    logger,
	}
}

// Protected requires a valid JWT token
func (m *AuthMiddleware) Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return common.Unauthorized(c, "Token không được cung cấp")
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return common.Unauthorized(c, "Định dạng token không hợp lệ")
		}

		tokenString := parts[1]

		// Check if token is blacklisted
		ctx := context.Background()
		blacklisted, _ := m.redis.Get(ctx, "blacklist:"+tokenString).Result()
		if blacklisted != "" {
			return common.Unauthorized(c, "Token đã bị thu hồi")
		}

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Phương thức ký không hợp lệ")
			}
			return []byte(m.jwtConfig.Secret), nil
		})

		if err != nil || !token.Valid {
			return common.Unauthorized(c, "Token không hợp lệ hoặc đã hết hạn")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return common.Unauthorized(c, "Không thể đọc claims từ token")
		}

		// Store user info in context
		c.Locals("user_id", claims["sub"])
		c.Locals("email", claims["email"])
		c.Locals("username", claims["username"])

		return c.Next()
	}
}

// RateLimit limits requests per user
func RateLimit(redis *redis.Client, maxRequests int, window time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := "ratelimit:" + c.IP() + ":" + c.Path()

		ctx := context.Background()
		current, err := redis.Incr(ctx, key).Result()
		if err != nil {
			// If Redis is down, allow the request
			return c.Next()
		}

		if current == 1 {
			redis.Expire(ctx, key, window)
		}

		if current > int64(maxRequests) {
			return common.Error(c, fiber.StatusTooManyRequests, "RATE_LIMIT_EXCEEDED", "Quá nhiều yêu cầu. Vui lòng thử lại sau.")
		}

		c.Set("X-RateLimit-Limit", string(rune(maxRequests)))
		c.Set("X-RateLimit-Remaining", string(rune(maxRequests-int(current))))

		return c.Next()
	}
}
