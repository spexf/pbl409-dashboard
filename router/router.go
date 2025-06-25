package router

import (
	"pbl409-dashboard/pkg/agents"
	"pbl409-dashboard/pkg/auth"
	middleware "pbl409-dashboard/pkg/middleware"
	service "pbl409-dashboard/pkg/services"
	user "pbl409-dashboard/pkg/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *mux.Router {
	serviceHandler := &service.ServiceHandler{DB: db}
	userHandler := &user.UserHandler{DB: db}
	authHandler := &auth.AuthHandler{DB: db}
	agentHandler := &agents.AgentHandler{DB: db}

	r := mux.NewRouter()

	// Public routes
	public := r.PathPrefix("/api/v1").Subrouter()
	public.HandleFunc("/auth/login", authHandler.LoginHandler).Methods("POST")

	// Protected routes
	private := r.PathPrefix("/api/v1").Subrouter()
	private.Use(middleware.JWTAuth)

	// Services
	private.HandleFunc("/services", serviceHandler.GetService).Methods("GET")
	private.HandleFunc("/services", serviceHandler.StoreService).Methods("POST")
	private.HandleFunc("/services/{id}", serviceHandler.ShowService).Methods("GET")
	private.HandleFunc("/services/{id}", serviceHandler.UpdateService).Methods("PUT", "OPTIONS")
	private.HandleFunc("/services/{id}", serviceHandler.DeleteService).Methods("DELETE", "OPTIONS")

	// Users
	private.HandleFunc("/users", userHandler.GetUser).Methods("GET")
	private.HandleFunc("/users", userHandler.StoreUser).Methods("POST")
	private.HandleFunc("/users/{id}", userHandler.ShowUser).Methods("GET")
	private.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT", "OPTIONS")
	private.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE", "OPTIONS")

	// Agents
	private.HandleFunc("/wazuh/{id}/agents", agentHandler.GetAgents).Methods("GET")
	private.HandleFunc("/wazuh/{id}/agents/{agentName}", agentHandler.GetAgentData).Methods("GET")
	private.HandleFunc("/wazuh/{id}/agents", agentHandler.CreateAgents).Methods("POST")
	private.HandleFunc("/wazuh/{id}/agents", agentHandler.DeleteAgents).Methods("DELETE", "OPTIONS")
	return r
}
