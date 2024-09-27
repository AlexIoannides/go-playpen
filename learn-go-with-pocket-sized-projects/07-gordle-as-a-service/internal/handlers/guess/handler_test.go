package guess

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/07-gordle-as-a-service/internal/api"
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/07-gordle-as-a-service/internal/gordle"
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/07-gordle-as-a-service/internal/session"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	game, _ := gordle.New("pocket")
	handle := Handler(successGameGuesserStub{session.Game{
		ID:           "123456",
		Gordle:       *game,
		AttemptsLeft: 5,
		Guesses:      nil,
		Status:       session.StatusPlaying,
	}})

	req, err := http.NewRequest(http.MethodPost, "/games/", strings.NewReader(`{"guess":"pocket"}`))
	require.NoError(t, err)

	// add path parameters
	req.SetPathValue(api.GameID, "123456")

	recorder := httptest.NewRecorder()

	handle(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"123456","attempts_left":4,"guesses":[{"word":"pocket", "feedback":"++++++"}],"word_length":6,"status":"Won"}`, recorder.Body.String())
}

type successGameGuesserStub struct {
	game session.Game
}

func (g successGameGuesserStub) Find(id session.GameID) (session.Game, error) {
	return g.game, nil
}

func (g successGameGuesserStub) Update(game session.Game) error {
	return nil
}
