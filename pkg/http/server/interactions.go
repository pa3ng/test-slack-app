package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/pa3ng/test-slack-app/pkg/slack"
)

var token = os.Getenv("VERIFICATION_TOKEN")

func Interactions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("INTERACTIONS RECEIVED")

	var ds slack.DialogSubmission
	err := json.Unmarshal([]byte(r.FormValue("payload")), &ds)
	if err != nil {
		fmt.Println("[INTERACTIONS ERROR] ", err.Error())
		return
	}
	fmt.Println("TYPE: ", ds.Type)
	fmt.Printf("SUBMISSION: %+v\n", ds.Submission)
	w.WriteHeader(http.StatusOK)
}
