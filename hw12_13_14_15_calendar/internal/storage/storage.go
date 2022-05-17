package storage

import (
	"context"
	"time"
)

var ()

type Storage interface {
	CreateEvent(ctx context.Context, e Event) error
	UpdateEvent(ctx context.Context, e Event) error
	DeleteEvent(ctx context.Context, e Event) error
	GetEvents(ctx context.Context, start, end time.Time) ([]Event, error)
}
