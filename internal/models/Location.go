package models

import (
	"github.com/bww/go-postgis"
	"time"
)

type Location struct {
	LocationId  uint64
	Coordinates postgis.PointS
	Address     string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
