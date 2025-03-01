package grpc

import (
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	cfg  Config
	Grpc *grpc.Server
}

type Config struct {
	Network string `config:"network" validate:"required"`
	Address string `config:"address" validate:"required,hostname_port"`
}

func New(cfg Config) *Server {
	return &Server{
		cfg:  cfg,
		Grpc: grpc.NewServer(),
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen(s.cfg.Network, s.cfg.Address)
	if err != nil {
		return fmt.Errorf("starting listener: %w", err)
	}

	err = s.Grpc.Serve(listener)
	if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		return fmt.Errorf("starting server: %w", err)
	}

	return nil
}

func (s *Server) Stop() {
	s.Grpc.GracefulStop()
}
