package service

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"pbl409-dashboard/pkg/utils"

	"gorm.io/gorm"
)

type ServiceHandler struct {
	DB *gorm.DB
}

func (h *ServiceHandler) GetService(w http.ResponseWriter, r *http.Request) {
	service, err := GetService(h.DB)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Gagal mengambil service")
	}
	w.Header().Set("Content-Type", "application/json")
	utils.RespondWithJSON(w, http.StatusOK, service, "Getting data success")
}

func (h *ServiceHandler) ShowService(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ValidateAndParseIDParam(w, r, "id")
	if !ok {
		return
	}

	service, err := ShowService(h.DB, uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondWithError(w, http.StatusNotFound, "Service not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Getting service data failed")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.RespondWithJSON(w, http.StatusAccepted, service, "Getting data success")
}

func (h *ServiceHandler) StoreService(w http.ResponseWriter, r *http.Request) {
	var input ServiceStore
	if ok := utils.ParseAndValidateJSON(w, r, &input); !ok {
		return
	}

	if err := StoreService(h.DB, input); err != nil {
		log.Println("Store error", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to store service")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, input, "Service created succcessfully")
}

func (h *ServiceHandler) UpdateService(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ValidateAndParseIDParam(w, r, "id")
	if !ok {
		return
	}

	var updated map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if err := Update(h.DB, uint(id), updated); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondWithError(w, http.StatusNotFound, "Service not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Getting service data failed")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, nil, "Service modified successfully")
}

func (h *ServiceHandler) DeleteService(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ValidateAndParseIDParam(w, r, "id")
	if !ok {
		return
	}

	err := Delete(h.DB, uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondWithError(w, http.StatusNotFound, "Service not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Getting service data failed")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.RespondWithJSON(w, http.StatusOK, nil, "Service deletion success")

}
