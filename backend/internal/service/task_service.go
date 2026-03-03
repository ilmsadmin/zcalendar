package service

import (
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"zplus-calendar/internal/common"
	"zplus-calendar/internal/dto"
	"zplus-calendar/internal/model"
	"zplus-calendar/internal/repository"
)

// TaskService defines the interface for task business logic
type TaskService interface {
	Create(userID uuid.UUID, req dto.CreateTaskRequest) (*dto.TaskResponse, error)
	GetByID(id, userID uuid.UUID) (*dto.TaskResponse, error)
	List(userID uuid.UUID, params repository.TaskListParams) ([]dto.TaskResponse, int64, error)
	Update(id, userID uuid.UUID, req dto.UpdateTaskRequest) (*dto.TaskResponse, error)
	Delete(id, userID uuid.UUID) error
	UpdateStatus(id, userID uuid.UUID, status string) error
	GetByDate(userID uuid.UUID, date time.Time) ([]dto.TaskResponse, error)
	GetByDateRange(userID uuid.UUID, from, to time.Time) ([]dto.TaskResponse, error)
	Search(userID uuid.UUID, query string, limit, offset int) ([]dto.TaskResponse, int64, error)
}

type taskService struct {
	taskRepo repository.TaskRepository
	logger   *zap.Logger
}

func NewTaskService(taskRepo repository.TaskRepository, logger *zap.Logger) TaskService {
	return &taskService{
		taskRepo: taskRepo,
		logger:   logger,
	}
}

func (s *taskService) Create(userID uuid.UUID, req dto.CreateTaskRequest) (*dto.TaskResponse, error) {
	task := model.Task{
		UserID:         userID,
		CategoryID:     req.CategoryID,
		ParentTaskID:   req.ParentTaskID,
		Title:          req.Title,
		Description:    req.Description,
		IsLunarBased:   req.IsLunarBased,
		LunarDay:       req.LunarDay,
		LunarMonth:     req.LunarMonth,
		LunarYear:      req.LunarYear,
		Priority:       req.Priority,
		Status:         model.TaskStatusTodo,
		IsRecurring:    req.IsRecurring,
		RecurrenceRule: req.RecurrenceRule,
		EstimatedHours: req.EstimatedHours,
		DueTime:        req.DueTime,
		StartTime:      req.StartTime,
	}

	if task.Priority == "" {
		task.Priority = model.TaskPriorityMedium
	}

	// Parse dates
	if req.DueDate != nil {
		t, err := time.Parse("2006-01-02", *req.DueDate)
		if err != nil {
			return nil, common.NewBadRequestError("Định dạng due_date không hợp lệ (YYYY-MM-DD)")
		}
		task.DueDate = &t
	}
	if req.StartDate != nil {
		t, err := time.Parse("2006-01-02", *req.StartDate)
		if err != nil {
			return nil, common.NewBadRequestError("Định dạng start_date không hợp lệ (YYYY-MM-DD)")
		}
		task.StartDate = &t
	}
	if req.RecurrenceEndDate != nil {
		t, err := time.Parse("2006-01-02", *req.RecurrenceEndDate)
		if err != nil {
			return nil, common.NewBadRequestError("Định dạng recurrence_end_date không hợp lệ (YYYY-MM-DD)")
		}
		task.RecurrenceEndDate = &t
	}

	if err := s.taskRepo.Create(&task); err != nil {
		s.logger.Error("Failed to create task", zap.Error(err))
		return nil, common.NewInternalError("Không thể tạo task", err)
	}

	// Assign tags
	if len(req.TagIDs) > 0 {
		if err := s.taskRepo.ReplaceTags(task.ID, req.TagIDs); err != nil {
			s.logger.Error("Failed to assign tags", zap.Error(err))
			// Non-fatal: task was created
		}
	}

	// Re-fetch with preloads
	return s.GetByID(task.ID, userID)
}

func (s *taskService) GetByID(id, userID uuid.UUID) (*dto.TaskResponse, error) {
	task, err := s.taskRepo.GetByID(id, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NewNotFoundError("Task không tồn tại")
		}
		s.logger.Error("Failed to get task", zap.Error(err))
		return nil, common.NewInternalError("Không thể lấy thông tin task", err)
	}

	resp := dto.ToTaskResponse(task)
	return &resp, nil
}

func (s *taskService) List(userID uuid.UUID, params repository.TaskListParams) ([]dto.TaskResponse, int64, error) {
	tasks, total, err := s.taskRepo.List(userID, params)
	if err != nil {
		s.logger.Error("Failed to list tasks", zap.Error(err))
		return nil, 0, common.NewInternalError("Không thể lấy danh sách task", err)
	}

	return dto.ToTaskListResponse(tasks), total, nil
}

