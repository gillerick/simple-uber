package models

import "github.com/bww/go-postgis"

type Ride struct {
	RideId        uint64
	StartLocation postgis.PointS
	EndLocation   postgis.PointS
	RiderId       uint64
}
