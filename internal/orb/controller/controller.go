package controller

import (
	"context"

	"github.com/miladabc/tfh-orb/internal/orb/model"
	"github.com/miladabc/tfh-orb/internal/orb/proto"
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

	c.repo.StoreHeartbeat(model.Heartbeat{
		DeviceID:  req.DeviceId,
		Lat:       req.Latitude,
		Lng:       req.Longitude,
		Timestamp: req.Timestamp.AsTime(),
	})

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
		Timestamp: timestamppb.New(hb.Timestamp),
	}, nil
}
