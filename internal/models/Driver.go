package models

import (
	"time"
)

type Driver struct {
	DriverId  uint64
	Vehicle   Vehicle   `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Driver) TableName() string {
	return "driver"
}
