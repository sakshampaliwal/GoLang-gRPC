package main

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of UserService.
type MockUserService struct {
    mock.Mock
}

// GetUserInfo mocks the GetUserInfo method of UserService.
func (m *MockUserService) GetUserInfo(userID int) (string, error) {
    args := m.Called(userID)
    return args.String(0), args.Error(1)
}

func TestWelcomeMessage(t *testing.T) {
    // Create a new instance of the mock UserService
    userService := new(MockUserService)

    // Set up the behavior of the mock
    userService.On("GetUserInfo", 1).Return("John Doe", nil)

    // Call the function under test
    message, err := WelcomeMessage(1, userService)

    // Assert the result
    assert.NoError(t, err)
    assert.Equal(t, "Welcome, John Doe!", message)

    // Assert that GetUserInfo was called with the expected argument
    userService.AssertCalled(t, "GetUserInfo", 1)
}
