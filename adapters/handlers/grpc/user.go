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

func (handler UserHandler) Get(ctx context.Context, input *userGRPC.GetUserInput) (*userGRPC.UserResult, error) {
	user, err := handler.userService.Get(ctx, domain.ID(input.GetId()))
	if err != nil {
		return nil, err
	}
	return &userGRPC.UserResult{Pid: string(user.Pid), Name: user.Name}, err
}

func (handler UserHandler) Create(ctx context.Context, in *userGRPC.CreateUserInput) (*userGRPC.UserResult, error) {
	user, err := handler.userService.Create(ctx, domain.CreateUserInput{
		Pid:       domain.PID(in.Pid),
		Email:     in.Email,
		Password:  in.Password,
		LastName:  in.LastName,
		FirstName: in.FirstName,
	})
	if err != nil {
		return nil, err
	}

	return &userGRPC.UserResult{
		Pid:       string(user.Pid),
		Email:     user.Email,
		Name:      user.Name,
		LastName:  user.LastName,
		FirstName: user.FirstName,
	}, nil
}

func (handler UserHandler) CheckDuplicated(ctx context.Context, in *userGRPC.CheckDuplicatedUserInput) (*userGRPC.CheckDuplicatedUserResult, error) {
	ok, err := handler.userService.CheckDuplicated(ctx, in.Email)
	if err != nil {
		return nil, err
	}
	return &userGRPC.CheckDuplicatedUserResult{Ok: ok}, nil
}
