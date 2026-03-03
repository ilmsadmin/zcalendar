package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"zplus-calendar/internal/common"
	"zplus-calendar/internal/dto"
	"zplus-calendar/internal/service"
)

// AuthHandler handles authentication HTTP requests
type AuthHandler struct {
	authService service.AuthService
	validator   *validator.Validate
	logger      *zap.Logger
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authService service.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator.New(),
		logger:      logger,
	}
}

// Register handles user registration
// @Summary Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.RegisterRequest true "Register request"
// @Success 201 {object} common.APIResponse{data=dto.AuthResponse}
// @Failure 400 {object} common.APIResponse
// @Failure 409 {object} common.APIResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	resp, err := h.authService.Register(c.Context(), &req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Created(c, resp)
}

// Login handles user login
// @Summary Login
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.LoginRequest true "Login request"
// @Success 200 {object} common.APIResponse{data=dto.AuthResponse}
// @Failure 401 {object} common.APIResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	userAgent := c.Get("User-Agent")
	ip := c.IP()

	resp, err := h.authService.Login(c.Context(), &req, userAgent, ip)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// Logout handles user logout
// @Summary Logout
// @Tags auth
// @Security BearerAuth
// @Success 204
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}
	tokenID, _ := c.Locals("token_id").(string)

	if err := h.authService.Logout(c.Context(), userID, tokenID); err != nil {
		return handleServiceError(c, err)
	}

	return common.NoContent(c)
}

// RefreshToken handles token refresh
// @Summary Refresh access token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.RefreshTokenRequest true "Refresh token request"
// @Success 200 {object} common.APIResponse{data=dto.TokenResponse}
// @Failure 401 {object} common.APIResponse
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req dto.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	resp, err := h.authService.RefreshToken(c.Context(), &req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// ChangePassword handles password change
// @Summary Change password
// @Tags auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body dto.ChangePasswordRequest true "Change password request"
// @Success 200 {object} common.APIResponse
// @Failure 400 {object} common.APIResponse
// @Router /auth/change-password [post]
func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	var req dto.ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	if err := h.authService.ChangePassword(c.Context(), userID, &req); err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, fiber.Map{"message": "Đổi mật khẩu thành công"})
}

// GetProfile returns the current user's profile
// @Summary Get user profile
// @Tags auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} common.APIResponse{data=dto.UserResponse}
// @Router /auth/profile [get]
func (h *AuthHandler) GetProfile(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	resp, err := h.authService.GetProfile(c.Context(), userID)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// UpdateProfile updates the current user's profile
// @Summary Update user profile
// @Tags auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body dto.UpdateProfileRequest true "Update profile request"
// @Success 200 {object} common.APIResponse{data=dto.UserResponse}
// @Router /auth/profile [put]
func (h *AuthHandler) UpdateProfile(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	var req dto.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	resp, err := h.authService.UpdateProfile(c.Context(), userID, &req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// UpdatePreferences updates the current user's preferences
// @Summary Update user preferences
// @Tags auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body dto.UpdatePreferencesRequest true "Update preferences request"
// @Success 200 {object} common.APIResponse{data=dto.UserResponse}
// @Router /auth/preferences [put]
func (h *AuthHandler) UpdatePreferences(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	var req dto.UpdatePreferencesRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	resp, err := h.authService.UpdatePreferences(c.Context(), userID, &req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// DeleteAccount deletes the current user's account
// @Summary Delete user account
// @Tags auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Router /auth/account [delete]
func (h *AuthHandler) DeleteAccount(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	var req struct {
		Password string `json:"password" validate:"required"`
	}
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}
	if req.Password == "" {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Mật khẩu là bắt buộc")
	}

	if err := h.authService.DeleteAccount(c.Context(), userID, req.Password); err != nil {
		return handleServiceError(c, err)
	}

	return common.NoContent(c)
}

// ─── Helpers ────────────────────────────────────────────────────────

// handleServiceError converts service-layer errors to HTTP responses
func handleServiceError(c *fiber.Ctx, err error) error {
	if appErr, ok := common.IsAppError(err); ok {
		return common.Error(c, appErr.StatusCode, appErr.Code, appErr.Message)
	}
	return common.InternalError(c, "Lỗi hệ thống")
}

// formatValidationErrors converts validator errors to common.ValidationError slice
func formatValidationErrors(err error) []common.ValidationError {
	var errs []common.ValidationError
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := e.Field()
			var msg string
			switch e.Tag() {
			case "required":
				msg = field + " là bắt buộc"
			case "email":
				msg = "Email không hợp lệ"
			case "min":
				msg = field + " phải có ít nhất " + e.Param() + " ký tự"
			case "max":
				msg = field + " không được vượt quá " + e.Param() + " ký tự"
			case "alphanum":
				msg = field + " chỉ được chứa chữ và số"
			default:
				msg = field + " không hợp lệ"
			}
			errs = append(errs, common.ValidationError{Field: field, Message: msg})
		}
	}
	return errs
}
