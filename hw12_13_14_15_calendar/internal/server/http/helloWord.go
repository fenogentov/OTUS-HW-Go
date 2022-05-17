package http_server

import (
	"net/http"
)

// helloWorld ...
func (s *Server) helloWorld(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("/hello")
	w.Write([]byte("Hello World!"))
}
