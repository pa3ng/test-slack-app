package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pa3ng/test-slack-app/pkg/command"
)

func Commands(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, command.Create(r.FormValue("text")))
}
