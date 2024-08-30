package gateway

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func StartGRPCGateway() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := RegisterEventServiceGateway(ctx, mux, "event-service:50051", opts)
	if err != nil {
		log.Fatalf("Failed to start Event gRPC-Gateway: %v", err)
	}

	err = RegisterBookingServiceGateway(ctx, mux, "booking-service:50052", opts)
	if err != nil {
		log.Fatalf("Failed to start Booking gRPC-Gateway: %v", err)
	}

	log.Println("Starting HTTP/1.1 REST server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to serve gRPC-Gateway: %v", err)
	}
}
