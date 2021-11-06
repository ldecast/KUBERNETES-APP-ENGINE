package main

import (
	"fmt"
	"log"
	"net"
	"server/games"

	"google.golang.org/grpc"
)

const server_PORT = 9000

func main() {

	fmt.Println("Go gRPC Server for API Kafka!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", server_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := games.Server{}

	grpcServer := grpc.NewServer()

	games.RegisterGameServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
