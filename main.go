package main

import (
	"log"
	"net/http"
	"pbl409-dashboard/config"
	"pbl409-dashboard/database/migration"
	"pbl409-dashboard/router"

	"github.com/rs/cors"
)

func main() {
	db := config.ConnectDB()
	router := router.Router(db)
	migration.MigrateAll()

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
