package main

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	defer func() { _ = listener.Close() }()

	t.Logf("bound to port %q", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()
		}(conn)
	}
}
