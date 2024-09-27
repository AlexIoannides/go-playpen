package server

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/internal/log"
)

func Test_timerInterceptor(t *testing.T) {
	info := &grpc.UnaryServerInfo{FullMethod: "TestingFunc"}
	handler := func(ctx context.Context, req any) (any, error) {
		return "123", nil
	}

	bfr := bytes.NewBuffer([]byte{})
	lgr := log.New(bfr)

	// Use the t variable for logging
	interceptor := timerInterceptor(lgr)

	_, _ = interceptor(context.Background(), "request", info, handler)
	assert.Contains(t, bfr.String(), "time in TestingFunc: ")
}
