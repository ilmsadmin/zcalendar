package dto

import (
	"time"

	"github.com/google/uuid"
)

// ─── Auth Request DTOs ──────────────────────────────────────────────

// RegisterRequest is the request body for user registration
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Username string `json:"username" validate:"required,min=3,max=50,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=128"`
	FullName string `json:"full_name" validate:"omitempty,max=100"`
}

// LoginRequest is the request body for user login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RefreshTokenRequest is the request body for token refresh
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// ForgotPasswordRequest is the request body for forgot password
type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// ResetPasswordRequest is the request body for reset password
type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=128"`
}

// ChangePasswordRequest is the request body for change password
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=128"`
}

// ─── Auth Response DTOs ─────────────────────────────────────────────

// AuthResponse is the response for login/register
type AuthResponse struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresAt    time.Time    `json:"expires_at"`
}

// TokenResponse is the response for token refresh
type TokenResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// UserResponse is the public user data
type UserResponse struct {
	ID          uuid.UUID                `json:"id"`
	Email       string                   `json:"email"`
	Username    string                   `json:"username"`
	FullName    string                   `json:"full_name,omitempty"`
	AvatarURL   string                   `json:"avatar_url,omitempty"`
	Role        string                   `json:"role"`
	Preferences *UserPreferencesResponse `json:"preferences,omitempty"`
	CreatedAt   time.Time                `json:"created_at"`
}

// UserPreferencesResponse contains user preferences
type UserPreferencesResponse struct {
	Theme        string `json:"theme"`
	Locale       string `json:"locale"`
	CalendarView string `json:"calendar_view"`
	WeekStartDay int    `json:"week_start_day"`
}

// ─── User Profile DTOs ──────────────────────────────────────────────

// UpdateProfileRequest is the request body for updating user profile
type UpdateProfileRequest struct {
	FullName string `json:"full_name" validate:"omitempty,max=100"`
	Username string `json:"username" validate:"omitempty,min=3,max=50,alphanum"`
}

// UpdatePreferencesRequest is the request body for updating preferences
type UpdatePreferencesRequest struct {
	Theme        string `json:"theme" validate:"omitempty,oneof=light dark system"`
	Locale       string `json:"locale" validate:"omitempty,oneof=vi en"`
	CalendarView string `json:"calendar_view" validate:"omitempty,oneof=month week day"`
	WeekStartDay *int   `json:"week_start_day" validate:"omitempty,min=0,max=6"`
}
