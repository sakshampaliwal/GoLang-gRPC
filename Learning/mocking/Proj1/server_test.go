// server_test.go
package main

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockDataFetcher is a mock implementation of DataFetcher.
type MockDataFetcher struct {
    mock.Mock
}

// FetchData mocks the FetchData method.
func (m *MockDataFetcher) FetchData() string {
    args := m.Called()
    return args.String(0)
}

func TestHandleRequest(t *testing.T) {
    // Create a new instance of the mock DataFetcher
    mockFetcher := new(MockDataFetcher)

    // Set up expectations
    mockFetcher.On("FetchData").Return("Mocked data")

    // Create a new instance of the server with the mock data fetcher
    server := NewServer(mockFetcher)

    // Call the method we want to test
    result := server.HandleRequest()

    // Assert that the result is what we expect
    assert.Equal(t, "Data received: Mocked data", result)

    // Verify that the FetchData method was called exactly once
    mockFetcher.AssertExpectations(t)
}
