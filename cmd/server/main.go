package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/rhc07/chatservice-gRPC-golang/pkg/apis"
	"github.com/rhc07/chatservice-gRPC-golang/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 500051: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterChatServiceServer(s, &server.Server{})

	fmt.Println("Server is starting, say hello!")

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
