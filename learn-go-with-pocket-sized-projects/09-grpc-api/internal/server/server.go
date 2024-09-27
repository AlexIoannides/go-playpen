package server

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/api"
	"fmt"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

// Logger used by the Server
type Logger interface {
	Logf(format string, args ...any)
}

// Server is the implementation of the gRPC server.
type Server struct {
	api.UnimplementedHabitsServer
	lgr Logger
}

// New returns a Server that can ListenAndServe.
func New(lgr Logger) *Server {
	return &Server{
		lgr: lgr,
	}
}

func (s *Server) ListenAndServe(port int) error {
	const addr = "127.0.0.1"

	listner, err := net.Listen("tcp", net.JoinHostPort(addr, strconv.Itoa(port)))
	if err != nil {
		return fmt.Errorf("unable to listen on TCP port %d: %w", port, err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s)

	s.lgr.Logf("starting server on port: %d", port)

	err = grpcServer.Serve(listner)
	if err != nil {
		return fmt.Errorf("error while listening: %w", err)
	}

	// Stop of GracefulStop was called, no reason to be alarmed
	return nil
}
