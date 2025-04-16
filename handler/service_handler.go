package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"pbl409-dashboard/services"
	"pbl409-dashboard/utils"
	"strconv"

	"github.com/gorilla/mux"
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
	if err := json.NewEncoder(w).Encode(service); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Encode fail")
	}
}

func (h *ServiceHandler) ShowService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "ID Tidak Ditemukan")
	}

	if idStr < "0" {
		utils.RespondWithError(w, http.StatusBadRequest, "ID Harus Bernilai Positif")
	}

	id, err := strconv.Atoi(idStr)

	service, err := services.ShowService(h.DB, uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondWithError(w, http.StatusNotFound, "Service tidak ditemukan")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Gagal menggambil data service")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.RespondWithJSON(w, http.StatusAccepted, service)
}

func (h *ServiceHandler) UpdateService(w http.ResponseWriter, r *http.Request) {

}

func (h *ServiceHandler) DeleteService(w http.ResponseWriter, r *http.Request) {

}
