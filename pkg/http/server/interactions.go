package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/pa3ng/test-slack-app/pkg/slack"
)

var token = os.Getenv("VERIFICATION_TOKEN")

func Interactions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("INTERACTIONS RECEIVED")

	var ds slack.DialogSubmission
	err := json.Unmarshal([]byte(r.FormValue("payload")), &ds)
	if err != nil {
		// Error: there was an error submitting the request
		fmt.Println("[INTERACTIONS ERROR] ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("PAYLOAD: ", r.FormValue("payload"))
	switch ds.Type {
	case "dialog_submission":
		switch ds.CallbackID {
		case "repo-create":
			// TODO - actually create repo
			// fmt.Fprintln(w, "{\"errors\":[{\"name\":\"repo_name\",\"error\":\"Uh-oh.\"}]}")
			// WARNING: EPHEMERAL DOES NOT WORK ON PRIVATE CHANNELS
			err := postEphemeral(ds.Channel.ID, "Yay you did it!", ds.User.ID)
			if err != nil {
				fmt.Println("[INTERACTIONS ERROR] ", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		default:
			fmt.Println("[INTERACTIONS ERROR] Unrecognized callback ID")
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		fmt.Println("[INTERACTIONS ERROR] Unrecognized interactions event")
		w.WriteHeader(http.StatusInternalServerError)
	}

}

type Eph struct {
	ChannelID string `json:"channel"`
	UserID    string `json:"user"`
	Text      string `json:"text"`
}

func postEphemeral(channelID, msg, userID string) error {
	token := os.Getenv("SLACK_BOT_USER_OAUTH_ACCESS_TOKEN")
	url := "https://slack.com/api/chat.postEphemeral"

	eph := &Eph{
		ChannelID: channelID,
		Text:      msg,
		UserID:    userID,
	}
	payload, err := json.Marshal(eph)
	if err != nil {
		return fmt.Errorf("[ERROR] %s", err.Error())
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	if err != nil {
		return fmt.Errorf("[ERROR] %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[ERROR] %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[ERROR] %s", err.Error())
	}
	hasError, message, err := slack.ContainsErrors(string(b))
	if hasError {
		if err != nil {
			return fmt.Errorf("[ERROR] %s", err.Error())
		}
		return fmt.Errorf("[ERROR] Contains error%s", message)
	}
	return nil
}
