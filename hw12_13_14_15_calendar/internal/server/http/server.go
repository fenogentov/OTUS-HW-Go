package httpserver

import (
	"context"
	"encoding/json"
	"hw12_13_14_15_calendar/internal/storage"
	"hw12_13_14_15_calendar/internal/util/logger"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type Server struct {
	handler    *mux.Router
	storage    storage.Storage
	logger     logger.Logger
	httpServer *http.Server
}

func NewServer(logger logger.Logger, host, port string, storage storage.Storage) *Server {
	addr := net.JoinHostPort(host, port)
	return &Server{
		handler: &mux.Router{},
		storage: storage,
		logger:  logger,
		httpServer: &http.Server{
			Addr:    addr,
			Handler: nil,
		},
	}
}

func (s *Server) RegisterHandlers() {
	router := mux.NewRouter()
	//	router.StrictSlash(true)
	router.HandleFunc("/hello", s.helloWorld).Methods("GET")
	router.HandleFunc("/v1/create", s.createEvent).Methods("POST")
	router.HandleFunc("/v1/update", s.updateEvent).Methods("POST")
	router.HandleFunc("/v1/get", s.getEvents).Methods("GET")
	router.HandleFunc("/v1/delete", s.deleteEvent).Methods("POST")
	router.Use(s.loggingMiddleware)

	s.httpServer.Handler = router

	return
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("http server starting ...")

	s.RegisterHandlers()
	err := s.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("http server error started: " + err.Error())
		return err
	}

	<-ctx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Error("http server error stoped: " + err.Error())
	}

	s.logger.Info("http server stop")

	return
}

// createEvent ...
// curl http://localhost:8081/v1/create -v -H "Content-Type: application/json" -iX "POST" -d '{"id":777, "title":"test curl", "start":"2022-05-16T20:00:00.00+03:00", "end":"2022-05-19T20:00:00.00+03:00", "descript":"event curl create","user": "usr"}'
func (s *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	var event storage.Event
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot decode to struct event: " + err.Error() + "\n"))
		return
	}

	if err = s.storage.CreateEvent(r.Context(), event); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create event in database: " + err.Error() + "\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("event created in database\n"))
}

func (s *Server) updateEvent(w http.ResponseWriter, r *http.Request) {
	var event storage.Event
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot decode to struct event: " + err.Error() + "\n"))
		return
	}

	if err = s.storage.UpdateEvent(r.Context(), event); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to update event in database: " + err.Error() + "\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("event updated in database\n"))
}

func (s *Server) getEvents(w http.ResponseWriter, r *http.Request) {
	var times struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	err := json.NewDecoder(r.Body).Decode(&times)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot decode to struct event: " + err.Error() + "\n"))
		return
	}

	events, err := s.storage.GetEvents(r.Context(), times.Start, times.End)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed getting events from database: " + err.Error() + "\n"))
		return
	}

	if events == nil && err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte("events not found: " + err.Error()))
		return
	}

	result, err := json.Marshal(events)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("cannot execute marshal: " + err.Error()))
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// }
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	var event storage.Event
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("cannot decode to struct event: " + err.Error() + "\n"))
		return
	}

	if err = s.storage.DeleteEvent(r.Context(), event); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("failed to delete event from database: " + err.Error() + "\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("event deleted from database\n"))
}
