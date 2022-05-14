package memorystorage

import (
	"fmt"
	"testing"
	"time"

	"github.com/fenogentov/OTUS-HW-Go/hw12_13_14_15_calendar/internal/storage"
)

var testCases = []struct {
	name  string
	event storage.Event
}{
	{
		name: "Create Event One",
		event: storage.Event{
			ID:        1234,
			Title:     "one event",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Second * 30),
			Descript:  "test One event",
			UserID:    "qwerty",
		},
	},
	{
		name: "Create Event Two",
		event: storage.Event{
			ID:        2345,
			Title:     "two event",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Second * 30),
			Descript:  "test Two event",
			UserID:    "qwerty",
		},
	},
	{
		name: "Create Event Three",
		event: storage.Event{
			ID:        3456,
			Title:     "three event",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Second * 30),
			Descript:  "test Three event",
			UserID:    "qwerty",
		},
	},
	{
		name: "Create Event 4",
		event: storage.Event{
			ID:        4567,
			Title:     "4 event",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Second * 30),
			Descript:  "test Three event",
			UserID:    "qwerty",
		},
	},
	{
		name: "Update Event 1",
		event: storage.Event{
			ID:        1234,
			Title:     "5 event",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Second * 30),
			Descript:  "test Three event",
			UserID:    "qwerty",
		},
	},
}

func TestStorageMemory(t *testing.T) {
	t.Run("Storage Memory", func(t *testing.T) {
		tStorage := New()
		tStorage.CreateEvent(storage.Event{})
		tStorage.CreateEvent(testCases[0].event)
		tStorage.CreateEvent(testCases[1].event)
		tStorage.CreateEvent(testCases[2].event)

		tStorage.UpdateEvent(storage.Event{})
		tStorage.UpdateEvent(testCases[3].event)
		tStorage.UpdateEvent(testCases[4].event)

		tStorage.DeleteEvent(storage.Event{})
		tStorage.DeleteEvent(testCases[3].event)
		tStorage.DeleteEvent(testCases[4].event)

		fmt.Printf("%+v\n", tStorage.GetEvents(time.Now(), time.Now().Add(time.Second*35)))
	})
}
