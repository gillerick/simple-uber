package models

import "time"

type Vehicle struct {
	VehicleId uint64
	Make      string
	Model     string
	Plate     string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
