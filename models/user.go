package models

import (
	"time"
)

type User struct {
	UserId          uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Email           string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"` // 指定 varchar(255)
	UserName        string    `gorm:"not null" json:"user_name"`
	Password        string    `gorm:"not null" json:"password"`
	ConfirmPassword string    `gorm:"not null" json:"confirm_password"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	IsLogin         bool      `gorm:"default:false" json:"is_login"`
	AvatarURL       string    `gorm:"default:''" json:"avatar_url"`
}
