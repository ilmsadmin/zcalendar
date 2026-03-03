package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Email        string    `gorm:"uniqueIndex;size:255;not null" json:"email"`
	Username     string    `gorm:"uniqueIndex;size:100;not null" json:"username"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	FullName     string    `gorm:"size:255" json:"full_name"`
	AvatarURL    string    `gorm:"size:500" json:"avatar_url"`
	Timezone     string    `gorm:"size:50;default:Asia/Ho_Chi_Minh" json:"timezone"`
	Locale       string    `gorm:"size:10;default:vi-VN" json:"locale"`
	Role         string    `gorm:"size:20;default:user" json:"role"` // admin, user

	// Preferences
	DefaultCalendarView string `gorm:"size:20;default:month" json:"default_calendar_view"`
	ShowLunar           bool   `gorm:"default:true" json:"show_lunar"`
	ShowGoodBadDays     bool   `gorm:"default:true" json:"show_good_bad_days"`
	Theme               string `gorm:"size:20;default:light" json:"theme"`

	// OAuth
	GoogleID   string `gorm:"size:255" json:"-"`
	FacebookID string `gorm:"size:255" json:"-"`

	// Status
	IsActive    bool       `gorm:"default:true" json:"is_active"`
	IsVerified  bool       `gorm:"default:false" json:"is_verified"`
	LastLoginAt *time.Time `json:"last_login_at"`

	// Timestamps
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "users"
}

// BeforeCreate hook to generate UUID
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
