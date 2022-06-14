package httpserver

// import (
// 	"encoding/json"
// 	"hw12_13_14_15_calendar/internal/storage"
// 	"net/http"
// 	"strings"
// 	"time"
// )

// // http_api ...
// func (s *Server) http_api(w http.ResponseWriter, r *http.Request) {
// 	s.logger.Info("handler api v1")
// 	path := strings.Trim(r.URL.Path, "/")
// 	pathParts := strings.Split(path, "/")
// 	if len(pathParts) < 2 {
// 		s.logger.Error("expect /v1/<metod> in task handler")
// 		http.Error(w, "expect /v1/<metod> in task handler", http.StatusBadRequest)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
// 	var event storage.Event
// 	err := json.NewDecoder(r.Body).Decode(&event)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		_, err = w.Write([]byte("cannot decode to struct: " + err.Error()))
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	switch pathParts[1] {
// 	case "create":
// 		if err = s.storage.CreateEvent(event); err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)

// 			searchErrorJson, _ := json.Marshal(time.Now())
// 			w.Write(searchErrorJson)

// 			_, err = w.Write([]byte(err.Error()))
// 			if err != nil {
// 				w.WriteHeader(http.StatusInternalServerError)
// 			}
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)

// 		searchErrorJson, _ := json.Marshal(time.Now())
// 		w.Write(searchErrorJson)

// 		_, err = w.Write([]byte("Event created"))
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 	case "update":
// 		if err = s.storage.UpdateEvent(event); err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			_, err = w.Write([]byte(err.Error()))
// 			if err != nil {
// 				w.WriteHeader(http.StatusInternalServerError)
// 			}
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		_, err = w.Write([]byte("Event updated"))
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 	case "delete":
// 		if err = s.storage.DeleteEvent(event); err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			_, err = w.Write([]byte(err.Error()))
// 			if err != nil {
// 				w.WriteHeader(http.StatusInternalServerError)
// 			}
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		_, err = w.Write([]byte("Event deleted"))
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 	case "get":
// 		get, err := s.storage.GetEvents(event.StartTime, event.EndTime)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		jsonGet, err := json.Marshal(get)
// 		if err != nil {
// 			return
// 		}
// 		w.Write(jsonGet)

// 	default:
// 		s.logger.Error("expect /v1/<metod> in task handler")
// 		http.Error(w, "expect /v1/<metod> in task handler", http.StatusBadRequest)
// 		return
// 	}

// 	s.logger.Info("/v1/" + pathParts[1])
// }

// func createEvent(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
// 	var event storage.Event
// 	err := json.NewDecoder(r.Body).Decode(&event)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		_, err = w.Write([]byte("cannot decode to struct: " + err.Error()))
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	if err = s.storage.CreateEvent(event); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)

// 		searchErrorJson, _ := json.Marshal(time.Now())
// 		w.Write(searchErrorJson)

// 		_, err = w.Write([]byte(err.Error()))
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 		}
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)

// 	searchErrorJson, _ := json.Marshal(time.Now())
// 	w.Write(searchErrorJson)

// 	_, err = w.Write([]byte("Event created"))
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }

// func updateEvent(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
// 	var event storage.Event
// 	err := json.NewDecoder(r.Body).Decode(&event)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		_, err = w.Write([]byte("cannot decode to struct: " + err.Error()))
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	if err = s.storage.UpdateEvent(event); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		_, err = w.Write([]byte(err.Error()))
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 		}
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	_, err = w.Write([]byte("Event updated"))
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }

// func getEvents
