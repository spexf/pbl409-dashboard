package router

import (
	service "pbl409-dashboard/pkg/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *mux.Router {

	serviceHandler := &service.ServiceHandler{
		DB: db,
	}
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1/").Subrouter()
	api.HandleFunc("/services", serviceHandler.GetService).Methods("GET")
	api.HandleFunc("/services", serviceHandler.StoreService).Methods("POST")
	api.HandleFunc("/services/{id}", serviceHandler.ShowService).Methods("GET")
	api.HandleFunc("/services/{id}", serviceHandler.DeleteService).Methods("DELETE", "OPTIONS")
	return r
}
