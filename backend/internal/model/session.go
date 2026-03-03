package model

import (
	"time"

	"github.com/google/uuid"
)

type UserSession struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	RefreshToken string    `gorm:"size:500;not null;index" json:"-"`
	DeviceInfo   string    `gorm:"type:jsonb" json:"device_info"`
	IPAddress    string    `gorm:"size:45" json:"ip_address"`
	ExpiresAt    time.Time `gorm:"not null;index" json:"expires_at"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"-"`

	CreatedAt time.Time `json:"created_at"`
}

func (UserSession) TableName() string {
	return "user_sessions"
}
