package main

import (
	"context"
	"net"
	"testing"
)

func TestCancelFanout(t *testing.T) {
	listener, err := net.Listen("tcp", "10.0.0.1:")
	if err != nil {
		t.Error(err)
	}
	defer func() { _ = listener.Close() }()

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Error(err)
		}

		conn.Close()
	}()

	ctx, cancel := context.WithCancel(context.Background())
	dial = func() {
		var d net.Dialer

		d.DialContext()
	}
	// generate multiple dials
	// cancel on first successfull connection
}
