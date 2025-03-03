package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	validator "github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	koanf "github.com/knadh/koanf/v2"
	"github.com/rs/zerolog/log"
)

const (
	keyDelim  = "."
	tag       = "config"
	fileName  = "config.yml"
	envPrefix = "TFH_ORB_"
	envDelim  = "_"
)

func New() (*Config, error) {
	k := koanf.New(keyDelim)

	err := k.Load(structs.Provider(Default, tag), nil)
	if err != nil {
		return nil, fmt.Errorf("loading default config: %w", err)
	}

	err = k.Load(file.Provider(fileName), yaml.Parser())

	switch {
	case errors.As(err, new(*os.PathError)):
		log.Warn().Msg("config file not found, using default config")
	case err != nil:
		return nil, fmt.Errorf("loading config file: %w", err)
	}

	err = k.Load(env.Provider(envPrefix, envDelim, func(envKey string) string {
		return strings.ToLower(strings.TrimPrefix(envKey, envPrefix))
	}), nil)
	if err != nil {
		return nil, fmt.Errorf("loading env vars: %w", err)
	}

	var c Config

	err = k.UnmarshalWithConf("", &c, koanf.UnmarshalConf{Tag: tag})
	if err != nil {
		return nil, fmt.Errorf("unmarshalling config: %w", err)
	}

	err = validator.New().Struct(&c)
	if err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}

	return &c, nil
}