func (s *taskService) Update(id, userID uuid.UUID, req dto.UpdateTaskRequest) (*dto.TaskResponse, error) {
	task, err := s.taskRepo.GetByID(id, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NewNotFoundError("Task không tồn tại")
		}
		return nil, common.NewInternalError("Không thể lấy thông tin task", err)
	}

	// Apply partial updates
	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.CategoryID != nil {
		task.CategoryID = req.CategoryID
	}
	if req.ParentTaskID != nil {
		task.ParentTaskID = req.ParentTaskID
	}
	if req.DueTime != nil {
		task.DueTime = req.DueTime
	}
	if req.StartTime != nil {
		task.StartTime = req.StartTime
	}
	if req.IsLunarBased != nil {
		task.IsLunarBased = *req.IsLunarBased
	}
	if req.LunarDay != nil {
		task.LunarDay = req.LunarDay
	}
	if req.LunarMonth != nil {
		task.LunarMonth = req.LunarMonth
	}
	if req.LunarYear != nil {
		task.LunarYear = req.LunarYear
	}
	if req.Priority != nil {
		task.Priority = *req.Priority
	}
	if req.IsRecurring != nil {
		task.IsRecurring = *req.IsRecurring
	}
	if req.RecurrenceRule != nil {
		task.RecurrenceRule = req.RecurrenceRule
	}
	if req.Progress != nil {
		task.Progress = *req.Progress
	}
	if req.EstimatedHours != nil {
		task.EstimatedHours = req.EstimatedHours
	}
	if req.ActualHours != nil {
		task.ActualHours = req.ActualHours
	}
	if req.SortOrder != nil {
		task.SortOrder = *req.SortOrder
	}

	// Parse dates
	if req.DueDate != nil {
		t, err := time.Parse("2006-01-02", *req.DueDate)
		if err != nil {
			return nil, common.NewBadRequestError("Định dạng due_date không hợp lệ")
		}
		task.DueDate = &t
	}
	if req.StartDate != nil {
		t, err := time.Parse("2006-01-02", *req.StartDate)
		if err != nil {
			return nil, common.NewBadRequestError("Định dạng start_date không hợp lệ")
		}
		task.StartDate = &t
	}
	if req.RecurrenceEndDate != nil {
		t, err := time.Parse("2006-01-02", *req.RecurrenceEndDate)
		if err != nil {
			return nil, common.NewBadRequestError("Định dạng recurrence_end_date không hợp lệ")
		}
		task.RecurrenceEndDate = &t
	}

	if err := s.taskRepo.Update(task); err != nil {
		s.logger.Error("Failed to update task", zap.Error(err))
		return nil, common.NewInternalError("Không thể cập nhật task", err)
	}

	// Update tags if provided
	if req.TagIDs != nil {
		if err := s.taskRepo.ReplaceTags(task.ID, *req.TagIDs); err != nil {
			s.logger.Error("Failed to update tags", zap.Error(err))
		}
	}

	return s.GetByID(id, userID)
}

func (s *taskService) Delete(id, userID uuid.UUID) error {
	// Verify existence
	if _, err := s.taskRepo.GetByID(id, userID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NewNotFoundError("Task không tồn tại")
		}
		return common.NewInternalError("Không thể xoá task", err)
	}

	if err := s.taskRepo.Delete(id, userID); err != nil {
		s.logger.Error("Failed to delete task", zap.Error(err))
		return common.NewInternalError("Không thể xoá task", err)
	}
	return nil
}

func (s *taskService) UpdateStatus(id, userID uuid.UUID, status string) error {
	if !model.IsValidTaskStatus(status) {
		return common.NewBadRequestError("Trạng thái không hợp lệ")
	}

	// Verify existence
	if _, err := s.taskRepo.GetByID(id, userID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NewNotFoundError("Task không tồn tại")
		}
		return common.NewInternalError("Không thể cập nhật trạng thái", err)
	}

	if err := s.taskRepo.UpdateStatus(id, userID, status); err != nil {
		s.logger.Error("Failed to update task status", zap.Error(err))
		return common.NewInternalError("Không thể cập nhật trạng thái", err)
	}
	return nil
}

func (s *taskService) GetByDate(userID uuid.UUID, date time.Time) ([]dto.TaskResponse, error) {
	tasks, err := s.taskRepo.GetByDate(userID, date)
	if err != nil {
		s.logger.Error("Failed to get tasks by date", zap.Error(err))
		return nil, common.NewInternalError("Không thể lấy task theo ngày", err)
	}
	return dto.ToTaskListResponse(tasks), nil
}

func (s *taskService) GetByDateRange(userID uuid.UUID, from, to time.Time) ([]dto.TaskResponse, error) {
	tasks, err := s.taskRepo.GetByDateRange(userID, from, to)
	if err != nil {
		s.logger.Error("Failed to get tasks by date range", zap.Error(err))
		return nil, common.NewInternalError("Không thể lấy task theo khoảng thời gian", err)
	}
	return dto.ToTaskListResponse(tasks), nil
}

func (s *taskService) Search(userID uuid.UUID, query string, limit, offset int) ([]dto.TaskResponse, int64, error) {
	tasks, total, err := s.taskRepo.Search(userID, query, limit, offset)
	if err != nil {
		s.logger.Error("Failed to search tasks", zap.Error(err))
		return nil, 0, common.NewInternalError("Không thể tìm kiếm task", err)
	}
	return dto.ToTaskListResponse(tasks), total, nil
}
