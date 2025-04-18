package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ParseAndValidateJSON(w http.ResponseWriter, r *http.Request, dst interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		log.Println("Decode error:", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return false
	}
	if err := validate.Struct(dst); err != nil {
		log.Println("Validation error:", err)
		RespondWithError(w, http.StatusBadRequest, "Validation failed")
		return false
	}
	return true
}
