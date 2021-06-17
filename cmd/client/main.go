package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rhc07/chatservice-gRPC-golang/pkg/apis"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error from client: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.SayHello(ctx, &pb.MessageRequest{Body: "Rorie"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Greeting: %s", r.GetBody())

}
