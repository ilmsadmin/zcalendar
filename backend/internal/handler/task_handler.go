package handler

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"zplus-calendar/internal/common"
	"zplus-calendar/internal/dto"
	"zplus-calendar/internal/repository"
	"zplus-calendar/internal/service"
)

// TaskHandler handles task-related HTTP requests
type TaskHandler struct {
	taskService service.TaskService
	validator   *validator.Validate
	logger      *zap.Logger
}

func NewTaskHandler(taskService service.TaskService, logger *zap.Logger) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
		validator:   validator.New(),
		logger:      logger,
	}
}

// Create creates a new task
// @Summary Create a task
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.CreateTaskRequest true "Create task request"
// @Success 201 {object} common.Response{data=dto.TaskResponse}
// @Router /tasks [post]
func (h *TaskHandler) Create(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	var req dto.CreateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	resp, err := h.taskService.Create(userID, req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Created(c, resp)
}

// List lists tasks with filters and pagination
// @Summary List tasks
// @Tags tasks
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Param status query string false "Filter by status"
// @Param priority query string false "Filter by priority"
// @Param category_id query string false "Filter by category"
// @Param date_from query string false "Filter from date"
// @Param date_to query string false "Filter to date"
// @Param sort_by query string false "Sort field"
// @Param sort_order query string false "Sort order (asc/desc)"
// @Success 200 {object} common.PaginatedResponse{data=[]dto.TaskResponse}
// @Router /tasks [get]
func (h *TaskHandler) List(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	pagination := common.NewPagination(c)
	sort := common.NewSortParams(c, "created_at", "desc")

	params := repository.TaskListParams{
		Page:      pagination.Page,
		Limit:     pagination.Limit,
		Offset:    pagination.Offset,
		Status:    c.Query("status"),
		Priority:  c.Query("priority"),
		SortBy:    sort.SortBy,
		SortOrder: sort.SortOrder,
	}

	// Optional category ID filter
	if catID := c.Query("category_id"); catID != "" {
		id, err := uuid.Parse(catID)
		if err == nil {
			params.CategoryID = &id
		}
	}

	// Optional tag ID filter
	if tagID := c.Query("tag_id"); tagID != "" {
		id, err := uuid.Parse(tagID)
		if err == nil {
			params.TagID = &id
		}
	}

	// Date range filters
	if dateFrom := c.Query("date_from"); dateFrom != "" {
		t, err := time.Parse("2006-01-02", dateFrom)
		if err == nil {
			params.DateFrom = &t
		}
	}
	if dateTo := c.Query("date_to"); dateTo != "" {
		t, err := time.Parse("2006-01-02", dateTo)
		if err == nil {
			params.DateTo = &t
		}
	}

	tasks, total, err := h.taskService.List(userID, params)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Paginated(c, tasks, total, pagination.Page, pagination.Limit)
}

// Get retrieves a single task by ID
// @Summary Get a task
// @Tags tasks
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Success 200 {object} common.Response{data=dto.TaskResponse}
// @Router /tasks/{id} [get]
func (h *TaskHandler) Get(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	taskID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_ID", "ID không hợp lệ")
	}

	resp, err := h.taskService.GetByID(taskID, userID)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// Update updates an existing task
// @Summary Update a task
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Param body body dto.UpdateTaskRequest true "Update request"
// @Success 200 {object} common.Response{data=dto.TaskResponse}
// @Router /tasks/{id} [put]
func (h *TaskHandler) Update(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	taskID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_ID", "ID không hợp lệ")
	}

	var req dto.UpdateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	resp, err := h.taskService.Update(taskID, userID, req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, resp)
}

// Delete deletes a task (soft delete)
// @Summary Delete a task
// @Tags tasks
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Success 204 "No Content"
// @Router /tasks/{id} [delete]
func (h *TaskHandler) Delete(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	taskID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_ID", "ID không hợp lệ")
	}

	if err := h.taskService.Delete(taskID, userID); err != nil {
		return handleServiceError(c, err)
	}

	return common.NoContent(c)
}

// UpdateStatus updates task status
// @Summary Update task status
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Param body body dto.UpdateTaskStatusRequest true "Status update"
// @Success 200 {object} common.Response
// @Router /tasks/{id}/status [patch]
func (h *TaskHandler) UpdateStatus(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	taskID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_ID", "ID không hợp lệ")
	}

	var req dto.UpdateTaskStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Dữ liệu không hợp lệ")
	}

	if err := h.validator.Struct(&req); err != nil {
		return common.ValidationErr(c, formatValidationErrors(err))
	}

	if err := h.taskService.UpdateStatus(taskID, userID, req.Status); err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, fiber.Map{"message": "Cập nhật trạng thái thành công"})
}

// GetByDate retrieves tasks by date
// @Summary Get tasks by date
// @Tags tasks
// @Produce json
// @Security BearerAuth
// @Param date path string true "Date (YYYY-MM-DD)"
// @Success 200 {object} common.Response{data=[]dto.TaskResponse}
// @Router /tasks/by-date/{date} [get]
func (h *TaskHandler) GetByDate(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	date, err := time.Parse("2006-01-02", c.Params("date"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_DATE", "Ngày không hợp lệ (YYYY-MM-DD)")
	}

	tasks, err := h.taskService.GetByDate(userID, date)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Success(c, tasks)
}

// Search searches tasks by keyword
// @Summary Search tasks
// @Tags tasks
// @Produce json
// @Security BearerAuth
// @Param q query string true "Search query"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} common.PaginatedResponse{data=[]dto.TaskResponse}
// @Router /tasks/search [get]
func (h *TaskHandler) Search(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return common.Unauthorized(c, "Không xác thực")
	}

	query := c.Query("q")
	if query == "" {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Vui lòng nhập từ khoá tìm kiếm")
	}

	pagination := common.NewPagination(c)

	tasks, total, err := h.taskService.Search(userID, query, pagination.Limit, pagination.Offset)
	if err != nil {
		return handleServiceError(c, err)
	}

	return common.Paginated(c, tasks, total, pagination.Page, pagination.Limit)
}

// getUserID extracts the user ID from Fiber context (set by auth middleware)
func getUserID(c *fiber.Ctx) (uuid.UUID, error) {
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	return uuid.Parse(userIDStr)
}
