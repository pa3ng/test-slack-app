package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pa3ng/test-slack-app/pkg/processor"
	"github.com/pa3ng/test-slack-app/pkg/slack"
)

func Commands(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusCreated)

	cmd, args, err := processor.ProcessCommand(r.Form["text"])
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	fmt.Println("CMD: ", cmd)
	switch cmd {
	case "repo":
		action := args[0]
		fmt.Println("ARGS: ", args)
		switch action {
		case "create":
			triggerID := r.FormValue("trigger_id")
			err := slack.OpenDialog(triggerID)
			if err != nil {
				fmt.Fprintln(w, err.Error())
				return
			}
			fmt.Println("Repo Creation Dialog Opened!")
		default:
			fmt.Fprintln(w, "Sorry, I do not understand that action.")
			return
		}
	case "help":
		// TODO
		fmt.Fprintln(w, "Help message coming soon!")
		return
	default:
		fmt.Fprintln(w, "Sorry, I do not understand that command.")
		return
	}

	fmt.Fprintln(w, cmd)
}
