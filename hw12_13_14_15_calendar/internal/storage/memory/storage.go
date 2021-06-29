package memorystorage

import (
	"sync"
	"time"

	"hw12_13_14_15_calendar/internal/storage"
)

//Storage ...
type Storage struct {
	//	events map[uuid.UUID]storage.Event
	events map[int]storage.Event
	mu     sync.RWMutex
}

// New ...
func New() *Storage {
	return &Storage{
		//		events: make(map[uuid.UUID]storage.Event),
		events: make(map[int]storage.Event),
	}
}

// CreateEvent ...
func (s *Storage) CreateEvent(evnt storage.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()

	//	evnt.ID = uuid.New()
	//	s.events[evnt.ID] = evnt

	s.events[evnt.ID] = evnt
}

// UpdateEvent ...
func (s *Storage) UpdateEvent(evnt storage.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events[evnt.ID] = evnt
}

// DeleteEvent ...
func (s *Storage) DeleteEvent(evnt storage.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.events, evnt.ID)
}

// GetEvents ...
func (s *Storage) GetEvents(startDT, endDT time.Time) []storage.Event {
	s.mu.Lock()
	defer s.mu.Unlock()

	var events []storage.Event
	for _, e := range s.events {
		if e.StartTime.Second() >= startDT.Second() && e.EndTime.Second() <= endDT.Second() {
			events = append(events, e)
		}
	}

	return events
}

// TODO
