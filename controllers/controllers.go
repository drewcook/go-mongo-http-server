package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/drewcook/go-mongo-http-server/repository"
)

// Get the collection that the controller maps to
const COLL_NAME = "books"

var collection = repository.GetCollection(COLL_NAME)

func SeedBooks(w http.ResponseWriter, r *http.Request) {
	// Set headers
	r.Header.Set("Access-Control-Allow-Methods", "POST")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Access-Control-Allow-Origin", "*")

	// Seed the collection
	repository.Seed(collection)

	json.NewEncoder(w).Encode(true)
}

// GetBooks returns all books from the collection
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Set headers
	r.Header.Set("Access-Control-Allow-Methods", "GET")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Access-Control-Allow-Origin", "*")

	// Encode data as JSON and write to response
	data := repository.GetBooks(collection)
	json.NewEncoder(w).Encode(data)
}
