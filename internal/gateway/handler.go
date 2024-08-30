package gateway

import (
	"context"
	"net/http"

	bookingpb "github.com/tech-and-avinash/go-shared-ms-poc/api/proto/booking"
	eventpb "github.com/tech-and-avinash/go-shared-ms-poc/api/proto/event"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func RegisterHandlers(ctx context.Context, mux *runtime.ServeMux, eventEndpoint, bookingEndpoint string, opts []grpc.DialOption) error {
	if err := eventpb.RegisterEventServiceHandlerFromEndpoint(ctx, mux, eventEndpoint, opts); err != nil {
		return err
	}
	if err := bookingpb.RegisterBookingServiceHandlerFromEndpoint(ctx, mux, bookingEndpoint, opts); err != nil {
		return err
	}
	return nil
}

func StartHTTPServer(mux *runtime.ServeMux) error {
	return http.ListenAndServe(":8080", mux)
}
