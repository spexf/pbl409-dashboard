package agents

import (
	"fmt"
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

	agentName, agentOK := utils.ValidateName(w, r, "agentName")
	fmt.Print(agentOK)
	if !agentOK {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed getting agent data")
		return

	}

	agent, err := GetAgentDetails(h.DB, uint(id), agentName)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed getting agent data")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, agent, "Getting agent data success")

}

func (h *AgentHandler) GetAgents(w http.ResponseWriter, r *http.Request) {
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

func (h *AgentHandler) CreateAgents(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ValidateAndParseIDParam(w, r, "id")
	if !ok {
		return
	}

	var input AgentStoreData
	if ok := utils.ParseAndValidateJSON(w, r, &input); !ok {
		return
	}
	create, err := StoreAgents(h.DB, uint(id), input)
	fmt.Print(err)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed storing agent data")
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, create, "Agent data stored")
}

func (h *AgentHandler) DeleteAgents(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ValidateAndParseIDParam(w, r, "id")
	if !ok {
		return
	}

	var input AgentIDRequest
	if ok := utils.ParseAndValidateJSON(w, r, &input); !ok {
		return
	}
	data, err := DeleteAgents(h.DB, uint(id), input.AgentsID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed removing agent data")
		return
	}
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed removing agent data")
		return
	}

	var status int

	errMes, ok := utils.GetDeleteErrorMessage(dataMap)
	if !ok {
		status = http.StatusOK
	}

	if errMes == "Agent does not exist" {
		status = http.StatusNotFound
	}

	utils.RespondWithJSON(w, status, dataMap["data"], dataMap["message"].(string))

}
