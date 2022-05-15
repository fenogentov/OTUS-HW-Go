package memorystorage

import (
	"errors"
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
func (s *Storage) CreateEvent(evnt storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if (evnt == storage.Event{}) {
		return errors.New("error CreateEvent: empty event")
	}

	//	evnt.ID = uuid.New()
	//	s.events[evnt.ID] = evnt

	s.events[evnt.ID] = evnt
	return nil
}

// UpdateEvent ...
func (s *Storage) UpdateEvent(evnt storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if (evnt == storage.Event{}) {
		return errors.New("error UpdateEvent: empty event")
	}

	if _, ok := s.events[evnt.ID]; !ok {
		return errors.New("error UpdateEvent: defunct event")
	}

	s.events[evnt.ID] = evnt

	return nil
}

// DeleteEvent ...
func (s *Storage) DeleteEvent(evnt storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if (evnt == storage.Event{}) {
		return errors.New("error DeleteEvent: empty event")
	}

	delete(s.events, evnt.ID)

	return nil
}

// GetEvents ...
func (s *Storage) GetEvents(startDT, endDT time.Time) ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var events []storage.Event
	for _, e := range s.events {
		if e.StartTime.Second() >= startDT.Second() && e.EndTime.Second() <= endDT.Second() {
			events = append(events, e)
		}
	}
	return events, nil
}
