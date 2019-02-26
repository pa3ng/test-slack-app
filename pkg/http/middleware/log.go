package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Log function simply wraps the handler function for loging information about the incoming request
func Log(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params), name string) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()
		log.Printf(
			"%s: %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
		fn(w, r, param)
	}
}
