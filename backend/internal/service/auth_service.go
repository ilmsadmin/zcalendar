package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"zplus-calendar/internal/common"
	"zplus-calendar/internal/dto"
	"zplus-calendar/internal/model"
	"zplus-calendar/internal/repository"
	"zplus-calendar/pkg/utils"
)

// AuthService defines the interface for authentication operations
type AuthService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest, userAgent, ip string) (*dto.AuthResponse, error)
	Logout(ctx context.Context, userID uuid.UUID, tokenID string) error
	RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.TokenResponse, error)
	ChangePassword(ctx context.Context, userID uuid.UUID, req *dto.ChangePasswordRequest) error
	GetProfile(ctx context.Context, userID uuid.UUID) (*dto.UserResponse, error)
	UpdateProfile(ctx context.Context, userID uuid.UUID, req *dto.UpdateProfileRequest) (*dto.UserResponse, error)
	UpdatePreferences(ctx context.Context, userID uuid.UUID, req *dto.UpdatePreferencesRequest) (*dto.UserResponse, error)
	DeleteAccount(ctx context.Context, userID uuid.UUID, password string) error
}

type authService struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	jwtManager  *utils.JWTManager
	logger      *zap.Logger
}

// NewAuthService creates a new AuthService
func NewAuthService(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	jwtManager *utils.JWTManager,
	logger *zap.Logger,
) AuthService {
	return &authService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		jwtManager:  jwtManager,
		logger:      logger,
	}
}

func (s *authService) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	// Check if email already exists
	exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, common.NewInternalError("Lỗi kiểm tra email", err)
	}
	if exists {
		return nil, common.NewConflictError("Email đã được sử dụng")
	}

	// Check if username already exists
	exists, err = s.userRepo.ExistsByUsername(ctx, req.Username)
	if err != nil {
		return nil, common.NewInternalError("Lỗi kiểm tra username", err)
	}
	if exists {
		return nil, common.NewConflictError("Username đã được sử dụng")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, common.NewInternalError("Lỗi mã hóa mật khẩu", err)
	}

	// Create user
	user := &model.User{
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: hashedPassword,
		FullName:     req.FullName,
		IsActive:     true,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		s.logger.Error("Failed to create user", zap.Error(err))
		return nil, common.NewInternalError("Không thể tạo tài khoản", err)
	}

	// Generate token pair
	tokenPair, err := s.jwtManager.GenerateTokenPair(user.ID, user.Email)
	if err != nil {
		s.logger.Error("Failed to generate token pair", zap.Error(err))
		return nil, common.NewInternalError("Không thể tạo token", err)
	}

	s.logger.Info("User registered", zap.String("email", user.Email), zap.String("user_id", user.ID.String()))

	return &dto.AuthResponse{
		User:         toUserResponse(user),
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresAt:    tokenPair.ExpiresAt,
	}, nil
}

func (s *authService) Login(ctx context.Context, req *dto.LoginRequest, userAgent, ip string) (*dto.AuthResponse, error) {
	// Find user by email
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		s.logger.Error("Failed to get user by email", zap.Error(err))
		return nil, common.NewInternalError("Lỗi hệ thống", err)
	}
	if user == nil {
		return nil, common.NewUnauthorizedError("Email hoặc mật khẩu không đúng")
	}

	// Check password
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, common.NewUnauthorizedError("Email hoặc mật khẩu không đúng")
	}

	// Check user status
	if !user.IsActive {
		return nil, common.NewForbiddenError("Tài khoản đã bị vô hiệu hóa")
	}

	// Generate token pair
	tokenPair, err := s.jwtManager.GenerateTokenPair(user.ID, user.Email)
	if err != nil {
		s.logger.Error("Failed to generate token pair", zap.Error(err))
		return nil, common.NewInternalError("Không thể tạo token", err)
	}

	// Store session in Redis
	sessionID := uuid.New().String()
	sessionData := &repository.SessionData{
		UserID:       user.ID,
		RefreshToken: tokenPair.RefreshToken,
		UserAgent:    userAgent,
		IP:           ip,
		CreatedAt:    time.Now(),
	}
	if err := s.sessionRepo.Create(ctx, sessionID, sessionData, 7*24*time.Hour); err != nil {
		s.logger.Error("Failed to create session", zap.Error(err))
		// Don't fail login if session storage fails
	}

	// Update last login
	now := time.Now()
	user.LastLoginAt = &now
	if err := s.userRepo.Update(ctx, user); err != nil {
		s.logger.Error("Failed to update last login", zap.Error(err))
	}

	s.logger.Info("User logged in", zap.String("email", user.Email), zap.String("ip", ip))

	return &dto.AuthResponse{
		User:         toUserResponse(user),
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresAt:    tokenPair.ExpiresAt,
	}, nil
}

func (s *authService) Logout(ctx context.Context, userID uuid.UUID, tokenID string) error {
	// Blacklist the current access token
	if tokenID != "" {
		if err := s.sessionRepo.BlacklistToken(ctx, tokenID, 15*time.Minute); err != nil {
			s.logger.Error("Failed to blacklist token", zap.Error(err))
		}
	}

	s.logger.Info("User logged out", zap.String("user_id", userID.String()))
	return nil
}

