package container

import (
	"github.com/miladabc/tfh-orb/internal/config"
	"github.com/miladabc/tfh-orb/internal/grpc"
	"github.com/miladabc/tfh-orb/internal/log"
)

type Container struct {
	Config     *config.Config
	GrpcServer *grpc.Server
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

	c.InitGrpcServer()

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

func (c *Container) InitGrpcServer() {
	if notNil(c.GrpcServer) {
		return
	}

	c.GrpcServer = grpc.New(c.Config.GrpcServer)
}

func (c *Container) Shutdown() {
	if notNil(c.GrpcServer) {
		c.GrpcServer.Stop()
	}
}

func notNil[T any](p *T) bool {
	return p != nil
}
