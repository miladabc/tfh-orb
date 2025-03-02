package model

import "time"

type Heartbeat struct {
	DeviceID  string
	Latitude  float64
	Longitude float64
	Timestamp time.Time
}
