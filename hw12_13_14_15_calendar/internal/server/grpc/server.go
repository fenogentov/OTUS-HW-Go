package grpc_server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/fenogentov/OTUS-HW-Go/hw12_13_14_15_calendar/internal/server/grpc/proto"

	"google.golang.org/grpc"
)

// Server ...
type Server struct {
	lis    net.Listener
	server *grpc.Server
}

// NewServer ...
func NewServer(host, port string) *Server {
	addr := net.JoinHostPort(host, port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	serv := pb.UnimplementedCalendarServer{}
	pb.RegisterCalendarServer(server, serv)

	return &Server{lis: lis, server: server}
}

// Start ...
func (s *Server) Start(ctx context.Context) error {
	fmt.Println("start grpc")
	if err := s.server.Serve(s.lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}
