package app

import (
	"github.com/miladabc/tfh-orb/internal/container"
	"github.com/rs/zerolog/log"
)

func StartGrpcServer() error {
	c := container.New()
	defer c.Shutdown()

	err := c.Init()
	if err != nil {
		return err
	}

	go func() {
		log.Info().Msgf("grpc server listening on `%s`", c.Config.GrpcServer.Address)

		err := c.GrpcServer.Start()
		if err != nil {
			log.Fatal().Err(err).Msg("starting grpc server")
		}
	}()

	<-handleInterrupts()

	log.Info().Msgf("grpc server shutting down...")

	return nil
}
