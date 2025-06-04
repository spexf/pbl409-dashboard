package utils

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	res := Result{Code: code, Data: payload, Message: message}
	json.NewEncoder(w).Encode(res)

}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, nil, message)
}
