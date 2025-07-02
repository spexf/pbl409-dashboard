package router

import (
	"html/template"
	"net/http"
	"path/filepath"

	"pbl409-dashboard/pkg/agents"
	"pbl409-dashboard/pkg/auth"
	"pbl409-dashboard/pkg/middleware"
	service "pbl409-dashboard/pkg/services"
	user "pbl409-dashboard/pkg/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Render HTML dengan base.html
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("templates", tmpl)
	tmpls, err := template.ParseFiles("templates/base.html", tmplPath)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpls.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}

// Middleware untuk halaman HTML, cek cookie token
func RequireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func Router(db *gorm.DB) *mux.Router {
	serviceHandler := &service.ServiceHandler{DB: db}
	userHandler := &user.UserHandler{DB: db}
	authHandler := &auth.AuthHandler{DB: db}
	agentHandler := &agents.AgentHandler{DB: db}

	r := mux.NewRouter()

	// Public page: login
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "login.html", map[string]interface{}{"Title": "Login"})
	})

	// Logout handler
	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "token",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// Protected pages
	r.HandleFunc("/dashboard", RequireLogin(func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "dashboard.html", map[string]interface{}{"Title": "Dashboard"})
	}))

	r.HandleFunc("/performa", RequireLogin(func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "performa.html", map[string]interface{}{"Title": "Performa"})
	}))

	r.HandleFunc("/agents", RequireLogin(func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "agent/list_agent.html", map[string]interface{}{"Title": "Daftar Agent"})
	}))
	r.HandleFunc("/agent/add", RequireLogin(func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "agent/add_agent.html", map[string]interface{}{"Title": "Tambah Agent"})
	}))
	r.HandleFunc("/agent/edit", RequireLogin(func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "agent/edit_agent.html", map[string]interface{}{"Title": "Edit Agent"})
	}))
	r.HandleFunc("/agent/delete", RequireLogin(func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "agent/delete_agent.html", map[string]interface{}{"Title": "Hapus Agent"})
	}))

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))

	// Public API
	public := r.PathPrefix("/api/v1").Subrouter()
	public.HandleFunc("/auth/login", authHandler.LoginHandler).Methods("POST")

	// Private API (JWT middleware)
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
