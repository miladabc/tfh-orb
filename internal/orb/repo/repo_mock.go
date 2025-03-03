package repo

import (
	"time"

	"github.com/miladabc/tfh-orb/internal/orb/model"
)

type MockRepository struct{}

func NewMock() *MockRepository {
	return &MockRepository{}
}

func (r *MockRepository) StoreHeartbeat(_ model.Heartbeat) {}

func (r *MockRepository) GetLatestHeartbeat(deviceID string) (model.Heartbeat, bool) {
	return model.Heartbeat{
		DeviceID:  deviceID,
		Latitude:  1,
		Longitude: 2,
		Timestamp: time.Date(2020, 1, 1, 12, 0, 0, 0, time.Local),
	}, true
}
