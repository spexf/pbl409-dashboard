package user

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"pbl409-dashboard/pkg/utils"

	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	service, err := GetUser(h.DB)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Gagal mengambil service")
	}
	w.Header().Set("Content-Type", "application/json")
	utils.RespondWithJSON(w, http.StatusOK, service, "Getting data success")
}

func (h *UserHandler) ShowUser(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ValidateAndParseIDParam(w, r, "id")
	if !ok {
		return
	}

	service, err := ShowUser(h.DB, uint(id))
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

func (h *UserHandler) StoreUser(w http.ResponseWriter, r *http.Request) {
	var input UserStore
	if ok := utils.ParseAndValidateJSON(w, r, &input); !ok {
		return
	}

	if err := StoreUser(h.DB, input); err != nil {
		log.Println("Store error", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to store service")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, input, "Service created succcessfully")
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ValidateAndParseIDParam(w, r, "id")
	if !ok {
		return
	}

	var updated map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if err := UpdateUser(h.DB, uint(id), updated); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondWithError(w, http.StatusNotFound, "Service not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Getting service data failed")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, nil, "Service modified successfully")
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
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
