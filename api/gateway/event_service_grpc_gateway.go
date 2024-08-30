package gateway

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	eventpb "github.com/tech-and-avinash/go-shared-ms-poc/api/proto/event"
)

func RegisterEventServiceGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return eventpb.RegisterEventServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
