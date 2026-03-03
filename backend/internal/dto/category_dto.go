package dto

import (
	"time"

	"zplus-calendar/internal/model"

	"github.com/google/uuid"
)

// ── Category Request DTOs ──────────────────────────────────────────

type CreateCategoryRequest struct {
	Name      string `json:"name" validate:"required,min=1,max=100"`
	Color     string `json:"color" validate:"omitempty,hexcolor"`
	Icon      string `json:"icon" validate:"omitempty,max=50"`
	Type      string `json:"type" validate:"required,oneof=task note both"`
	SortOrder int    `json:"sort_order"`
}

type UpdateCategoryRequest struct {
	Name      *string `json:"name" validate:"omitempty,min=1,max=100"`
	Color     *string `json:"color" validate:"omitempty,hexcolor"`
	Icon      *string `json:"icon" validate:"omitempty,max=50"`
	Type      *string `json:"type" validate:"omitempty,oneof=task note both"`
	SortOrder *int    `json:"sort_order"`
}

// ── Category Response DTOs ─────────────────────────────────────────

type CategoryResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Icon      string    `json:"icon"`
	Type      string    `json:"type"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToCategoryResponse(c *model.Category) CategoryResponse {
	return CategoryResponse{
		ID:        c.ID,
		Name:      c.Name,
		Color:     c.Color,
		Icon:      c.Icon,
		Type:      c.Type,
		SortOrder: c.SortOrder,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ToCategoryListResponse(categories []model.Category) []CategoryResponse {
	result := make([]CategoryResponse, len(categories))
	for i, c := range categories {
		result[i] = ToCategoryResponse(&c)
	}
	return result
}

// ── Tag Request DTOs ───────────────────────────────────────────────

type CreateTagRequest struct {
	Name  string `json:"name" validate:"required,min=1,max=50"`
	Color string `json:"color" validate:"omitempty,hexcolor"`
}

type UpdateTagRequest struct {
	Name  *string `json:"name" validate:"omitempty,min=1,max=50"`
	Color *string `json:"color" validate:"omitempty,hexcolor"`
}

// ── Tag Response DTOs ──────────────────────────────────────────────

type TagResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
}

func ToTagResponse(t *model.Tag) TagResponse {
	return TagResponse{
		ID:        t.ID,
		Name:      t.Name,
		Color:     t.Color,
		CreatedAt: t.CreatedAt,
	}
}

func ToTagListResponse(tags []model.Tag) []TagResponse {
	result := make([]TagResponse, len(tags))
	for i, t := range tags {
		result[i] = ToTagResponse(&t)
	}
	return result
}
