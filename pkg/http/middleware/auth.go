package middleware

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var token = os.Getenv("VERIFICATION_TOKEN")

// Auth function simply wraps the handler function for verifying incoming requests with a provided verificatio token
func Auth(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		if r.FormValue("token") != "" {
			if token != r.FormValue("token") {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		fn(w, r, param)
	}
}
