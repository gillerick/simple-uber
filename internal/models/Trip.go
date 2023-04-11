package models

import (
	"github.com/bww/go-postgis"
	"math/big"
	"time"
)

type Trip struct {
	TripId        uint64
	RideId        uint64
	StartLocation postgis.PointS
	EndLocation   postgis.PointS
	Distance      int64
	Duration      time.Time
	Fare          big.Float
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
