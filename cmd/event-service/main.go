package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	eventpb "github.com/tech-and-avinash/go-shared-ms-poc/api/proto/event"
	"github.com/tech-and-avinash/go-shared-ms-poc/internal/app/event"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	eventServer := event.NewEventServiceServer()

	eventpb.RegisterEventServiceServer(grpcServer, eventServer)
	reflection.Register(grpcServer)

	log.Println("Starting Event gRPC server on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
