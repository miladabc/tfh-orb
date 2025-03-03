package controller

import (
	"context"
	"testing"
	"time"

	"github.com/miladabc/tfh-orb/internal/orb/proto"
	"github.com/miladabc/tfh-orb/internal/orb/repo"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestControllerSendHeartbeat(t *testing.T) {
	t.Parallel()

	r := repo.NewMock()
	c := New(r)

	tests := []struct {
		name string
		req  *proto.SendHeartbeatRequest
		res  *proto.SendHeartbeatResponse
		err  string
	}{
		{
			name: "empty device_id",
			req: &proto.SendHeartbeatRequest{
				DeviceId:  "",
				Latitude:  1,
				Longitude: 2,
				Timestamp: timestamppb.New(time.Now()),
			},
			res: nil,
			err: "code = InvalidArgument desc = device_id is required",
		},
		{
			name: "empty timestamp",
			req: &proto.SendHeartbeatRequest{
				DeviceId:  "8a5456a5-db60-4030-a12a-8927d6174594",
				Latitude:  1,
				Longitude: 2,
				Timestamp: nil,
			},
			res: nil,
			err: "code = InvalidArgument desc = timestamp is required",
		},
		{
			name: "valid request",
			req: &proto.SendHeartbeatRequest{
				DeviceId:  "8a5456a5-db60-4030-a12a-8927d6174594",
				Latitude:  1,
				Longitude: 2,
				Timestamp: timestamppb.New(time.Now()),
			},
			res: &proto.SendHeartbeatResponse{
				Success: true,
				Message: "Heartbeat received successfully",
			},
			err: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, err := c.SendHeartbeat(context.Background(), tc.req)
			if tc.err == "" {
				require.Equal(t, tc.res.Success, res.Success)
				require.Equal(t, tc.res.Message, res.Message)
			} else {
				require.ErrorContains(t, err, tc.err)
				require.Equal(t, tc.res, res)
			}
		})
	}
}

func TestControllerGetLatestLocation(t *testing.T) {
	t.Parallel()

	r := repo.NewMock()
	c := New(r)

	tests := []struct {
		name string
		req  *proto.GetLatestLocationRequest
		res  *proto.GetLatestLocationResponse
		err  string
	}{
		{
			name: "empty device_id",
			req: &proto.GetLatestLocationRequest{
				DeviceId: "",
			},
			res: nil,
			err: "code = InvalidArgument desc = device_id is required",
		},
		{
			name: "valid request",
			req: &proto.GetLatestLocationRequest{
				DeviceId: "8a5456a5-db60-4030-a12a-8927d6174594",
			},
			res: &proto.GetLatestLocationResponse{
				Found:     true,
				Latitude:  1,
				Longitude: 2,
				Timestamp: timestamppb.New(time.Date(2020, 1, 1, 12, 0, 0, 0, time.Local)),
			},
			err: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, err := c.GetLatestLocation(context.Background(), tc.req)
			if tc.err == "" {
				require.Equal(t, tc.res.Found, res.Found)
				require.Equal(t, tc.res.Latitude, res.Latitude)
				require.Equal(t, tc.res.Longitude, res.Longitude)
				require.Equal(t, tc.res.Timestamp, res.Timestamp)
			} else {
				require.ErrorContains(t, err, tc.err)
				require.Equal(t, tc.res, res)
			}
		})
	}
}
