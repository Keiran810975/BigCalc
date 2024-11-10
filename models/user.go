package models

import (
	"time"
)

type User struct {
	UserId          uint           `gorm:"primaryKey;autoIncrement"`// 主键
	Email           string         `gorm:"uniqueIndex;not null"` 
	UserName        string         `gorm:"not null"`                        
	Password        string         `gorm:"not null"`                        
	ConfirmPassword string         `gorm:"not null"`                      
	CreatedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	IsLogin         bool            `gorm:"default:false"`
	AvatarURL   	string    		`gorm:"default:''"`// 头像URL
}