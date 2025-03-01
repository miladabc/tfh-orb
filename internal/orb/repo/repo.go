package repo

import (
	"time"

	"github.com/miladabc/tfh-orb/pkg/memory"
)

type Repository struct {
	mem *memory.Memory[string, Heartbeat]
}

type Heartbeat struct {
	DeviceID  string
	Lat       float64
	Lng       float64
	Timestamp time.Time
}

func New() *Repository {
	return &Repository{
		mem: memory.New[string, Heartbeat](),
	}
}

func (r *Repository) StoreHeartbeat(hb Heartbeat) {
	r.mem.Store(hb.DeviceID, hb)
}

func (r *Repository) GetLatestHeartbeat(deviceID string) (Heartbeat, bool) {
	return r.mem.Get(deviceID)
}
