package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"zplus-calendar/internal/common"
	"zplus-calendar/internal/dto"
	"zplus-calendar/internal/service"
)

// ── Category Handler ───────────────────────────────────────────────

type CategoryHandler struct {
	service   service.CategoryService
	validator *validator.Validate
	logger    *zap.Logger
}

func NewCategoryHandler(svc service.CategoryService, logger *zap.Logger) *CategoryHandler {
	return &CategoryHandler{
		service:   svc,
		validator: validator.New(),
		logger:    logger,
	}
}

// Create creates a new category
// @Summary Create a category
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.CreateCategoryRequest true "Create category"
// @Success 201 {object} common.Response{data=dto.CategoryResponse}
// @Router /categories [post]
func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	var req dto.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	resp, err := h.service.Create(userID, req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Created(c, resp)
}

// List lists all categories for the current user
// @Summary List categories
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param type query string false "Filter by type (task, note, both)"
// @Success 200 {object} common.Response{data=[]dto.CategoryResponse}
// @Router /categories [get]
func (h *CategoryHandler) List(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	catType := c.Query("type")
	resp, err := h.service.List(userID, catType)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// Update updates a category
// @Summary Update a category
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Param body body dto.UpdateCategoryRequest true "Update category"
// @Success 200 {object} common.Response{data=dto.CategoryResponse}
// @Router /categories/{id} [put]
func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	catID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_ID", "ID không hợp lệ")
	}

	var req dto.UpdateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	resp, err := h.service.Update(catID, userID, req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// Delete deletes a category
// @Summary Delete a category
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Success 204 "No Content"
// @Router /categories/{id} [delete]
func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	catID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_ID", "ID không hợp lệ")
	}

	if err := h.service.Delete(catID, userID); err != nil {
		return handleServiceError(c, err)
	}

	return common.NoContent(c)
}

// ── Tag Handler ────────────────────────────────────────────────────

type TagHandler struct {
	service   service.TagService
	validator *validator.Validate
	logger    *zap.Logger
}

func NewTagHandler(svc service.TagService, logger *zap.Logger) *TagHandler {
	return &TagHandler{
		service:   svc,
		validator: validator.New(),
		logger:    logger,
	}
}

// Create creates a new tag
// @Summary Create a tag
// @Tags tags
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.CreateTagRequest true "Create tag"
// @Success 201 {object} common.Response{data=dto.TagResponse}
// @Router /tags [post]
func (h *TagHandler) Create(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	var req dto.CreateTagRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	resp, err := h.service.Create(userID, req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Created(c, resp)
}

// List lists all tags for the current user
// @Summary List tags
// @Tags tags
// @Produce json
// @Security BearerAuth
// @Success 200 {object} common.Response{data=[]dto.TagResponse}
// @Router /tags [get]
func (h *TagHandler) List(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	resp, err := h.service.List(userID)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// Update updates a tag
// @Summary Update a tag
// @Tags tags
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Tag ID"
// @Param body body dto.UpdateTagRequest true "Update tag"
// @Success 200 {object} common.Response{data=dto.TagResponse}
// @Router /tags/{id} [put]
func (h *TagHandler) Update(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	tagID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_ID", "ID không hợp lệ")
	}

	var req dto.UpdateTagRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	resp, err := h.service.Update(tagID, userID, req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// Delete deletes a tag
// @Summary Delete a tag
// @Tags tags
// @Produce json
// @Security BearerAuth
// @Param id path string true "Tag ID"
// @Success 204 "No Content"
// @Router /tags/{id} [delete]
func (h *TagHandler) Delete(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	tagID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_ID", "ID không hợp lệ")
	}

	if err := h.service.Delete(tagID, userID); err != nil {
		return handleServiceError(c, err)
	}

	return common.NoContent(c)
}
