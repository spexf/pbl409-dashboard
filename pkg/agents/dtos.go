package agents

type AgentStoreData struct {
	Host  string `json:"ip"`
	Name  string `json:"name"`
	Group string `json:"groups"`
}

type AgentIds struct {
	AgentIds []string `json:"agents_id"`
}

type AgentKeyResponse struct {
	Error int `json:"error"`
	Data  struct {
		ID  string `json:"id"`
		Key string `json:"key"`
	} `json:"data"`
}

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

type AgentItem struct {
	OS                OSInfo `json:"os"`
	Status            string `json:"status"`
	Version           string `json:"version"`
	DateAdd           string `json:"dateAdd"`
	IP                string `json:"ip"`
	ID                string `json:"id"`
	RegisterIP        string `json:"registerIP"`
	Manager           string `json:"manager"`
	Name              string `json:"name"`
	GroupConfigStatus string `json:"group_config_status"`
	LastKeepAlive     string `json:"lastKeepAlive"`
	StatusCode        int    `json:"status_code"`
	NodeName          string `json:"node_name"`
}

type AgentListResponse struct {
	Data struct {
		AffectedItems []AgentDTO `json:"affected_items"`
	} `json:"data"`
	Error int `json:"error"`
}

type AgentItemResponse struct {
	Data struct {
		AffectedItems []AgentItem `json:"affected_items"`
	} `json:"data"`
}

type DeleteAgentResponseFull struct {
	Data struct {
		AffectedItems []string `json:"affected_items"`
	} `json:"data"`
	Message string `json:"message"`
	Error   int    `json:"error"`
}

type AgentIDRequest struct {
	AgentsID []string `json:"agents_id" validate:"required,dive,required"`
}

type WazuhDeleteResponse struct {
	Data    interface{} `json:"data"` // Bisa diubah jadi struct jika kamu tahu bentuk pastinya
	Message string      `json:"message"`
	Error   int         `json:"error"`
}
