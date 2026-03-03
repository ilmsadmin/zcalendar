package dto

import (
	"time"
)

// ── Note Request DTOs ──────────────────────────────────────────────

type CreateNoteRequest struct {
	Title      string   `json:"title" validate:"required,min=1,max=500"`
	Content    string   `json:"content" validate:"max=50000"`
	Date       string   `json:"date" validate:"required"` // "2024-01-15"
	CategoryID *string  `json:"category_id"`
	Tags       []string `json:"tags"`
	IsPinned   bool     `json:"is_pinned"`
	Color      string   `json:"color"`

	// Lunar reference
	LunarDay   *int `json:"lunar_day"`
	LunarMonth *int `json:"lunar_month"`
	LunarYear  *int `json:"lunar_year"`
}

type UpdateNoteRequest struct {
	Title      *string   `json:"title" validate:"omitempty,min=1,max=500"`
	Content    *string   `json:"content" validate:"omitempty,max=50000"`
	Date       *string   `json:"date"`
	CategoryID *string   `json:"category_id"`
	Tags       *[]string `json:"tags"`
	IsPinned   *bool     `json:"is_pinned"`
	Color      *string   `json:"color"`

	LunarDay   *int `json:"lunar_day"`
	LunarMonth *int `json:"lunar_month"`
	LunarYear  *int `json:"lunar_year"`
}

// ── Note Document (MongoDB) ────────────────────────────────────────

type NoteDocument struct {
	ID         string   `bson:"_id,omitempty" json:"id"`
	UserID     string   `bson:"user_id" json:"user_id"`
	Title      string   `bson:"title" json:"title"`
	Content    string   `bson:"content" json:"content"`
	Date       string   `bson:"date" json:"date"`
	CategoryID *string  `bson:"category_id,omitempty" json:"category_id"`
	Tags       []string `bson:"tags,omitempty" json:"tags"`
	IsPinned   bool     `bson:"is_pinned" json:"is_pinned"`
	Color      string   `bson:"color" json:"color"`

	LunarDay   *int `bson:"lunar_day,omitempty" json:"lunar_day"`
	LunarMonth *int `bson:"lunar_month,omitempty" json:"lunar_month"`
	LunarYear  *int `bson:"lunar_year,omitempty" json:"lunar_year"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// ── Note Response DTOs ─────────────────────────────────────────────

type NoteResponse struct {
	ID         string   `json:"id"`
	UserID     string   `json:"user_id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Date       string   `json:"date"`
	CategoryID *string  `json:"category_id"`
	Tags       []string `json:"tags"`
	IsPinned   bool     `json:"is_pinned"`
	Color      string   `json:"color"`

	LunarDay   *int `json:"lunar_day"`
	LunarMonth *int `json:"lunar_month"`
	LunarYear  *int `json:"lunar_year"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToNoteResponse(doc *NoteDocument) NoteResponse {
	return NoteResponse{
		ID:         doc.ID,
		UserID:     doc.UserID,
		Title:      doc.Title,
		Content:    doc.Content,
		Date:       doc.Date,
		CategoryID: doc.CategoryID,
		Tags:       doc.Tags,
		IsPinned:   doc.IsPinned,
		Color:      doc.Color,
		LunarDay:   doc.LunarDay,
		LunarMonth: doc.LunarMonth,
		LunarYear:  doc.LunarYear,
		CreatedAt:  doc.CreatedAt,
		UpdatedAt:  doc.UpdatedAt,
	}
}

func ToNoteListResponse(docs []NoteDocument) []NoteResponse {
	result := make([]NoteResponse, len(docs))
	for i, d := range docs {
		result[i] = ToNoteResponse(&d)
	}
	return result
}
