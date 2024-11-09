package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"column:email;unique"`
	Password  string    `json:"-" gorm:"column:password"`
	Username  string    `json:"username" gorm:"column:username"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}
