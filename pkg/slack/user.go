package slack

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	TeamID string `json:"team_id"`
}
