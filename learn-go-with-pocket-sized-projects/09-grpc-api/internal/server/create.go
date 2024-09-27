package server

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/api"
	"context"
)

// CreateHabit is the endpoint that registers a habit.
func (s *Server) CreateHabit(_ context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	s.lgr.Logf("CreateHabit request received: %s", request)

	return &api.CreateHabitResponse{
		Habit: &api.Habit{},
	}, nil
}
