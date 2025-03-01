package container

import (
	"context"

	"github.com/miladabc/tfh-orb/internal/config"
	"github.com/miladabc/tfh-orb/internal/log"
)

type Container struct {
	Config *config.Config
}

func New() *Container {
	return &Container{}
}

func (c *Container) Init() error {
	err := c.InitConfig()
	if err != nil {
		return err
	}

	err = c.InitLogger()
	if err != nil {
		return err
	}

	return nil
}

func (c *Container) InitConfig() (err error) {
	if notNil(c.Config) {
		return
	}

	c.Config, err = config.New()

	return
}

func (c *Container) InitLogger() error {
	return log.Init(c.Config.Log)
}

func (c *Container) Shutdown(ctx context.Context) {
}

func notNil[T any](p *T) bool {
	return p != nil
}
