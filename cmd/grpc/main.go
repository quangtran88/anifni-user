package main

import (
	"fmt"
	"github.com/joho/godotenv"
	baseConstants "github.com/quangtran88/anifni-base/libs/constants"
	baseUtils "github.com/quangtran88/anifni-base/libs/utils"
	grpcHandler "github.com/quangtran88/anifni-user/adapters/handlers/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	err := godotenv.Load()
	env := baseUtils.GetEnvManager()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", env.GetEnv(baseConstants.GRPCPortEnvKey)))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcHandler.InitGRPCServices(s)

	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
