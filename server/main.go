package main

import (
	"log"
	"net"

	"github.com/peter-crist/grpc-health-checks/protos"
	"google.golang.org/grpc"
)

const (
	port     = ":5000"
	hostname = "AngryServer"
)

func main() {
	startGrpcServer()
}

func startGrpcServer() {
	//open tcp listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("❌ failed to listen: %v", err)
	}

	log.Println("🥳 Server is listening on port", port)
	log.Printf("💻 Server's name is %s", hostname)

	s := Server{}

	grpcServer := grpc.NewServer()
	protos.RegisterAngryServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ failed to serve: %v", err)
	}
}
