package agents

import (
	"net/http"
	"pbl409-dashboard/pkg/utils"

	"gorm.io/gorm"
)

type AgentHandler struct {
	DB *gorm.DB
}

func (h *AgentHandler) GetAgentData(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ValidateAndParseIDParam(w, r, "id")
	if !ok {
		return
	}
	agentList, err := GetAgents(h.DB, uint(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed getting agent data")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, agentList, "Getting agent data success")
}
