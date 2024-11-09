package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	FromEmail string     `json:"from_email"`
	ToEmail   string     `json:"to_email"`
	Content   string     `json:"content"`
	IsRead    bool       `json:"is_read" gorm:"default:false"`
	ReadAt    *time.Time `json:"read_at"`
	CreatedAt time.Time  `json:"created_at"`
}
