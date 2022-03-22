package http_server

import (
	"fmt"
	"net/http"
	"time"
)

func (h *Handlers) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)

		resp := fmt.Sprintf("%s %s %s %s %d %v %s", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, http.StatusOK, time.Since(startTime), r.UserAgent())
		h.logger.Info(resp)
	})
}
