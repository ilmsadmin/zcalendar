package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"zplus-calendar/internal/model"
)

// TaskRepository defines the interface for task data access
type TaskRepository interface {
	Create(task *model.Task) error
	GetByID(id, userID uuid.UUID) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id, userID uuid.UUID) error
	List(userID uuid.UUID, params TaskListParams) ([]model.Task, int64, error)
	GetByDate(userID uuid.UUID, date time.Time) ([]model.Task, error)
	GetByDateRange(userID uuid.UUID, from, to time.Time) ([]model.Task, error)
	UpdateStatus(id, userID uuid.UUID, status string) error
	Search(userID uuid.UUID, query string, limit, offset int) ([]model.Task, int64, error)
	ReplaceTags(taskID uuid.UUID, tagIDs []uuid.UUID) error
}

// TaskListParams holds filtering & pagination params for listing tasks
type TaskListParams struct {
	Page       int
	Limit      int
	Offset     int
	Status     string
	Priority   string
	CategoryID *uuid.UUID
	TagID      *uuid.UUID
	DateFrom   *time.Time
	DateTo     *time.Time
	SortBy     string
	SortOrder  string
}

// ── GORM Implementation ────────────────────────────────────────────

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetByID(id, userID uuid.UUID) (*model.Task, error) {
	var task model.Task
	err := r.db.
		Preload("Category").
		Preload("Tags").
		Preload("SubTasks").
		Where("id = ? AND user_id = ?", id, userID).
		First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id, userID uuid.UUID) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Task{}).Error
}

func (r *taskRepository) List(userID uuid.UUID, params TaskListParams) ([]model.Task, int64, error) {
	var tasks []model.Task
	var total int64

	query := r.db.Model(&model.Task{}).Where("user_id = ?", userID)

	// Filters
	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}
	if params.Priority != "" {
		query = query.Where("priority = ?", params.Priority)
	}
	if params.CategoryID != nil {
		query = query.Where("category_id = ?", *params.CategoryID)
	}
	if params.TagID != nil {
		query = query.Where("id IN (SELECT task_id FROM task_tags WHERE tag_id = ?)", *params.TagID)
	}
	if params.DateFrom != nil {
		query = query.Where("due_date >= ?", *params.DateFrom)
	}
	if params.DateTo != nil {
		query = query.Where("due_date <= ?", *params.DateTo)
	}

	// Only top-level tasks (no subtasks)
	query = query.Where("parent_task_id IS NULL")

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Sort
	sortBy := params.SortBy
	if sortBy == "" {
		sortBy = "created_at"
	}
	sortOrder := params.SortOrder
	if sortOrder == "" {
		sortOrder = "desc"
	}
	query = query.Order(sortBy + " " + sortOrder)

	// Paginate and preload
	err := query.
		Preload("Category").
		Preload("Tags").
		Preload("SubTasks").
		Offset(params.Offset).
		Limit(params.Limit).
		Find(&tasks).Error

	return tasks, total, err
}

func (r *taskRepository) GetByDate(userID uuid.UUID, date time.Time) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.
		Preload("Category").
		Preload("Tags").
		Where("user_id = ? AND due_date = ?", userID, date.Format("2006-01-02")).
		Order("sort_order ASC, due_time ASC").
		Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetByDateRange(userID uuid.UUID, from, to time.Time) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.
		Preload("Category").
		Preload("Tags").
		Where("user_id = ? AND due_date >= ? AND due_date <= ?", userID, from.Format("2006-01-02"), to.Format("2006-01-02")).
		Order("due_date ASC, sort_order ASC, due_time ASC").
		Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateStatus(id, userID uuid.UUID, status string) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if status == model.TaskStatusDone {
		now := time.Now()
		updates["completed_at"] = &now
		updates["progress"] = 100
	} else {
		updates["completed_at"] = nil
	}

	return r.db.Model(&model.Task{}).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(updates).Error
}

func (r *taskRepository) Search(userID uuid.UUID, query string, limit, offset int) ([]model.Task, int64, error) {
	var tasks []model.Task
	var total int64

	search := "%" + query + "%"
	q := r.db.Model(&model.Task{}).
		Where("user_id = ? AND (title ILIKE ? OR description ILIKE ?)", userID, search, search)

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := q.
		Preload("Category").
		Preload("Tags").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&tasks).Error

	return tasks, total, err
}

func (r *taskRepository) ReplaceTags(taskID uuid.UUID, tagIDs []uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Remove existing tags
		if err := tx.Where("task_id = ?", taskID).Delete(&model.TaskTag{}).Error; err != nil {
			return err
		}

		// Add new tags
		for _, tagID := range tagIDs {
			taskTag := model.TaskTag{TaskID: taskID, TagID: tagID}
			if err := tx.Create(&taskTag).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
