package grpcHandler

import (
	baseUtils "github.com/quangtran88/anifni-base/libs/utils"
	userGRPC "github.com/quangtran88/anifni-grpc/user"
	"github.com/quangtran88/anifni-user/adapters/repositories"
	"github.com/quangtran88/anifni-user/core/services"
	"google.golang.org/grpc"
)

func InitGRPCServices(s *grpc.Server) {
	userRepo := repositories.NewUserRepository()
	hash := baseUtils.GetHashGenerator()

	userService := services.NewUserService(userRepo, hash)

	userHandler := NewUserHandler(userService)

	userGRPC.RegisterUserServiceServer(s, userHandler)
}
