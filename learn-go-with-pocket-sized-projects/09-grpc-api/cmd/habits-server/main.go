package main

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/internal/log"
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/internal/server"
	"os"
)

const port = 28710

func main() {
	// Set the writing output of our logger.
	lgr := log.New(os.Stdout)

	srv := server.New(lgr)

	err := srv.ListenAndServe(port)
	if err != nil {
		lgr.Logf("Error while running the server: %s", err.Error())
		os.Exit(1)
	}
}
