package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ValidateAndParseIDParam(w http.ResponseWriter, r *http.Request, paramName string) (int, bool) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		RespondWithError(w, http.StatusBadRequest, paramName+" tidak ditemukan")
		return -1, false
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		RespondWithError(w, http.StatusBadRequest, paramName+" harus berupa angka positif")
		return -1, false
	}

	return id, true
}
