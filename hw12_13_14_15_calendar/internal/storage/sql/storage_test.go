package sqlstorage

import (
	"fmt"
	"hw12_13_14_15_calendar/internal/logger"
	"hw12_13_14_15_calendar/internal/storage"
	"testing"
	"time"
)

var testCases = []struct {
	name  string
	event storage.Event
}{
	{
		name: "Create Event One",
		event: storage.Event{
			ID:        1230,
			Title:     "one event",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Second * 30),
			Descript:  "test One event",
			UserID:    "qwertyz",
		},
	},
	{
		name: "Create Event Two",
		event: storage.Event{
			ID:        23450,
			Title:     "two event",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Second * 30),
			Descript:  "test Two event",
			UserID:    "qwertyz",
		},
	},
	{
		name: "Create Event Three",
		event: storage.Event{
			ID:        234560,
			Title:     "three event",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Second * 30),
			Descript:  "test Three event",
			UserID:    "qwertyz",
		},
	},
}

func TestStorage(t *testing.T) {
	fmt.Println(" tst ")
	logg := logger.New("logrus.log", "DEBUG")
	t.Run(testCases[0].name, func(t *testing.T) {
		tStorage, _ := New(logg, "root", "12345", "127.0.0.1", "5438", "calendar")
		tStorage.CreateEvent(testCases[0].event)
		tStorage.CreateEvent(testCases[1].event)
		tStorage.UpdateEvent(storage.Event{})
		tStorage.UpdateEvent(testCases[2].event)
		tStorage.DeleteEvent(testCases[0].event)
		// tStorage.DeleteEvent(storage.Event{})
		//		log.Printf("%+v\n", tStorage.GetEvents(time.Now(), time.Now().Add(time.Duration(time.Second*35))))
	})

}
