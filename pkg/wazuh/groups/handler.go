package groups

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pbl409-dashboard/pkg/utils"
	"strings"

	"gorm.io/gorm"
)

type GroupHandler struct {
	DB *gorm.DB
}

func (h *GroupHandler) GetGroups(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ValidateAndParseIDParam(w, r, "id")
	if !ok {
		return
	}
	groupList, err := GetGroups(h.DB, uint(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed getting group data")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, groupList, "Getting group data success")
}

func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ValidateAndParseIDParam(w, r, "id")
	if !ok {
		return
	}

	var input GroupStoreData
	if ok := utils.ParseAndValidateJSON(w, r, &input); !ok {
		return
	}
	create, err := StoreGroup(h.DB, uint(id), input)

	fmt.Print(err.Error())
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, `"title"`) {
			var errResp GroupCreateError
			json.Unmarshal([]byte(errStr), &errResp)
			utils.RespondWithError(w, http.StatusBadRequest, errResp.Detail)
			return

		}
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed storing group data")
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, create, "Group data stored")
}
