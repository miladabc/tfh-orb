package repo

import (
	"testing"
	"time"

	"github.com/miladabc/tfh-orb/internal/orb/model"
	"github.com/stretchr/testify/require"
)

func TestRepoStoreAndGetHeartbeat(t *testing.T) {
	t.Parallel()

	r := New()

	beat := model.Heartbeat{
		DeviceID:  "8a5456a5-db60-4030-a12a-8927d6174594",
		Latitude:  37.7749,
		Longitude: -122.4194,
		Timestamp: time.Now(),
	}

	r.StoreHeartbeat(beat)

	hb, exists := r.GetLatestHeartbeat(beat.DeviceID)
	require.True(t, exists)
	require.Equal(t, beat, hb)
}
