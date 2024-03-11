package main

import (
	"net"
	"testing"
)

func TestServerStarts(t *testing.T) {
	// Try to listen on the port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	lis.Close() // Close the listener
}
