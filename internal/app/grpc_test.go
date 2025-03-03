package app

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/miladabc/tfh-orb/internal/container"
	orbclient "github.com/miladabc/tfh-orb/internal/orb/client"
	"github.com/miladabc/tfh-orb/internal/orb/model"
	"github.com/stretchr/testify/require"
)

const baseURL = "127.0.0.1:50051"

func TestIntegrationHeartbeat(t *testing.T) {
	t.Run("initialize app", func(t *testing.T) {
		c := container.New()

		t.Cleanup(func() {
			c.Shutdown()
		})

		err := c.Init()
		require.NoError(t, err)

		go func() {
			err := c.GrpcServer.Start()
			require.NoError(t, err)
		}()

		waitForServerStart(t)

		client, err := orbclient.New(baseURL)
		require.NoError(t, err)
		t.Cleanup(func() {
			client.Close()
		})

		ctx := context.Background()
		hb := model.Heartbeat{
			DeviceID:  "8a5456a5-db60-4030-a12a-8927d6174594",
			Latitude:  37.7749,
			Longitude: -122.4194,
			Timestamp: time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
		}

		t.Run("get latest heartbeat request, it should not exist", func(t *testing.T) {
			_, err := client.GetLatestLocation(ctx, hb.DeviceID)
			require.ErrorIs(t, err, orbclient.ErrHeartbeatNotFound)
		})

		t.Run("send heartbeat request", func(t *testing.T) {
			err := client.SendHeartbeat(ctx, hb)
			require.NoError(t, err)
		})

		t.Run("get latest heartbeat request", func(t *testing.T) {
			storedHB, err := client.GetLatestLocation(ctx, hb.DeviceID)
			require.NoError(t, err)
			require.Equal(t, hb, storedHB)
		})
	})
}

func waitForServerStart(t *testing.T) {
	t.Helper()

	require.Eventually(t, func() bool {
		conn, err := net.Dial("tcp", baseURL)
		if err != nil || conn == nil {
			return false
		}

		err = conn.Close()

		return err == nil
	}, time.Second, time.Millisecond)
}
