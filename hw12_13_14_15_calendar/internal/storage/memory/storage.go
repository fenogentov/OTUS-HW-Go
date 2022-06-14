package memorystorage

import (
	"context"
	"fmt"
	"sync"
	"time"

	"hw12_13_14_15_calendar/internal/storage"
)

// Storage ...
type Storage struct {
	//	events map[uuid.UUID]storage.Event
	events map[int64]storage.Event
	mu     sync.RWMutex
}

// New ...
func New() *Storage {
	return &Storage{
		//		events: make(map[uuid.UUID]storage.Event),
		events: make(map[int64]storage.Event),
	}
}

// CreateEvent ...
func (s *Storage) CreateEvent(ctx context.Context, event storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[event.ID]; ok {
		return fmt.Errorf("exists event with id=%d", event.ID)
	}

	if m, ok := storage.EnoughData(event); !ok {
		//	m := strings.Join(m, ", ")
		return fmt.Errorf("not enough data: %+v", m)
	}

	s.events[event.ID] = event
	return nil
}

// UpdateEvent ...
func (s *Storage) UpdateEvent(ctx context.Context, event storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if m, ok := storage.EnoughData(event); !ok {
		//	m := strings.Join(m, ", ")
		return fmt.Errorf("not enough data: %+v", m)
	}

	if _, ok := s.events[event.ID]; !ok {
		return fmt.Errorf("no such event id=%d", event.ID)
	}

	s.events[event.ID] = event

	return nil
}

// DeleteEvent ...
func (s *Storage) DeleteEvent(ctx context.Context, event storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[event.ID]; !ok {
		return fmt.Errorf("no such event id=%d", event.ID)
	}

	delete(s.events, event.ID)

	return nil
}

// GetEvents ...
func (s *Storage) GetEvents(ctx context.Context, start, end time.Time) (result []storage.Event, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if (start == time.Time{} || end == time.Time{}) {
		return result, fmt.Errorf("two time values must be specified")
	}

	for _, e := range s.events {
		if e.StartTime.Before(end) && e.EndTime.After(start) {
			result = append(result, e)
		}
	}
	return
}
