package main

import (
    "testing"

    "github.com/stretchr/testify/mock"
)

// Mock for the Add function
type MockAdd struct {
    mock.Mock
}

func (m *MockAdd) Add(a, b int) int {
    args := m.Called(a, b)
    return args.Int(0)
}

func TestCalculate(t *testing.T) {
    // Create a new instance of the mock
    mockAdd := new(MockAdd)

    // Set up the behavior of the mock
    mockAdd.On("Add", 3, 5).Return(8) // Mocking Add(3, 5) to return 8

    // Call the function under test
    result := Calculate(3, 5)

    // Assert the result
    expected := 8 // Since we mocked Add(3, 5) to return 8
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }

    // Assert that Add was called with the expected arguments
    mockAdd.AssertCalled(t, "Add", 3, 5)
}
