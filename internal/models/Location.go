package models

import (
	"github.com/bww/go-postgis"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

type Location struct {
	DriverId    uuid.UUID
	Status      string
	Coordinates postgis.PointS
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
