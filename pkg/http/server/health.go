package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
