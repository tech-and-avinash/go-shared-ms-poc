package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	bookingpb "github.com/tech-and-avinash/go-shared-ms-poc/api/proto/booking"
	eventpb "github.com/tech-and-avinash/go-shared-ms-poc/api/proto/event"
	"github.com/tech-and-avinash/go-shared-ms-poc/internal/app/booking"
)

func main() {
	// Set up a connection to the Event service.
	conn, err := grpc.Dial("event-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Event service: %v", err)
	}
	defer conn.Close()

	eventClient := eventpb.NewEventServiceClient(conn)

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	grpcServer := grpc.NewServer()
	bookingServer := booking.NewBookingServiceServer(eventClient)

	bookingpb.RegisterBookingServiceServer(grpcServer, bookingServer)
	reflection.Register(grpcServer)

	log.Println("Starting Booking gRPC server on port 50052...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
