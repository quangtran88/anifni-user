package grpcHandler

import (
	"context"
	"github.com/quangtran88/anifni-grpc/user"
	"github.com/quangtran88/anifni-user/core/domain"
	"github.com/quangtran88/anifni-user/core/ports"
)

type UserHandler struct {
	userGRPC.UnimplementedUserServiceServer
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (_ UserHandler) Ping(_ context.Context, _ *userGRPC.PingInput) (*userGRPC.PingResult, error) {
	return &userGRPC.PingResult{Message: "Pong"}, nil
}

func (handler UserHandler) GetUser(_ context.Context, input *userGRPC.GetUserInput) (*userGRPC.GetUserResult, error) {
	user, err := handler.userService.Get(domain.ID(input.GetId()))
	if err != nil {
		return nil, err
	}
	return &userGRPC.GetUserResult{Id: string(user.Id), Name: user.Name}, err
}
