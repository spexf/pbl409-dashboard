package router

import (
	"pbl409-dashboard/handler"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *mux.Router {

	serviceHandler := &handler.ServiceHandler{
		DB: db,
	}
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1/").Subrouter()
	api.HandleFunc("/services/{id}", serviceHandler.ShowService).Methods("GET")
	return r
}
