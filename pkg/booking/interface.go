package booking

import (
	"context"

	bookingpb "github.com/tech-and-avinash/go-shared-ms-poc/api/proto/booking"
)

type BookingService interface {
	CreateBooking(ctx context.Context, userID, eventID string) (*bookingpb.CreateBookingResponse, error)
	GetBooking(ctx context.Context, id string) (*bookingpb.GetBookingResponse, error)
}
