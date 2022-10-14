package main

import (
	"context"
	"fmt"
	userGRPC "github.com/quangtran88/anifni-grpc/user"

	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	userGRPC.UnimplementedUserServiceServer
}

func (s server) Ping(ctx context.Context, message *userGRPC.PingMessage) (*userGRPC.PingResult, error) {
	return &userGRPC.PingResult{Message: "Pong"}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 6000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	userGRPC.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
