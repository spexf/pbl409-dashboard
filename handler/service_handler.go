package handler

import (
	"errors"
	"log"
	"net/http"
	"pbl409-dashboard/dtos"
	"pbl409-dashboard/services"
	"pbl409-dashboard/utils"

	"gorm.io/gorm"
)

type ServiceHandler struct {
	DB *gorm.DB
}

func (h *ServiceHandler) GetService(w http.ResponseWriter, r *http.Request) {
	service, err := services.GetService(h.DB)
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

	service, err := services.ShowService(h.DB, uint(id))
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
	var input dtos.ServiceStore
	if ok := utils.ParseAndValidateJSON(w, r, &input); !ok {
		return
	}

	if err := services.StoreService(h.DB, input); err != nil {
		log.Println("Store error", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to store service")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, input, "Service created succcessfully")
}

func (h *ServiceHandler) UpdateService(w http.ResponseWriter, r *http.Request) {

}

func (h *ServiceHandler) DeleteService(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ValidateAndParseIDParam(w, r, "id")
	if !ok {
		return
	}

	err := services.DeleteService(h.DB, uint(id))
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
