package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Dialog struct {
	TriggerID      string          `json:"trigger_id"`      // Required
	CallbackID     string          `json:"callback_id"`     // Required
	State          string          `json:"state,omitempty"` // Optional
	Title          string          `json:"title"`
	SubmitLabel    string          `json:"submit_label,omitempty"`
	NotifyOnCancel bool            `json:"notify_on_cancel"`
	Elements       []DialogElement `json:"elements"`
}

type DialogElement struct {
	Type        string   `json:"type"`
	Label       string   `json:"label"`
	Name        string   `json:"name"`
	Placeholder string   `json:"placeholder"`
	Optional    bool     `json:"optional"`
	Options     []Option `json:"options"`
}

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type DialogSubmission struct {
	Type            string            `json:"type"`
	Token           string            `json:"token"`
	ActionTimeStamp string            `json:"action_ts"`
	CallbackID      string            `json:"callback_id"`
	ResponseURL     string            `json:"response_url"`
	State           string            `json:"state"`
	Submission      map[string]string `json:"submission"`
	Team            Team              `json:"team"`
	User            User              `json:"user"`
	Channel         Channel           `json:"channel"`
}

// OpenDialog opens a Slack dialog box of the repo request form
func OpenDialog(triggerID string) error {
	token := os.Getenv("SLACK_BOT_USER_OAUTH_ACCESS_TOKEN")

	d, err := getRepoCreateDialog()
	if err != nil {
		return fmt.Errorf("[ERROR] %s", err.Error())
	}

	payload, err := json.Marshal(&d)
	if err != nil {
		return fmt.Errorf("[ERROR] %s", err.Error())
	}

	// convert payload to UTF formatted string
	t := &url.URL{Path: string(payload)}
	encodedPayload := t.String()
	encodedPayload = strings.Trim(encodedPayload, "./")

	urlfmt := "https://slack.com/api/dialog.open?token=%s&trigger_id=%s&dialog=%s"
	url := fmt.Sprintf(urlfmt, token, triggerID, encodedPayload)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("[ERROR] %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")

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

	hasErrors, errMsg, err := containsErrors(string(b))
	if err != nil {
		return fmt.Errorf("[ERROR] %s", err.Error())
	}
	if hasErrors {
		return fmt.Errorf(errMsg)
	}
	return nil
}

func getRepoCreateDialog() (*Dialog, error) {
	absPath, err := filepath.Abs("resources/dialog.json")
	if err != nil {
		fmt.Println("GetRepoCreateDialog ERROR: ", err.Error())
		return nil, err
	}

	b, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println("GetRepoCreateDialog ERROR: ", err.Error())
		return nil, err
	}

	var d Dialog
	err = json.Unmarshal(b, &d)
	if err != nil {
		fmt.Println("GetRepoCreateDialog ERROR: ", err.Error())
		return nil, err
	}

	return &d, nil
}
