package grpc_server

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	pb "github.com/fenogentov/OTUS-HW-Go/hw12_13_14_15_calendar/internal/server/grpc/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func TestServerGRPC(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go func() {
		srv := NewServer("localhost", "50051")
		srv.Start(ctx)
	}()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalendarClient(conn)

	event := &pb.Event{}

	evn, err := c.CreateEvent(ctx, event)
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	fmt.Println(evn.ID)
	// log.Printf("Product ID: %s added successfully", evn)

	// ev, err := c.UpdateEvent(ctx, event)
	// if err != nil {
	// 	log.Fatalf("Could not get product: %v", err)
	// }
	// log.Printf("Product: %v", ev)

}
