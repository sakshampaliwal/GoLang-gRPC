// server.go
package main

import (
    "fmt"
)

// DataFetcher is an interface for fetching data.
type DataFetcher interface {
    FetchData() string
}

// Server is a struct representing our server.
type Server struct {
    dataFetcher DataFetcher
}

// NewServer creates a new Server instance with the given dataFetcher.
func NewServer(dataFetcher DataFetcher) *Server {
    return &Server{dataFetcher}
}

// HandleRequest handles incoming requests by fetching data and returning it.
func (s *Server) HandleRequest() string {
    data := s.dataFetcher.FetchData()
    return fmt.Sprintf("Data received: %s", data)
}
