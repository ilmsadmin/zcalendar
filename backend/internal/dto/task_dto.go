package dto

import (
	"time"

	"zplus-calendar/internal/model"

	"github.com/google/uuid"
)

// ── Task Request DTOs ──────────────────────────────────────────────

type CreateTaskRequest struct {
	Title        string     `json:"title" validate:"required,min=1,max=500"`
	Description  string     `json:"description" validate:"max=5000"`
	CategoryID   *uuid.UUID `json:"category_id"`
	ParentTaskID *uuid.UUID `json:"parent_task_id"`

	// Scheduling
	DueDate   *string `json:"due_date"`   // "2024-01-15"
	DueTime   *string `json:"due_time"`   // "14:30"
	StartDate *string `json:"start_date"` // "2024-01-14"
	StartTime *string `json:"start_time"` // "09:00"

	// Lunar
	IsLunarBased bool `json:"is_lunar_based"`
	LunarDay     *int `json:"lunar_day"`
	LunarMonth   *int `json:"lunar_month"`
	LunarYear    *int `json:"lunar_year"`

	// Priority & Recurrence
	Priority          string  `json:"priority" validate:"omitempty,oneof=low medium high urgent"`
	IsRecurring       bool    `json:"is_recurring"`
	RecurrenceRule    *string `json:"recurrence_rule"`
	RecurrenceEndDate *string `json:"recurrence_end_date"`

	// Meta
	EstimatedHours *float64    `json:"estimated_hours"`
	TagIDs         []uuid.UUID `json:"tag_ids"`
}

type UpdateTaskRequest struct {
	Title        *string    `json:"title" validate:"omitempty,min=1,max=500"`
	Description  *string    `json:"description" validate:"omitempty,max=5000"`
	CategoryID   *uuid.UUID `json:"category_id"`
	ParentTaskID *uuid.UUID `json:"parent_task_id"`

	DueDate   *string `json:"due_date"`
	DueTime   *string `json:"due_time"`
	StartDate *string `json:"start_date"`
	StartTime *string `json:"start_time"`

	IsLunarBased *bool `json:"is_lunar_based"`
	LunarDay     *int  `json:"lunar_day"`
	LunarMonth   *int  `json:"lunar_month"`
	LunarYear    *int  `json:"lunar_year"`

	Priority          *string      `json:"priority" validate:"omitempty,oneof=low medium high urgent"`
	IsRecurring       *bool        `json:"is_recurring"`
	RecurrenceRule    *string      `json:"recurrence_rule"`
	RecurrenceEndDate *string      `json:"recurrence_end_date"`
	Progress          *int         `json:"progress" validate:"omitempty,min=0,max=100"`
	EstimatedHours    *float64     `json:"estimated_hours"`
	ActualHours       *float64     `json:"actual_hours"`
	SortOrder         *int         `json:"sort_order"`
	TagIDs            *[]uuid.UUID `json:"tag_ids"`
}

type UpdateTaskStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=todo in_progress done cancelled"`
}

// ── Task Query DTOs ────────────────────────────────────────────────

type TaskListQuery struct {
	Status     string `query:"status"`
	Priority   string `query:"priority"`
	CategoryID string `query:"category_id"`
	TagID      string `query:"tag_id"`
	DateFrom   string `query:"date_from"` // "2024-01-01"
	DateTo     string `query:"date_to"`   // "2024-01-31"
	Search     string `query:"search"`
}

// ── Task Response DTOs ─────────────────────────────────────────────

type TaskResponse struct {
	ID           uuid.UUID  `json:"id"`
	UserID       uuid.UUID  `json:"user_id"`
	CategoryID   *uuid.UUID `json:"category_id"`
	ParentTaskID *uuid.UUID `json:"parent_task_id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`

	DueDate   *string `json:"due_date"`
	DueTime   *string `json:"due_time"`
	StartDate *string `json:"start_date"`
	StartTime *string `json:"start_time"`

	LunarDay     *int `json:"lunar_day"`
	LunarMonth   *int `json:"lunar_month"`
	LunarYear    *int `json:"lunar_year"`
	IsLunarBased bool `json:"is_lunar_based"`

	Status   string `json:"status"`
	Priority string `json:"priority"`

	IsRecurring       bool    `json:"is_recurring"`
	RecurrenceRule    *string `json:"recurrence_rule"`
	RecurrenceEndDate *string `json:"recurrence_end_date"`

	Progress    int        `json:"progress"`
	CompletedAt *time.Time `json:"completed_at"`

	SortOrder      int      `json:"sort_order"`
	EstimatedHours *float64 `json:"estimated_hours"`
	ActualHours    *float64 `json:"actual_hours"`

	Category *CategoryResponse `json:"category,omitempty"`
	Tags     []TagResponse     `json:"tags,omitempty"`
	SubTasks []TaskResponse    `json:"subtasks,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToTaskResponse converts a model.Task to TaskResponse
func ToTaskResponse(t *model.Task) TaskResponse {
	resp := TaskResponse{
		ID:             t.ID,
		UserID:         t.UserID,
		CategoryID:     t.CategoryID,
		ParentTaskID:   t.ParentTaskID,
		Title:          t.Title,
		Description:    t.Description,
		DueTime:        t.DueTime,
		StartTime:      t.StartTime,
		LunarDay:       t.LunarDay,
		LunarMonth:     t.LunarMonth,
		LunarYear:      t.LunarYear,
		IsLunarBased:   t.IsLunarBased,
		Status:         t.Status,
		Priority:       t.Priority,
		IsRecurring:    t.IsRecurring,
		RecurrenceRule: t.RecurrenceRule,
		Progress:       t.Progress,
		CompletedAt:    t.CompletedAt,
		SortOrder:      t.SortOrder,
		EstimatedHours: t.EstimatedHours,
		ActualHours:    t.ActualHours,
		CreatedAt:      t.CreatedAt,
		UpdatedAt:      t.UpdatedAt,
	}

	if t.DueDate != nil {
		d := t.DueDate.Format("2006-01-02")
		resp.DueDate = &d
	}
	if t.StartDate != nil {
		d := t.StartDate.Format("2006-01-02")
		resp.StartDate = &d
	}
	if t.RecurrenceEndDate != nil {
		d := t.RecurrenceEndDate.Format("2006-01-02")
		resp.RecurrenceEndDate = &d
	}

	if t.Category != nil {
		cat := ToCategoryResponse(t.Category)
		resp.Category = &cat
	}

	for _, tag := range t.Tags {
		resp.Tags = append(resp.Tags, ToTagResponse(&tag))
	}

	for _, sub := range t.SubTasks {
		resp.SubTasks = append(resp.SubTasks, ToTaskResponse(&sub))
	}

	return resp
}

// ToTaskListResponse converts a list of tasks
func ToTaskListResponse(tasks []model.Task) []TaskResponse {
	result := make([]TaskResponse, len(tasks))
	for i, t := range tasks {
		result[i] = ToTaskResponse(&t)
	}
	return result
}
