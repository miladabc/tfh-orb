package app

import (
	"context"
	"time"

	"github.com/miladabc/tfh-orb/internal/container"
	"github.com/miladabc/tfh-orb/internal/orb/client"
	"github.com/miladabc/tfh-orb/internal/orb/model"
	"github.com/rs/zerolog/log"
)

func SendHeartbeatRequest() error {
	c := container.New()
	defer c.Shutdown()

	err := c.InitConfig()
	if err != nil {
		return err
	}

	err = c.InitLogger()
	if err != nil {
		return err
	}

	client, err := client.New("localhost:50051")
	if err != nil {
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	deviceID := "8a5456a5-db60-4030-a12a-8927d6174594"

	err = client.SendHeartbeat(ctx, model.Heartbeat{
		DeviceID:  deviceID,
		Latitude:  37.7749,
		Longitude: -122.4194,
		Timestamp: time.Now(),
	})
	if err != nil {
		return err
	}

	log.Info().Msg("sent heartbeat response")

	getLatestLocationResponse, err := client.GetLatestLocation(ctx, deviceID)
	if err != nil {
		return err
	}

	log.Info().Any("res", getLatestLocationResponse).Msg("get latest location response")

	return nil
}
