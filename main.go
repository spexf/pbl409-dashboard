package main

import (
	"log"
	"net"
	"pbl409-dashboard/database/migration"
)

func main() {
	migration.MigrateAll()
	log.Println("Server running at http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080"))
	net.Listen("localhost", "8080")
}
