package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

type User struct {
	UserId      uuid.UUID
	Name        string
	Email       string
	PhoneNumber string
	AuthToken   string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
