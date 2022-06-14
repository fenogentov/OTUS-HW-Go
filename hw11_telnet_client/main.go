package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var timeout time.Duration

func init() {
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "timeout connection (default 10s)")
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 2 {
		log.Fatal("host and port must be provided as two arguments\n")
	}
	address := net.JoinHostPort(flag.Arg(0), flag.Arg(1))

	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		fmt.Println(fmt.Errorf("error connect: %w", err))
		return
	}
	defer func() {
		if err := client.Close(); err != nil {
			fmt.Println(fmt.Errorf("error close connect: %w", err))
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	go signalNotify(cancel)

	go send(cancel, client)

	go receive(cancel, client)

	<-ctx.Done()
}

func send(cancel context.CancelFunc, client TelnetClient) {
	client.Send()
	cancel()
}

func receive(cancel context.CancelFunc, client TelnetClient) {
	client.Receive()
	cancel()
}

func signalNotify(cancel context.CancelFunc) {
	chC := make(chan os.Signal, 1)
	signal.Notify(chC, syscall.SIGINT)
	<-chC
	cancel()
}
