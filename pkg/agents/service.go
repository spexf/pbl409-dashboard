package agents

import (
	"encoding/json"
	wazuh "pbl409-dashboard/pkg/client"
	service "pbl409-dashboard/pkg/services"
	"pbl409-dashboard/pkg/utils"

	"gorm.io/gorm"
)

type WazuhAgentResponse struct {
	Data struct {
		AffectedItems []WazuhAgentItem `json:"affected_items"`
	} `json:"data"`
}

type WazuhAgentItem struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	IP                string   `json:"ip"`
	RegisterIP        string   `json:"registerIP"`
	Status            string   `json:"status"`
	StatusCode        int      `json:"status_code"`
	Version           string   `json:"version"`
	DateAdd           string   `json:"dateAdd"`
	LastKeepAlive     string   `json:"lastKeepAlive"`
	Manager           string   `json:"manager"`
	NodeName          string   `json:"node_name"`
	Group             []string `json:"group,omitempty"`
	GroupConfigStatus string   `json:"group_config_status"`
	MergedSum         string   `json:"mergedSum,omitempty"`
	ConfigSum         string   `json:"configSum,omitempty"`

	OS struct {
		Arch     string `json:"arch"`
		Codename string `json:"codename,omitempty"`
		Major    string `json:"major"`
		Minor    string `json:"minor"`
		Name     string `json:"name"`
		Platform string `json:"platform"`
		Uname    string `json:"uname"`
		Version  string `json:"version"`
	} `json:"os"`
}

func GetAgents(db *gorm.DB, id uint) (interface{}, error) {
	host, err := service.SetWazuhHost(db, id)
	if err != nil {
		return nil, err
	}
	token, err := utils.GetWazuhToken(host)
	if err != nil {
		return nil, err
	}

	body, err := wazuh.WazuhGet(host, token, "/agents")
	if err != nil {
		return nil, err
	}

	var rawResp AgentListResponse
	if err := json.Unmarshal(body, &rawResp); err != nil {
		return nil, err
	}
	var agents []AgentDTO
	for _, item := range rawResp.Data.AffectedItems {
		agent := AgentDTO{
			ID:                item.ID,
			Name:              item.Name,
			IP:                item.IP,
			RegisterIP:        item.RegisterIP,
			Status:            item.Status,
			StatusCode:        item.StatusCode,
			Version:           item.Version,
			DateAdd:           item.DateAdd,
			LastKeepAlive:     item.LastKeepAlive,
			Manager:           item.Manager,
			NodeName:          item.NodeName,
			Group:             item.Group,
			GroupConfigStatus: item.GroupConfigStatus,
			MergedSum:         item.MergedSum,
			ConfigSum:         item.ConfigSum,
			OS: OSInfo{
				Arch:     item.OS.Arch,
				Codename: item.OS.Codename,
				Major:    item.OS.Major,
				Minor:    item.OS.Minor,
				Name:     item.OS.Name,
				Platform: item.OS.Platform,
				Uname:    item.OS.Uname,
				Version:  item.OS.Version,
			},
		}
		agents = append(agents, agent)
	}
	return agents, nil

}
