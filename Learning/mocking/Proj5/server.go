package main

import (
	"fmt"
)

// UserService is an interface for interacting with the user service.
type UserService interface {
	GetUserInfo(userID int) (string, error)
}

// UserServiceImpl is the implementation of UserService.
type UserServiceImpl struct{}

// GetUserInfo retrieves user information from the database.
func (us *UserServiceImpl) GetUserInfo(userID int) (string, error) {
	// Simulating fetching user information from a database
	// In a real scenario, this would interact with a database
	if userID == 1 {
		return "John Doe", nil
	} else {
		return "", fmt.Errorf("User not found")
	}
}

// WelcomeMessage generates a welcome message for a user.
func WelcomeMessage(userID int, userService UserService) (string, error) {
	userInfo, err := userService.GetUserInfo(userID)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Welcome, %s!", userInfo), nil
}
