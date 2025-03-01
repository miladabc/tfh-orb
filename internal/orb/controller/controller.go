package controller

import (
	"context"
	"time"

	"github.com/miladabc/tfh-orb/internal/orb/proto"
	"github.com/miladabc/tfh-orb/internal/orb/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Controller struct {
	proto.UnimplementedOrbManagerServiceServer
	repo *repo.Repository
}

func New(repo *repo.Repository) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c *Controller) SendHeartbeat(_ context.Context, req *proto.SendHeartbeatRequest) (*proto.SendHeartbeatResponse, error) {
	// TODO: Validate other params
	if req.DeviceId == "" {
		return nil, status.Error(codes.InvalidArgument, "device_id is required")
	}

	timestamp, err := time.Parse(time.RFC3339, req.Timestamp)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid timestamp format: %v", err)
	}

	hb := repo.Heartbeat{
		DeviceID:  req.DeviceId,
		Lat:       req.Latitude,
		Lng:       req.Longitude,
		Timestamp: timestamp,
	}

	c.repo.StoreHeartbeat(hb)

	return &proto.SendHeartbeatResponse{
		Success: true,
		Message: "Heartbeat received successfully",
	}, nil
}

func (c *Controller) GetLatestLocation(_ context.Context, req *proto.GetLatestLocationRequest) (*proto.GetLatestLocationResponse, error) {
	if req.DeviceId == "" {
		return nil, status.Error(codes.InvalidArgument, "device_id is required")
	}

	hb, exist := c.repo.GetLatestHeartbeat(req.DeviceId)

	return &proto.GetLatestLocationResponse{
		Found:     exist,
		Latitude:  hb.Lat,
		Longitude: hb.Lng,
		Timestamp: hb.Timestamp.Format(time.RFC3339),
	}, nil
}
