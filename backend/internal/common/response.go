package common

import (
	"math"

	"github.com/gofiber/fiber/v2"
)

// Standard API Response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Error Response
type ErrorResponse struct {
	Status string      `json:"status"`
	Error  ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Details []ValidationError `json:"details,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Paginated Response
type PaginatedResponse struct {
	Status string         `json:"status"`
	Data   interface{}    `json:"data"`
	Meta   PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
}

// Success sends a success response
func Success(c *fiber.Ctx, data interface{}) error {
	return c.JSON(Response{
		Status: "success",
		Data:   data,
	})
}

// SuccessWithMessage sends a success response with a message
func SuccessWithMessage(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Created sends a 201 created response
func Created(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Status: "success",
		Data:   data,
	})
}

// NoContent sends a 204 no content response
func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

// Paginated sends a paginated response
func Paginated(c *fiber.Ctx, data interface{}, total int64, page, limit int) error {
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return c.JSON(PaginatedResponse{
		Status: "success",
		Data:   data,
		Meta: PaginationMeta{
			Page:       page,
			Limit:      limit,
			TotalItems: total,
			TotalPages: totalPages,
			HasNext:    page < totalPages,
			HasPrev:    page > 1,
		},
	})
}

// Error sends an error response
func Error(c *fiber.Ctx, statusCode int, code, message string) error {
	return c.Status(statusCode).JSON(ErrorResponse{
		Status: "error",
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	})
}

// ValidationErr sends a validation error response
func ValidationErr(c *fiber.Ctx, errors []ValidationError) error {
	return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
		Status: "error",
		Error: ErrorDetail{
			Code:    "VALIDATION_ERROR",
			Message: "Dữ liệu không hợp lệ",
			Details: errors,
		},
	})
}

// NotFound sends a 404 not found response
func NotFound(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusNotFound, "NOT_FOUND", message)
}

// Unauthorized sends a 401 unauthorized response
func Unauthorized(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusUnauthorized, "UNAUTHORIZED", message)
}

// Forbidden sends a 403 forbidden response
func Forbidden(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusForbidden, "FORBIDDEN", message)
}

// InternalError sends a 500 internal server error
func InternalError(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", message)
}
