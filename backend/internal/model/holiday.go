package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VietnameseHoliday struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	NameEn      string    `gorm:"size:255" json:"name_en"`
	Description string    `gorm:"type:text" json:"description"`

	Day     int  `gorm:"not null" json:"day"`
	Month   int  `gorm:"not null" json:"month"`
	IsLunar bool `gorm:"default:false" json:"is_lunar"`

	Type         string `gorm:"size:30;not null" json:"type"` // public_holiday, cultural, memorial, traditional
	IsDayOff     bool   `gorm:"default:false" json:"is_day_off"`
	DaysOffCount int    `gorm:"default:0" json:"days_off_count"`
}

func (VietnameseHoliday) TableName() string {
	return "vietnamese_holidays"
}

func (h *VietnameseHoliday) BeforeCreate(tx *gorm.DB) error {
	if h.ID == uuid.Nil {
		h.ID = uuid.New()
	}
	return nil
}
