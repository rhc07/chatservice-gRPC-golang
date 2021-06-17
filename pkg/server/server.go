package server

import (
	"context"
	"log"

	pb "github.com/rhc07/chatservice-gRPC-golang/pkg/apis"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, input *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("Received: %v\n", input.GetBody())

	return &pb.MessageResponse{
		Body: "Hello " + input.GetBody(),
	}, nil
}
