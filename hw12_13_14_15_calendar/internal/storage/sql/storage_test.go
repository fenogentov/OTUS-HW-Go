package sqlstorage

import (
	"context"
	"fmt"

	"testing"
	"time"

	"hw12_13_14_15_calendar/internal/logger"
	"hw12_13_14_15_calendar/internal/storage"
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

func TestStorageSQL(t *testing.T) {
	logg := logger.New("logrus.log", "DEBUG")

	t.Run("Storage SQL", func(t *testing.T) {
		tStorage, _ := New(logg, "127.0.0.1", "5432", "calendar", "root", "12345")
		err := tStorage.Connect(context.Background())
		fmt.Println(err)

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
