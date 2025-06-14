package agents

type OSInfo struct {
	Arch     string `json:"arch"`
	Codename string `json:"codename,omitempty"`
	Major    string `json:"major"`
	Minor    string `json:"minor"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Uname    string `json:"uname"`
	Version  string `json:"version"`
}

type AgentDTO struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	IP                string   `json:"ip"`
	RegisterIP        string   `json:"register_ip"`
	Status            string   `json:"status"`
	StatusCode        int      `json:"status_code"`
	Version           string   `json:"version"`
	OS                OSInfo   `json:"os"`
	DateAdd           string   `json:"date_add"`
	LastKeepAlive     string   `json:"last_keep_alive"`
	Manager           string   `json:"manager"`
	NodeName          string   `json:"node_name"`
	Group             []string `json:"group,omitempty"`
	GroupConfigStatus string   `json:"group_config_status"`
	MergedSum         string   `json:"merged_sum,omitempty"`
	ConfigSum         string   `json:"config_sum,omitempty"`
}

type AgentListResponse struct {
	Data struct {
		AffectedItems []AgentDTO `json:"affected_items"`
	} `json:"data"`
	Error int `json:"error"`
}
