package orb

import (
	"github.com/miladabc/tfh-orb/internal/grpc"
	"github.com/miladabc/tfh-orb/internal/orb/controller"
	"github.com/miladabc/tfh-orb/internal/orb/proto"
	"github.com/miladabc/tfh-orb/internal/orb/repo"
)

func Init(server *grpc.Server) {
	repo := repo.New()
	cnt := controller.New(repo)

	proto.RegisterOrbManagerServiceServer(server.Grpc, cnt)
}
