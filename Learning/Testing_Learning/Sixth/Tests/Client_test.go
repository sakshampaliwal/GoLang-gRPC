package main

import (
    "context"
    "testing"
    "time"

    "google.golang.org/grpc"
)

func TestClientConnects(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Try to connect to the server
    conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithBlock(), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("failed to connect to server: %v", err)
    }
    conn.Close() // Close the connection
}
