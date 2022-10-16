package grpcHandler

import (
	userGRPC "github.com/quangtran88/anifni-grpc/user"
	"github.com/quangtran88/anifni-user/adapters/repositories"
	serviceAdapter "github.com/quangtran88/anifni-user/adapters/services"
	"github.com/quangtran88/anifni-user/core/services"
	"google.golang.org/grpc"
)

func InitGRPCServices(s *grpc.Server) {
	userRepo := repositories.NewUserRepository()

	authService := serviceAdapter.NewAuthenticationService()
	tokenService := serviceAdapter.NewTokenService()

	userService := services.NewUserService(userRepo, authService, tokenService)

	userHandler := NewUserHandler(userService)

	userGRPC.RegisterUserServiceServer(s, userHandler)
}