func (s *authService) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.TokenResponse, error) {
	// Validate refresh token
	claims, err := s.jwtManager.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, common.NewUnauthorizedError("Refresh token không hợp lệ")
	}

	// Check if token is blacklisted
	blacklisted, err := s.sessionRepo.IsTokenBlacklisted(ctx, claims.ID)
	if err != nil {
		s.logger.Error("Failed to check token blacklist", zap.Error(err))
	}
	if blacklisted {
		return nil, common.NewUnauthorizedError("Token đã bị thu hồi")
	}

	// Generate new token pair (token rotation)
	tokenPair, err := s.jwtManager.GenerateTokenPair(claims.UserID, claims.Email)
	if err != nil {
		s.logger.Error("Failed to generate token pair", zap.Error(err))
		return nil, common.NewInternalError("Không thể tạo token mới", err)
	}

	// Blacklist the old refresh token
	if err := s.sessionRepo.BlacklistToken(ctx, claims.ID, 7*24*time.Hour); err != nil {
		s.logger.Error("Failed to blacklist old refresh token", zap.Error(err))
	}

	return &dto.TokenResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresAt:    tokenPair.ExpiresAt,
	}, nil
}

func (s *authService) ChangePassword(ctx context.Context, userID uuid.UUID, req *dto.ChangePasswordRequest) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return common.NewInternalError("Lỗi hệ thống", err)
	}
	if user == nil {
		return common.NewNotFoundError("Không tìm thấy người dùng")
	}

	// Verify old password
	if !utils.CheckPassword(req.OldPassword, user.PasswordHash) {
		return common.NewBadRequestError("Mật khẩu cũ không đúng")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return common.NewInternalError("Lỗi mã hóa mật khẩu", err)
	}

	// Update password
	if err := s.userRepo.UpdatePassword(ctx, userID, hashedPassword); err != nil {
		return common.NewInternalError("Không thể cập nhật mật khẩu", err)
	}

	// Invalidate all sessions
	if err := s.sessionRepo.DeleteAllForUser(ctx, userID); err != nil {
		s.logger.Error("Failed to delete all sessions after password change", zap.Error(err))
	}

	return nil
}

func (s *authService) GetProfile(ctx context.Context, userID uuid.UUID) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, common.NewInternalError("Lỗi hệ thống", err)
	}
	if user == nil {
		return nil, common.NewNotFoundError("Không tìm thấy người dùng")
	}

	resp := toUserResponse(user)
	resp.Preferences = &dto.UserPreferencesResponse{
		Theme:        user.Theme,
		Locale:       user.Locale,
		CalendarView: user.DefaultCalendarView,
		WeekStartDay: 0,
	}
	return &resp, nil
}

func (s *authService) UpdateProfile(ctx context.Context, userID uuid.UUID, req *dto.UpdateProfileRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, common.NewInternalError("Lỗi hệ thống", err)
	}
	if user == nil {
		return nil, common.NewNotFoundError("Không tìm thấy người dùng")
	}

	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.Username != "" && req.Username != user.Username {
		exists, err := s.userRepo.ExistsByUsername(ctx, req.Username)
		if err != nil {
			return nil, common.NewInternalError("Lỗi kiểm tra username", err)
		}
		if exists {
			return nil, common.NewConflictError("Username đã được sử dụng")
		}
		user.Username = req.Username
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, common.NewInternalError("Không thể cập nhật hồ sơ", err)
	}

	resp := toUserResponse(user)
	return &resp, nil
}

func (s *authService) UpdatePreferences(ctx context.Context, userID uuid.UUID, req *dto.UpdatePreferencesRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, common.NewInternalError("Lỗi hệ thống", err)
	}
	if user == nil {
		return nil, common.NewNotFoundError("Không tìm thấy người dùng")
	}

	if req.Theme != "" {
		user.Theme = req.Theme
	}
	if req.Locale != "" {
		user.Locale = req.Locale
	}
	if req.CalendarView != "" {
		user.DefaultCalendarView = req.CalendarView
	}
	if req.WeekStartDay != nil {
		// Store in user model if needed; for now we just acknowledge
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, common.NewInternalError("Không thể cập nhật tùy chọn", err)
	}

	resp := toUserResponse(user)
	resp.Preferences = &dto.UserPreferencesResponse{
		Theme:        user.Theme,
		Locale:       user.Locale,
		CalendarView: user.DefaultCalendarView,
		WeekStartDay: 0,
	}
	return &resp, nil
}

func (s *authService) DeleteAccount(ctx context.Context, userID uuid.UUID, password string) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return common.NewInternalError("Lỗi hệ thống", err)
	}
	if user == nil {
		return common.NewNotFoundError("Không tìm thấy người dùng")
	}

	if !utils.CheckPassword(password, user.PasswordHash) {
		return common.NewBadRequestError("Mật khẩu không đúng")
	}

	if err := s.userRepo.Delete(ctx, userID); err != nil {
		return common.NewInternalError("Không thể xoá tài khoản", err)
	}

	// Invalidate all sessions
	if err := s.sessionRepo.DeleteAllForUser(ctx, userID); err != nil {
		s.logger.Error("Failed to delete sessions after account deletion", zap.Error(err))
	}

	s.logger.Info("User account deleted", zap.String("user_id", userID.String()))
	return nil
}

// toUserResponse converts a User model to UserResponse DTO
func toUserResponse(user *model.User) dto.UserResponse {
	resp := dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FullName:  user.FullName,
		AvatarURL: user.AvatarURL,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	return resp
}
