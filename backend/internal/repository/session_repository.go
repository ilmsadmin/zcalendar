package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// SessionData represents a user session stored in Redis
type SessionData struct {
	UserID       uuid.UUID `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	IP           string    `json:"ip"`
	CreatedAt    time.Time `json:"created_at"`
}

// SessionRepository defines the interface for session operations
type SessionRepository interface {
	Create(ctx context.Context, sessionID string, data *SessionData, ttl time.Duration) error
	Get(ctx context.Context, sessionID string) (*SessionData, error)
	Delete(ctx context.Context, sessionID string) error
	DeleteAllForUser(ctx context.Context, userID uuid.UUID) error
	BlacklistToken(ctx context.Context, tokenID string, ttl time.Duration) error
	IsTokenBlacklisted(ctx context.Context, tokenID string) (bool, error)
}

type sessionRepository struct {
	redis *redis.Client
}

// NewSessionRepository creates a new SessionRepository
func NewSessionRepository(redisClient *redis.Client) SessionRepository {
	return &sessionRepository{redis: redisClient}
}

func sessionKey(sessionID string) string {
	return fmt.Sprintf("session:%s", sessionID)
}

func userSessionsKey(userID uuid.UUID) string {
	return fmt.Sprintf("user_sessions:%s", userID.String())
}

func blacklistKey(tokenID string) string {
	return fmt.Sprintf("token_blacklist:%s", tokenID)
}

func (r *sessionRepository) Create(ctx context.Context, sessionID string, data *SessionData, ttl time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal session data: %w", err)
	}

	pipe := r.redis.Pipeline()

	// Store session data
	pipe.Set(ctx, sessionKey(sessionID), jsonData, ttl)

	// Add to user's session set
	pipe.SAdd(ctx, userSessionsKey(data.UserID), sessionID)
	pipe.Expire(ctx, userSessionsKey(data.UserID), ttl)

	if _, err := pipe.Exec(ctx); err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	return nil
}

func (r *sessionRepository) Get(ctx context.Context, sessionID string) (*SessionData, error) {
	jsonData, err := r.redis.Get(ctx, sessionKey(sessionID)).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	var data SessionData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session data: %w", err)
	}

	return &data, nil
}

func (r *sessionRepository) Delete(ctx context.Context, sessionID string) error {
	// Get session to find user ID
	session, err := r.Get(ctx, sessionID)
	if err != nil {
		return err
	}

	pipe := r.redis.Pipeline()
	pipe.Del(ctx, sessionKey(sessionID))

	if session != nil {
		pipe.SRem(ctx, userSessionsKey(session.UserID), sessionID)
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	return nil
}

func (r *sessionRepository) DeleteAllForUser(ctx context.Context, userID uuid.UUID) error {
	// Get all session IDs for the user
	sessionIDs, err := r.redis.SMembers(ctx, userSessionsKey(userID)).Result()
	if err != nil {
		return fmt.Errorf("failed to get user sessions: %w", err)
	}

	if len(sessionIDs) == 0 {
		return nil
	}

	pipe := r.redis.Pipeline()
	for _, sid := range sessionIDs {
		pipe.Del(ctx, sessionKey(sid))
	}
	pipe.Del(ctx, userSessionsKey(userID))

	if _, err := pipe.Exec(ctx); err != nil {
		return fmt.Errorf("failed to delete all sessions: %w", err)
	}

	return nil
}

func (r *sessionRepository) BlacklistToken(ctx context.Context, tokenID string, ttl time.Duration) error {
	if err := r.redis.Set(ctx, blacklistKey(tokenID), "1", ttl).Err(); err != nil {
		return fmt.Errorf("failed to blacklist token: %w", err)
	}
	return nil
}

func (r *sessionRepository) IsTokenBlacklisted(ctx context.Context, tokenID string) (bool, error) {
	exists, err := r.redis.Exists(ctx, blacklistKey(tokenID)).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check token blacklist: %w", err)
	}
	return exists > 0, nil
}
