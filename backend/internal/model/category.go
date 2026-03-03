package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Color     string    `gorm:"size:7;default:#3B82F6" json:"color"`
	Icon      string    `gorm:"size:50" json:"icon"`
	Type      string    `gorm:"size:20;not null" json:"type"` // task, note, both
	SortOrder int       `gorm:"default:0" json:"sort_order"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Category) TableName() string {
	return "categories"
}

func (c *Category) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

type Tag struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Name   string    `gorm:"size:50;not null" json:"name"`
	Color  string    `gorm:"size:7;default:#6B7280" json:"color"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"-"`

	CreatedAt time.Time `json:"created_at"`
}

func (Tag) TableName() string {
	return "tags"
}

func (t *Tag) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}
