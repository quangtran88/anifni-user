package grpcHandler

import (
	userGRPC "github.com/quangtran88/anifni-grpc/user"
	"github.com/quangtran88/anifni-user/adapters/controllers/grpc/handlers"
	"github.com/quangtran88/anifni-user/adapters/repositories"
	"github.com/quangtran88/anifni-user/core/services"
	"google.golang.org/grpc"
)

func InitGRPCServices(s *grpc.Server) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := grpcHandler.NewUserHandler(userService)

	userGRPC.RegisterUserServiceServer(s, userHandler)
}
