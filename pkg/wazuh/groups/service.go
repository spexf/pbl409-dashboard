package groups

import (
	"encoding/json"
	"fmt"
	service "pbl409-dashboard/pkg/services"
	"pbl409-dashboard/pkg/utils"
	wazuh "pbl409-dashboard/pkg/wazuh/client"

	"gorm.io/gorm"
)

func GetGroups(db *gorm.DB, id uint) (interface{}, error) {
	host, err := service.SetWazuhHost(db, id)
	if err != nil {
		return nil, err
	}
	token, err := utils.GetWazuhToken(host)
	if err != nil {
		return nil, err
	}

	body, err := wazuh.WazuhGet(host, token, "/groups")
	if err != nil {
		return nil, err
	}
	var rawResp GroupResponse
	if err := json.Unmarshal(body, &rawResp); err != nil {
		return nil, err
	}
	if len(rawResp.Data.AffectedItems) == 0 {
		return nil, fmt.Errorf("group tidak ada")
	}
	return rawResp.Data.AffectedItems, nil
}

func StoreGroup(db *gorm.DB, wazuh_id uint, group GroupStoreData) (interface{}, error) {
	host, err := service.SetWazuhHost(db, wazuh_id)
	if err != nil {
		return nil, err
	}
	token, err := utils.GetWazuhToken(host)
	if err != nil {
		return nil, err
	}

	payload := map[string]interface{}{
		"group_id": group.GroupID,
	}

	fmt.Print(payload)
	body, err := wazuh.WazuhPost(host, token, "/groups", payload)
	if err != nil {
		var errResp GroupCreateError
		if errval := json.Unmarshal([]byte(err.Error()), &errResp); errval != nil {
			return nil, err
		}
		return nil, err
	}

	fmt.Print(err)

	var result GroupCreatedResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Message, err

}

func DeleteGroup(db *gorm.DB, wazuh_id uint, group uint) (interface{}, error) {
	host, err := service.SetWazuhHost(db, wazuh_id)
	if err != nil {
		return nil, err
	}
	token, err := utils.GetWazuhToken(host)
	if err != nil {
		return nil, err
	}
	body, err := wazuh.WazuhDelete(host, token, "/groups")
	if err != nil {
		var errResp GroupCreateError
		if errval := json.Unmarshal([]byte(err.Error()), &errResp); errval != nil {
			return nil, err
		}
		return nil, err
	}

	return body, err

}
