package handlers

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/07-gordle-as-a-service/internal/api"
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/07-gordle-as-a-service/internal/handlers/getstatus"
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/07-gordle-as-a-service/internal/handlers/guess"
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/07-gordle-as-a-service/internal/handlers/newgame"
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/07-gordle-as-a-service/internal/repository"
	"net/http"
)

// NewRouter returns a router that listens for requests to the following endpoints:
//   - Create a new game;
//   - Get the status of a game;
//   - Make a guess in a game.
//
// The provided router is ready to serve.
func NewRouter(db *repository.GameRepository) *http.ServeMux {
	r := http.NewServeMux()

	// Register each endpoint.
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handler(db))    // curl -X POST -v http://localhost:8080/games
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handler(db)) // curl -X GET -v http://localhost:8080/games/1682279480
	r.HandleFunc(http.MethodPut+" "+api.GuessRoute, guess.Handler(db))         // curl -X PUT -v http://localhost:8080/games/1682279480 -d '{"guess":"faune"}'

	return r
}
