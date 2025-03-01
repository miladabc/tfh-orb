package config

import (
	"github.com/miladabc/tfh-orb/internal/grpc"
	"github.com/miladabc/tfh-orb/internal/log"
)

var Default = Config{
	Log: log.Config{
		Pretty: true,
		Level:  "trace",
	},
	GrpcServer: grpc.Config{
		Network: "tcp",
		Address: "0.0.0.0:50051",
	},
}
