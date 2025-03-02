package client

import (
	"context"
	"fmt"

	"github.com/miladabc/tfh-orb/internal/orb/model"
	"github.com/miladabc/tfh-orb/internal/orb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Client struct {
	conn *grpc.ClientConn
	cl   proto.OrbManagerServiceClient
}

var (
	ErrHeartbeatStoreFailed = fmt.Errorf("heartbeat store failed")
	ErrHeartbeatNotFound    = fmt.Errorf("heartbeat not found")
)

func New(target string) (*Client, error) {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("creating grpc client: %w", err)
	}

	return &Client{
		conn: conn,
		cl:   proto.NewOrbManagerServiceClient(conn),
	}, nil
}

func (c *Client) SendHeartbeat(ctx context.Context, beat model.Heartbeat) error {
	res, err := c.cl.SendHeartbeat(ctx, &proto.SendHeartbeatRequest{
		DeviceId:  beat.DeviceID,
		Latitude:  beat.Latitude,
		Longitude: beat.Longitude,
		Timestamp: timestamppb.New(beat.Timestamp),
	})
	if err != nil {
		return fmt.Errorf("sending heartbeat: %w", err)
	}

	if !res.Success {
		return ErrHeartbeatStoreFailed
	}

	return nil
}

func (c *Client) GetLatestLocation(ctx context.Context, deviceID string) (model.Heartbeat, error) {
	res, err := c.cl.GetLatestLocation(ctx, &proto.GetLatestLocationRequest{
		DeviceId: deviceID,
	})
	if err != nil {
		return model.Heartbeat{}, fmt.Errorf("getting latest location: %w", err)
	}

	if !res.Found {
		return model.Heartbeat{}, ErrHeartbeatNotFound
	}

	return model.Heartbeat{
		DeviceID:  deviceID,
		Latitude:  res.Latitude,
		Longitude: res.Longitude,
		Timestamp: res.Timestamp.AsTime(),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
