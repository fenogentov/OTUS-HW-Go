package sqlstorage

// import (
// 	"fmt"

// 	"hw12_13_14_15_calendar/internal/logger"
// 	"hw12_13_14_15_calendar/internal/storage"
// 	"testing"
// 	"time"
// )

// var testCases = []struct {
// 	name  string
// 	event storage.Event
// }{
// 	{
// 		name: "Create Event One",
// 		event: storage.Event{
// 			ID:          1234,
// 			Title:       "one event",
// 			StartTime:   time.Now(),
// 			EndTime:     time.Now().Add(time.Second * 30),
// 			Description: "test One event",
// 			UserID:      "qwerty",
// 		},
// 	},
// 	{
// 		name: "Create Event Two",
// 		event: storage.Event{
// 			ID:          23455,
// 			Title:       "two event",
// 			StartTime:   time.Now(),
// 			EndTime:     time.Now().Add(time.Second * 30),
// 			Description: "test Two event",
// 			UserID:      "qwerty",
// 		},
// 	},
// 	{
// 		name: "Create Event Three",
// 		event: storage.Event{
// 			ID:          23455,
// 			Title:       "three event",
// 			StartTime:   time.Now(),
// 			EndTime:     time.Now().Add(time.Second * 30),
// 			Description: "test Three event",
// 			UserID:      "qwerty",
// 		},
// 	},
// }

// func TestStorage(t *testing.T) {
// 	logg := logger.New("logrus.log", "DEBUG")
// 	t.Run(testCases[0].name, func(t *testing.T) {

// 		tStorage, _ := New(*logg, "postgres", "Fav660755", "127.0.0.1", "5432", "calendar")
// 		tStorage.CreateEvent(testCases[0].event)
// 		tStorage.CreateEvent(testCases[1].event)
// 		tStorage.UpdateEvent(storage.Event{})
// 		tStorage.UpdateEvent(testCases[2].event)
// 		tStorage.DeleteEvent(testCases[0].event)
// 		tStorage.DeleteEvent(storage.Event{})
// 		fmt.Println(tStorage)
// 		//		fmt.Printf("%+v\n", tStorage.GetEvents(time.Now(), time.Now().Add(time.Duration(time.Second*35))))
// 	})

// }
