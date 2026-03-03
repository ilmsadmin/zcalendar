package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reminder struct {
	ID     uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	TaskID *uuid.UUID `gorm:"type:uuid;index" json:"task_id"`

	Title    string    `gorm:"size:255;not null" json:"title"`
	Message  string    `gorm:"type:text" json:"message"`
	RemindAt time.Time `gorm:"not null;index" json:"remind_at"`

	Type   string     `gorm:"size:20;default:push" json:"type"` // push, email, both
	IsSent bool       `gorm:"default:false" json:"is_sent"`
	SentAt *time.Time `json:"sent_at"`

	IsRecurring    bool    `gorm:"default:false" json:"is_recurring"`
	RecurrenceRule *string `gorm:"size:255" json:"recurrence_rule"`

	// Relations
	User User  `gorm:"foreignKey:UserID" json:"-"`
	Task *Task `gorm:"foreignKey:TaskID" json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Reminder) TableName() string {
	return "reminders"
}

func (r *Reminder) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
