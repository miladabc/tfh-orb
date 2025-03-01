package config

import "github.com/miladabc/tfh-orb/internal/log"

var Default = Config{
	Log: log.Config{
		Pretty: true,
		Level:  "trace",
	},
}
