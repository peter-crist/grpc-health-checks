package main

import (
	"log"

	pb "github.com/peter-crist/grpc-health-checks/protos"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Greeting(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Message{Body: "Leave me alone!"}, nil
}
