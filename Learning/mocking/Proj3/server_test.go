package main

import (
    "github.com/stretchr/testify/mock"
    "testing"
)

// Define a mock struct
type MockedStruct struct {
    mock.Mock
}

// Mocking the function
func (m *MockedStruct) MyFunction() string {
    args := m.Called()
    return args.String(0)
}

func TestMyFunction(t *testing.T) {
    // Create a new instance of the mock
    mockInstance := new(MockedStruct)

    // Define the behavior of the mock
    mockInstance.On("MyFunction").Return("Mocked Function")

    // Call the function under test
    result := mockInstance.MyFunction()

    // Assert the result
    if result != "Mocked Function" {
        t.Errorf("Expected 'Mocked Function', got '%s'", result)
    }

    // Assert that the expected method was called
    mockInstance.AssertCalled(t, "MyFunction")
}
