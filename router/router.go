package router

import (
	"pbl409-dashboard/pkg/auth"
	service "pbl409-dashboard/pkg/services"
	user "pbl409-dashboard/pkg/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *mux.Router {

	serviceHandler := &service.ServiceHandler{
		DB: db,
	}
	userHandler := &user.UserHandler{
		DB: db,
	}
	authHandler := &auth.AuthHandler{
		DB: db,
	}

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1/").Subrouter()
	// Service
	api.HandleFunc("/services", serviceHandler.GetService).Methods("GET")
	api.HandleFunc("/services", serviceHandler.StoreService).Methods("POST")
	api.HandleFunc("/services/{id}", serviceHandler.ShowService).Methods("GET")
	api.HandleFunc("/services/{id}", serviceHandler.UpdateService).Methods("PUT", "OPTIONS")
	api.HandleFunc("/services/{id}", serviceHandler.DeleteService).Methods("DELETE", "OPTIONS")
	// Authenticate User
	api.HandleFunc("/auth/login", authHandler.LoginHandler).Methods("POST")
	// User
	api.HandleFunc("/users", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users", userHandler.StoreUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.ShowUser).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT", "OPTIONS")
	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE", "OPTIONS")
	return r
}
