package router

import (
	"github.com/drewcook/go-mongo-http-server/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// Create a new mux router
	r := mux.NewRouter()

	// Set up routes an their handlers
	r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/seed", controllers.SeedBooks).Methods("POST")

	return r
}
