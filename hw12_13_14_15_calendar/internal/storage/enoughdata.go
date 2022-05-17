package storage

import "time"

// enoughData
func EnoughData(event Event) (missing []string, ok bool) {
	if event.ID == 0 {
		missing = append(missing, "id")
	}
	if event.Title == "" {
		missing = append(missing, "title")
	}
	if (event.StartTime == time.Time{}) {
		missing = append(missing, "startTime")
	}
	if (event.EndTime == time.Time{}) {
		missing = append(missing, "endTime")
	}
	if event.User == "" {
		missing = append(missing, "user")
	}

	ok = (len(missing) == 0)

	return
}
