package slack

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Error struct {
	Ok               bool             `json:"ok"`
	ErrorType        string           `json:"error"`
	ResponseMetadata ResponseMetadata `json:"response_metadata"`
}

type ResponseMetadata struct {
	Messages []string `json:"messages"`
}

func containsErrors(responseBody string) (hasErrors bool, message string, functionError error) {
	var e Error
	err := json.Unmarshal([]byte(responseBody), &e)
	if err != nil {
		return false, "", fmt.Errorf("Could not unmarshall response body to check for errors")
	}
	if !e.Ok {
		errMsg := fmt.Sprintf("Oops, I got a `%s` error from Slack.", e.ErrorType)
		if len(e.ResponseMetadata.Messages) > 0 {
			moreErrs := " Here's more info:\n"
			for _, msg := range e.ResponseMetadata.Messages {
				msg = strings.Trim(msg, "[ERROR]")
				msg = strings.TrimSpace(msg)
				msg = "```" + msg + "```"
				moreErrs = moreErrs + msg + "\n"
			}
			errMsg = errMsg + moreErrs
			errMsg = errMsg + "Please message `@artifactory-admins` in `#artifactory` channel for assistance."
		}
		return true, errMsg, nil
	}
	return false, "", nil
}
