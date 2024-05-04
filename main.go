package main

import (
	"go-web-native/config"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// 	Database connection
	config.ConnectDB()
	// set routes
	route := mux.NewRouter()

	addApproutes(route)
	// Run server
	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", route))
}
