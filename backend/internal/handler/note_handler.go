package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"zplus-calendar/internal/common"
	"zplus-calendar/internal/dto"
	"zplus-calendar/internal/repository"
	"zplus-calendar/internal/service"
)

// NoteHandler handles note-related HTTP requests
type NoteHandler struct {
	noteService service.NoteService
	validator   *validator.Validate
	logger      *zap.Logger
}

func NewNoteHandler(noteService service.NoteService, logger *zap.Logger) *NoteHandler {
	return &NoteHandler{
		noteService: noteService,
		validator:   validator.New(),
		logger:      logger,
	}
}

// Create creates a new note
// @Summary Create a note
// @Tags notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.CreateNoteRequest true "Create note"
// @Success 201 {object} common.Response{data=dto.NoteResponse}
// @Router /notes [post]
func (h *NoteHandler) Create(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	var req dto.CreateNoteRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	resp, err := h.noteService.Create(c.Context(), userID.String(), req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Created(c, resp)
}

// List lists notes with filters and pagination
// @Summary List notes
// @Tags notes
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param date_from query string false "Filter from date"
// @Param date_to query string false "Filter to date"
// @Param tag query string false "Filter by tag"
// @Param pinned query bool false "Filter pinned"
// @Success 200 {object} common.PaginatedResponse{data=[]dto.NoteResponse}
// @Router /notes [get]
func (h *NoteHandler) List(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	pagination := common.NewPagination(c)

	params := repository.NoteListParams{
		Limit:    int64(pagination.Limit),
		Offset:   int64(pagination.Offset),
		DateFrom: c.Query("date_from"),
		DateTo:   c.Query("date_to"),
		Tag:      c.Query("tag"),
	}

	if pinned := c.Query("pinned"); pinned == "true" {
		t := true
		params.IsPinned = &t
	} else if pinned == "false" {
		f := false
		params.IsPinned = &f
	}

	notes, total, err := h.noteService.List(c.Context(), userID.String(), params)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Paginated(c, notes, total, pagination.Page, pagination.Limit)
}

// Get retrieves a single note
// @Summary Get a note
// @Tags notes
// @Produce json
// @Security BearerAuth
// @Param id path string true "Note ID"
// @Success 200 {object} common.Response{data=dto.NoteResponse}
// @Router /notes/{id} [get]
func (h *NoteHandler) Get(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	noteID := c.Params("id")
	if noteID == "" {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_ID", "ID không hợp lệ")
	}

	resp, err := h.noteService.GetByID(c.Context(), noteID, userID.String())
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// Update updates a note
// @Summary Update a note
// @Tags notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Note ID"
// @Param body body dto.UpdateNoteRequest true "Update note"
// @Success 200 {object} common.Response{data=dto.NoteResponse}
// @Router /notes/{id} [put]
func (h *NoteHandler) Update(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	noteID := c.Params("id")
	if noteID == "" {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_ID", "ID không hợp lệ")
	}

	var req dto.UpdateNoteRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	resp, err := h.noteService.Update(c.Context(), noteID, userID.String(), req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// Delete deletes a note
// @Summary Delete a note
// @Tags notes
// @Produce json
// @Security BearerAuth
// @Param id path string true "Note ID"
// @Success 204 "No Content"
// @Router /notes/{id} [delete]
func (h *NoteHandler) Delete(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	noteID := c.Params("id")
	if noteID == "" {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_ID", "ID không hợp lệ")
	}

	if err := h.noteService.Delete(c.Context(), noteID, userID.String()); err != nil {
		return handleServiceError(c, err)
	}

	return common.NoContent(c)
}

// GetByDate retrieves notes by date
// @Summary Get notes by date
// @Tags notes
// @Produce json
// @Security BearerAuth
// @Param date path string true "Date (YYYY-MM-DD)"
// @Success 200 {object} common.Response{data=[]dto.NoteResponse}
// @Router /notes/by-date/{date} [get]
func (h *NoteHandler) GetByDate(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	date := c.Params("date")
	if date == "" {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_DATE", "Ngày không hợp lệ")
	}

	notes, err := h.noteService.GetByDate(c.Context(), userID.String(), date)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, notes)
}

// Search searches notes by keyword
// @Summary Search notes
// @Tags notes
// @Produce json
// @Security BearerAuth
// @Param q query string true "Search query"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} common.PaginatedResponse{data=[]dto.NoteResponse}
// @Router /notes/search [get]
func (h *NoteHandler) Search(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	query := c.Query("q")
	if query == "" {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Vui lòng nhập từ khoá tìm kiếm")
	}

	pagination := common.NewPagination(c)

	notes, total, err := h.noteService.Search(c.Context(), userID.String(), query, int64(pagination.Limit), int64(pagination.Offset))
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Paginated(c, notes, total, pagination.Page, pagination.Limit)
}
