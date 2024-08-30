package event

import (
	"context"

	eventpb "github.com/tech-and-avinash/go-shared-ms-poc/api/proto/event"
)

type EventService interface {
	CreateEvent(ctx context.Context, name, location string, date *eventpb.Timestamp) (*eventpb.CreateEventResponse, error)
	GetEvent(ctx context.Context, id string) (*eventpb.GetEventResponse, error)
}
