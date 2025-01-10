package main

import (
	"crud-api/config"
	"crud-api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	config.InitDB()

	// Set up the router
	r := router.SetupRouter()

	// Start the server
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
