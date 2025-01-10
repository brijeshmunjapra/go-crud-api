// router/router.go
package router

import (
	"crud-api/handlers"

	"github.com/gorilla/mux"
)

// SetupRouter initializes the routes and returns a router
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/items", handlers.GetItems).Methods("GET")
	r.HandleFunc("/items/{id}", handlers.GetItem).Methods("GET")
	r.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")

	return r
}
