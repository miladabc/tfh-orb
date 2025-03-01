package cmd

import (
	"os"

	"github.com/miladabc/tfh-orb/internal/app"
	"github.com/rs/zerolog/log"
	cli "github.com/urfave/cli/v2"
)

func Execute() {
	cmd := &cli.App{
		Name:  "orb",
		Usage: "Orb Fleet Management Service",
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "run grpc server",
				Action: func(*cli.Context) error {
					return app.StartGrpcServer()
				},
			},
		},
	}

	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("running cmd")
	}
}
