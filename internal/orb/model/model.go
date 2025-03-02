package model

import "time"

type Heartbeat struct {
	DeviceID  string
	Lat       float64
	Lng       float64
	Timestamp time.Time
}
