package http_server

import (
	"log"
	"net/http"
)

// answerHello ...
func answerHello(w http.ResponseWriter, r *http.Request) {
	log.Println("answerHello()")
	w.Write([]byte("Hello World!"))
}
