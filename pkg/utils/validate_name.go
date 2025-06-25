package utils

import (
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func ValidateName(w http.ResponseWriter, r *http.Request, paramName string) (string, bool) {
	vars := mux.Vars(r)
	name := vars[paramName]

	if name == "" {
		RespondWithError(w, http.StatusBadRequest, paramName+" tidak ditemukan atau kosong")
		return "", false
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(name) {
		RespondWithError(w, http.StatusBadRequest, paramName+" mengandung karakter tidak valid")
		return "", false
	}

	return name, true
}
