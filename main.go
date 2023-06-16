package main

import (
	"log"
	"net/http"

	"github.com/drewcook/go-mongo-http-server/router"
)

const host = "localhost"
const port = "5280"

func main() {
	// Create a new mux router
	r := router.Router()
	// Open up server and listen on host and port
	log.Fatal(http.ListenAndServe(host+":"+port, r))
}
