package habit_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/internal/habit"
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/internal/habit/mocks"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	h := habit.Habit{
		Name:            "swim",
		WeeklyFrequency: 2,
		CreationTime:    time.Now(),
		ID:              "123",
	}

	dbErr := fmt.Errorf("db unavailable")

	tests := map[string]struct {
		db          func(ctl *minimock.Controller) *mocks.HabitCreatorMock
		expectedErr error
	}{
		"nominal": {
			db: func(ctl *minimock.Controller) *mocks.HabitCreatorMock {
				db := mocks.NewHabitCreatorMock(ctl)
				db.AddMock.Expect(minimock.AnyContext, h).Return(nil)
				return db
			},
			expectedErr: nil,
		},
		"error case": {
			db: func(ctl *minimock.Controller) *mocks.HabitCreatorMock {
				db := mocks.NewHabitCreatorMock(ctl)
				db.AddMock.Expect(minimock.AnyContext, h).Return(dbErr)
				return db
			},
			expectedErr: dbErr,
		},
		"db timeout": {
			db: func(ctl *minimock.Controller) *mocks.HabitCreatorMock {
				db := mocks.NewHabitCreatorMock(ctl)
				db.AddMock.Set(
					func(ctx context.Context, habit habit.Habit) error {
						select {
						// This duration is longer than a database call
						case <-time.After(2 * time.Second):
							return nil
						case <-ctx.Done():
							return ctx.Err()
						}
					})
				return db
			},
			expectedErr: context.DeadlineExceeded,
		},
	}

	for name, tt := range tests {
		name, tt := name, tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := minimock.NewController(t)

			db := tt.db(ctrl)

			got, err := habit.Create(context.Background(), db, h)
			assert.ErrorIs(t, err, tt.expectedErr)
			if tt.expectedErr == nil {
				assert.Equal(t, h.Name, got.Name)
			}
		})
	}
}
