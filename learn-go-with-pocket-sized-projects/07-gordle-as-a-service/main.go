package main

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/07-gordle-as-a-service/internal/handlers"
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/07-gordle-as-a-service/internal/repository"
	"log"
	"net/http"
)

func main() {
	// Create the games' repository.
	db := repository.New()

	addr := ":8080"

	log.Print("Listening on ", addr, "...")

	// Start the server.
	err := http.ListenAndServe(addr, handlers.NewRouter(db))
	if err != nil {
		panic(err)
	}
}
