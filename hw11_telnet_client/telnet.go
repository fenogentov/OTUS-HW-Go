package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

type telnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
		conn:    nil,
	}
}

func (t *telnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return fmt.Errorf("cannot connect %s : %w", t.address, err)
	}
	t.conn = conn
	fmt.Fprintf(os.Stderr, "...Connected to %s\n", t.address)
	return nil
}

func (t *telnetClient) Send() error {
	if _, err := io.Copy(t.conn, t.in); err != nil {
		return fmt.Errorf("cannot send: %w", err)
	}
	fmt.Fprintf(os.Stderr, "...EOF\n")
	return nil
}

func (t *telnetClient) Receive() error {
	if _, err := io.Copy(t.out, t.conn); err != nil {
		return fmt.Errorf("cannot receive: %w", err)
	}
	fmt.Fprintf(os.Stderr, "...Connection was closed by peer\n")
	return nil
}

func (t *telnetClient) Close() error {
	return t.conn.Close()
}
