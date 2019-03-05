package slack

type Team struct {
	ID             string `json:"id"`
	Domain         string `json:"domain"`
	EnterpriseID   string `json:"enterprise_id"`
	EnterpriseName string `json:"enterprise_name"`
}
