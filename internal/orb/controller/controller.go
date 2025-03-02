package controller

import (
	"context"

	"github.com/miladabc/tfh-orb/internal/orb/model"
	"github.com/miladabc/tfh-orb/internal/orb/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Controller struct {
	proto.UnimplementedOrbManagerServiceServer
	repo Repository
}

type Repository interface {
	StoreHeartbeat(hb model.Heartbeat)
	GetLatestHeartbeat(deviceID string) (model.Heartbeat, bool)
}

func New(repo Repository) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c *Controller) SendHeartbeat(_ context.Context, req *proto.SendHeartbeatRequest) (*proto.SendHeartbeatResponse, error) {
	if req.DeviceId == "" {
		return nil, status.Error(codes.InvalidArgument, "device_id is required")
	}

	if req.Timestamp == nil {
		return nil, status.Error(codes.InvalidArgument, "timestamp is required")
	}

	hb := model.Heartbeat{
		DeviceID:  req.DeviceId,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Timestamp: req.Timestamp.AsTime(),
	}

	c.repo.StoreHeartbeat(hb)

	log.Info().Any("beat", hb).Msg("new heartbeat received")

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
		Latitude:  hb.Latitude,
		Longitude: hb.Longitude,
		Timestamp: timestamppb.New(hb.Timestamp),
	}, nil
}
