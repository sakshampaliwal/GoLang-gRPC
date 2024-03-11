// package main

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // MockAddSubtract is a mock implementation of the Add and Subtract functions.
// type MockAddSubtract struct {
// 	mock.Mock
// }

// // Add mocks the Add function.
// func (m *MockAddSubtract) Add(a, b int) int {
// 	args := m.Called(a, b)
// 	return args.Int(0)
// }

// // Subtract mocks the Subtract function.
// func (m *MockAddSubtract) Subtract(a, b int) int {
// 	args := m.Called(a, b)
// 	return args.Int(0)
// }

// func TestCalculateResult(t *testing.T) {
// 	// Create a new instance of the mock
// 	mockAddSubtract := new(MockAddSubtract)

// 	// Set up the behavior of the mock
// 	mockAddSubtract.On("Add", 3, 5).Return(8)       // Mocking Add(3, 5) to return 8
// 	mockAddSubtract.On("Subtract", 8, 2).Return(6) // Mocking Subtract(8, 2) to return 6

// 	// Call the function under test
// 	result := CalculateResult(3, 5, 2)

// 	// Assert the result
// 	expected := 6 // Since we mocked Subtract(8, 2) to return 6
// 	assert.Equal(t, expected, result)

// 	// Assert that Add and Subtract were called with the expected arguments
// 	mockAddSubtract.AssertCalled(t, "Add", 3, 5)
// 	mockAddSubtract.AssertCalled(t, "Subtract", 8, 2)
// }
