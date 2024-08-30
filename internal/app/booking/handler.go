package booking

import (
	"context"
	"errors"

	bookingpb "github.com/tech-and-avinash/go-shared-ms-poc/api/proto/booking"
	eventpb "github.com/tech-and-avinash/go-shared-ms-poc/api/proto/event"
)

type BookingServiceServer struct {
	bookingpb.UnimplementedBookingServiceServer
	bookings    map[string]*bookingpb.GetBookingResponse
	eventClient eventpb.EventServiceClient
}

func NewBookingServiceServer(eventClient eventpb.EventServiceClient) *BookingServiceServer {
	return &BookingServiceServer{
		bookings:    make(map[string]*bookingpb.GetBookingResponse),
		eventClient: eventClient,
	}
}

func (s *BookingServiceServer) CreateBooking(ctx context.Context, req *bookingpb.CreateBookingRequest) (*bookingpb.CreateBookingResponse, error) {
	eventResp, err := s.eventClient.GetEvent(ctx, &eventpb.GetEventRequest{Id: req.EventId})
	if err != nil {
		return nil, errors.New("failed to fetch event details")
	}

	// Convert eventpb.GetEventResponse to bookingpb.Event
	event := &bookingpb.Event{
		Id:       eventResp.Id,
		Name:     eventResp.Name,
		Location: eventResp.Location,
		Date:     eventResp.Date,
	}

	id := "booking_" + generateID()
	booking := &bookingpb.GetBookingResponse{
		Id:     id,
		UserId: req.UserId,
		Event:  event, // Use the newly created bookingpb.Event
	}
	s.bookings[id] = booking
	return &bookingpb.CreateBookingResponse{
		Id:      id,
		Message: "Booking created successfully",
	}, nil
}

func (s *BookingServiceServer) GetBooking(ctx context.Context, req *bookingpb.GetBookingRequest) (*bookingpb.GetBookingResponse, error) {
	booking, exists := s.bookings[req.Id]
	if !exists {
		return nil, errors.New("booking not found")
	}
	return booking, nil
}
