package server

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/api"
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/internal/habit"
	r "alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/internal/repository"
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetHabitStatus is the endpoint that retrieves the status of a habit per week.
func (s *Server) GetHabitStatus(ctx context.Context, request *api.GetHabitStatusRequest) (*api.GetHabitStatusResponse, error) {
	s.lgr.Logf("GetHabitStatus request received: %s", request)

	err := validateGetHabitRequest(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request: "+err.Error())
	}

	// if empty, the timestamp is set to the current time
	var t time.Time
	if request.Timestamp == nil {
		t = time.Now()
	}

	h, ticksCount, err := habit.GetStatus(ctx, s.db, s.db, habit.ID(request.HabitId), t)
	if err != nil {
		switch {
		case errors.Is(err, r.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "couldn't find habit %q in repository", request.HabitId)
		default:
			return nil, status.Errorf(codes.Internal, "cannot get status %q: %s", h.ID, err.Error())
		}
	}

	return &api.GetHabitStatusResponse{
		Habit: &api.Habit{
			Id:              string(h.ID),
			Name:            string(h.Name),
			WeeklyFrequency: int32(h.WeeklyFrequency),
		},
		TicksCount: int32(ticksCount),
	}, nil
}

func validateGetHabitRequest(request *api.GetHabitStatusRequest) error {
	switch {
	case request == nil:
		return fmt.Errorf("empty request")
	case request.HabitId == "":
		return fmt.Errorf("missing habit id")
	}
	return nil
}
