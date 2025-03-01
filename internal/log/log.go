package log

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Pretty bool   `config:"pretty"`
	Level  string `config:"level" validate:"required"`
}

func Init(cfg Config) error {
	if cfg.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		return fmt.Errorf("parsing log level: %w", err)
	}

	zerolog.SetGlobalLevel(level)

	return nil
}
