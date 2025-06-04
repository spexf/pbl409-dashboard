package auth

import (
	"encoding/json"
	"net/http"
	user "pbl409-dashboard/pkg/users"
	"pbl409-dashboard/pkg/utils"

	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input user.LoginDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Format JSON tidak valid")
		return
	}
	token, err := Login(h.DB, input)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Email atau Password salah")
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token}, "Login Success")

}
