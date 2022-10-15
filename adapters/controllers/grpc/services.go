package grpcHandler

import (
	userGRPC "github.com/quangtran88/anifni-grpc/user"
	"github.com/quangtran88/anifni-user/adapters/controllers/grpc/handlers"
	"google.golang.org/grpc"
)

func InitGRPCServices(s *grpc.Server) {
	userHandler := grpcHandler.NewUserHandler()

	userGRPC.RegisterUserServiceServer(s, userHandler)
}
