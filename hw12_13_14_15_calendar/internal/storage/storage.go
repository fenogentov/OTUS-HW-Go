package storage

import "time"

type Storage interface {
	CreateEvent(e Event) error
	UpdateEvent(e Event) error
	DeleteEvent(e Event) error
	GetEvents(startData, endData time.Time) ([]Event, error)
}
