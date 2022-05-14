package internalhttp

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/fenogentov/OTUS-HW-Go/hw12_13_14_15_calendar/internal/logger"
	"github.com/pkg/errors"
)

// Server ...
type Server struct {
	address string
	server  *http.Server
	logger  logger.Logger
}

// Logger ...
type Logger interface { // TODO
}

// Application ...
type Application interface { // TODO
}

// NewServer ...
func NewServer(logger logger.Logger, host, port string) *Server {
	addr := net.JoinHostPort(host, port)
	return &Server{
		address: addr,
		logger:  logger,
	}
}

// Start ...
func (s *Server) Start(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", s.loggingMiddleware(s.answerHello))
	s.server = &http.Server{
		Addr:         s.address,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.logger.Info("http server start")
	err := s.server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	<-ctx.Done()
	s.logger.Info("http server stop")
	return nil
}

// Stop ...
func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return errors.Errorf("http server error stoped: %s", err)
	}
	return nil
}

// answerHello ...
func (s *Server) answerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
