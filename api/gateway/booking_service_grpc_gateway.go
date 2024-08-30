package gateway

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	bookingpb "github.com/tech-and-avinash/go-shared-ms-poc/api/proto/booking"
)

func RegisterBookingServiceGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return bookingpb.RegisterBookingServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
