package config

import "github.com/miladabc/tfh-orb/internal/log"

type Config struct {
	Log log.Config `config:"log" validate:"required"`
}
