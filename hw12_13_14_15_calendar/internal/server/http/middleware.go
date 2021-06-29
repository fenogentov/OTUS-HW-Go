package internalhttp

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func (s *Server) loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO loggin info
		startTime := time.Now()
		clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			clientIP = "\"error get ip client\""
		}
		next.ServeHTTP(w, r)
		resp := fmt.Sprintf("IP:%s [%v] %s %s %s %d %v %s", clientIP, startTime.Format("2006/01/02 15:04:05 MST"), r.Method, r.URL.Path, r.Proto, http.StatusOK, time.Since(startTime), r.UserAgent())
		s.logger.Info(resp)
	})
}
