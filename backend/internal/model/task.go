package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	CategoryID   *uuid.UUID `gorm:"type:uuid;index" json:"category_id"`
	ParentTaskID *uuid.UUID `gorm:"type:uuid;index" json:"parent_task_id"`

	Title       string `gorm:"size:500;not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`

	// Scheduling
	DueDate   *time.Time `gorm:"type:date;index" json:"due_date"`
	DueTime   *string    `gorm:"type:time" json:"due_time"`
	StartDate *time.Time `gorm:"type:date" json:"start_date"`
	StartTime *string    `gorm:"type:time" json:"start_time"`

	// Lunar date reference
	LunarDay     *int `json:"lunar_day"`
	LunarMonth   *int `json:"lunar_month"`
	LunarYear    *int `json:"lunar_year"`
	IsLunarBased bool `gorm:"default:false" json:"is_lunar_based"`

	// Status & Priority
	Status   string `gorm:"size:20;default:todo;index" json:"status"`
	Priority string `gorm:"size:10;default:medium" json:"priority"`

	// Recurrence
	IsRecurring       bool       `gorm:"default:false" json:"is_recurring"`
	RecurrenceRule    *string    `gorm:"size:255" json:"recurrence_rule"`
	RecurrenceEndDate *time.Time `gorm:"type:date" json:"recurrence_end_date"`

	// Progress
	Progress    int        `gorm:"default:0" json:"progress"`
	CompletedAt *time.Time `json:"completed_at"`

	// Meta
	SortOrder      int      `gorm:"default:0" json:"sort_order"`
	EstimatedHours *float64 `gorm:"type:decimal(5,2)" json:"estimated_hours"`
	ActualHours    *float64 `gorm:"type:decimal(5,2)" json:"actual_hours"`

	// Relations
	User       User      `gorm:"foreignKey:UserID" json:"-"`
	Category   *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	ParentTask *Task     `gorm:"foreignKey:ParentTaskID" json:"-"`
	SubTasks   []Task    `gorm:"foreignKey:ParentTaskID" json:"subtasks,omitempty"`
	Tags       []Tag     `gorm:"many2many:task_tags" json:"tags,omitempty"`

	// Timestamps
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Task) TableName() string {
	return "tasks"
}

func (t *Task) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

// TaskTag is the join table for tasks and tags
type TaskTag struct {
	TaskID uuid.UUID `gorm:"type:uuid;primaryKey" json:"task_id"`
	TagID  uuid.UUID `gorm:"type:uuid;primaryKey" json:"tag_id"`
}

func (TaskTag) TableName() string {
	return "task_tags"
}

// Task status constants
const (
	TaskStatusTodo       = "todo"
	TaskStatusInProgress = "in_progress"
	TaskStatusDone       = "done"
	TaskStatusCancelled  = "cancelled"
)

// Task priority constants
const (
	TaskPriorityLow    = "low"
	TaskPriorityMedium = "medium"
	TaskPriorityHigh   = "high"
	TaskPriorityUrgent = "urgent"
)

// IsValidStatus checks if the status is valid
func IsValidTaskStatus(status string) bool {
	switch status {
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusDone, TaskStatusCancelled:
		return true
	}
	return false
}

// IsValidPriority checks if the priority is valid
func IsValidTaskPriority(priority string) bool {
	switch priority {
	case TaskPriorityLow, TaskPriorityMedium, TaskPriorityHigh, TaskPriorityUrgent:
		return true
	}
	return false
}
