package config

import (
	"github.com/miladabc/tfh-orb/internal/grpc"
	"github.com/miladabc/tfh-orb/internal/log"
)

type Config struct {
	Log        log.Config  `config:"log" validate:"required"`
	GrpcServer grpc.Config `config:"grpc-server" validate:"required"`
}
