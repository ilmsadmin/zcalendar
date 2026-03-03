package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"zplus-calendar/internal/config"
)

// TokenType represents the type of JWT token
type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

// TokenClaims represents custom JWT claims
type TokenClaims struct {
	UserID    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	TokenType TokenType `json:"token_type"`
	jwt.RegisteredClaims
}

// TokenPair contains both access and refresh tokens
type TokenPair struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// JWTManager handles JWT token operations
type JWTManager struct {
	cfg config.JWTConfig
}

// NewJWTManager creates a new JWTManager
func NewJWTManager(cfg config.JWTConfig) *JWTManager {
	return &JWTManager{cfg: cfg}
}

// GenerateTokenPair creates a new access + refresh token pair
func (m *JWTManager) GenerateTokenPair(userID uuid.UUID, email string) (*TokenPair, error) {
	accessExpiry, err := parseDuration(m.cfg.AccessExpiry)
	if err != nil {
		return nil, fmt.Errorf("invalid access expiry: %w", err)
	}

	refreshExpiry, err := parseDuration(m.cfg.RefreshExpiry)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh expiry: %w", err)
	}

	accessToken, accessExp, err := m.generateToken(userID, email, AccessToken, accessExpiry)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, _, err := m.generateToken(userID, email, RefreshToken, refreshExpiry)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    accessExp,
	}, nil
}

// GenerateAccessToken creates a new access token only
func (m *JWTManager) GenerateAccessToken(userID uuid.UUID, email string) (string, time.Time, error) {
	accessExpiry, err := parseDuration(m.cfg.AccessExpiry)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("invalid access expiry: %w", err)
	}
	return m.generateToken(userID, email, AccessToken, accessExpiry)
}

// ValidateToken validates a JWT token and returns the claims
func (m *JWTManager) ValidateToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.cfg.Secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

// ValidateRefreshToken validates a refresh token specifically
func (m *JWTManager) ValidateRefreshToken(tokenString string) (*TokenClaims, error) {
	claims, err := m.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != RefreshToken {
		return nil, fmt.Errorf("token is not a refresh token")
	}

	return claims, nil
}

// generateToken creates a signed JWT token
func (m *JWTManager) generateToken(userID uuid.UUID, email string, tokenType TokenType, expiry time.Duration) (string, time.Time, error) {
	now := time.Now()
	expiresAt := now.Add(expiry)

	claims := TokenClaims{
		UserID:    userID,
		Email:     email,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "zplus-calendar",
			Subject:   userID.String(),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(m.cfg.Secret))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, expiresAt, nil
}

// parseDuration parses a duration string like "15m", "1h", "7d"
func parseDuration(s string) (time.Duration, error) {
	// Handle "d" (days) suffix which time.ParseDuration doesn't support
	if len(s) > 0 && s[len(s)-1] == 'd' {
		var days int
		_, err := fmt.Sscanf(s, "%dd", &days)
		if err != nil {
			return 0, fmt.Errorf("invalid duration: %s", s)
		}
		return time.Duration(days) * 24 * time.Hour, nil
	}
	return time.ParseDuration(s)
}
