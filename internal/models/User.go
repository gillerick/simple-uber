package models

import (
	"time"
)

type User struct {
	UserId      uint64
	Name        string
	Email       string
	PhoneNumber string
	AuthToken   string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
