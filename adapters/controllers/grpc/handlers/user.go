package grpcHandler

import (
	"context"
	"github.com/quangtran88/anifni-grpc/user"
)

type UserHandler struct {
	userGRPC.UnimplementedUserServiceServer
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (s UserHandler) Ping(ctx context.Context, message *userGRPC.PingInput) (*userGRPC.PingResult, error) {
	return &userGRPC.PingResult{Message: "Pong"}, nil
}
