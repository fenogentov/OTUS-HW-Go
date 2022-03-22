package grpc_server

import (
	"context"
	"fmt"
	"log"
	"net"

	"hw12_13_14_15_calendar/internal/logger"
	pb "hw12_13_14_15_calendar/internal/server/grpc/proto"

	"google.golang.org/grpc"
)

// Server ...
type Server struct {
	address string
	lis     net.Listener
	server  *grpc.Server
	logger  *logger.Logger
}

// NewServer ...
func NewServer(logger *logger.Logger, host, port string) *Server {
	addr := net.JoinHostPort(host, port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalendarServer(grpcServer, pb.UnimplementedCalendarServer{})

	return &Server{
		address: addr,
		lis:     lis,
		server:  grpcServer}
}

// Start ...
func (s *Server) Start(ctx context.Context) {
	fmt.Println("start grpc")
	err := s.server.Serve(s.lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) Stop() {

	if s.server != nil {
		s.server.Stop()
	}
	s.logger.Info("grpc server stop")
}
