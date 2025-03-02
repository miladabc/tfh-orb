package repo

import (
	"github.com/miladabc/tfh-orb/internal/orb/model"
	"github.com/miladabc/tfh-orb/pkg/memory"
)

type Repository struct {
	mem *memory.Memory[string, model.Heartbeat]
}

func New() *Repository {
	return &Repository{
		mem: memory.New[string, model.Heartbeat](),
	}
}

func (r *Repository) StoreHeartbeat(hb model.Heartbeat) {
	r.mem.Store(hb.DeviceID, hb)
}

func (r *Repository) GetLatestHeartbeat(deviceID string) (model.Heartbeat, bool) {
	return r.mem.Get(deviceID)
}
