package main

import (
    "github.com/stretchr/testify/mock"
    "testing"
)

type Database interface {
    GetUserName(userID int) string
}

type UserService struct {
    DB Database
}

func (s *UserService) GetUserName(userID int) string {
    return s.DB.GetUserName(userID)
}	

---


type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) GetUserName(userID int) string {
	args := m.Called(userID) //When the method associated with this mock object is called with the provided userID argument, record that call and store the arguments passed to it in the variable args
	return args.String(0) //Extract the first argument passed to the method call
}

func TestGetUserName(t *testing.T) {
	// Create a mock instance
	mockDB := new(MockDatabase)  //creating a new instance of the MockDatabase struct, 

	// Set up expectation for GetUserName method
	mockDB.On("GetUserName", 1).Return("John") //you're setting up an expectation for the GetUserName method of the MockDatabase object (mockDB).
	// 1 argument specifies the expected value for the userID parameter when this method is called.
	// Create a UserService instance with the mockDB
	// it's saying that when the GetUserName method of mockDB is called with a userID of 1, it should return "John".

	userService := &UserService{DB: mockDB} //setting up the userService instance to use the mocked database (mockDB) for its database-related operations during testing.

	// Call the method we want to test
	result := userService.GetUserName(1)

	// Check if the result is as expected
	if result != "John" {
		t.Errorf("Expected 'John', got '%s'", result)
	}

	// Assert that the expectations are met
	mockDB.AssertExpectations(t) //This is a method provided by the mock.Mock object. It checks whether all the expectations set on the mock object have been satisfied.
	// If any expectations were not met, the test will fail and an error will be reported through the *testing.T object.
}

func TestGetUserNameOnce(t *testing.T) {
	// Create a mock instance
	mockDB := new(MockDatabase)

	// Set up expectation for GetUserName method to be called once
	mockDB.On("GetUserName", 1).Return("John").Once()  //once specifies that the method should be called exactly once. If it's called more than once, it will fail the test.

	// Create a UserService instance with the mockDB
	userService := &UserService{DB: mockDB}

	// Call the method we want to test
	userService.GetUserName(1)

	// Assert that the expectations are met
	mockDB.AssertExpectations(t)
}

func TestGetUserNameAny(t *testing.T) {
	// Create a mock instance
	mockDB := new(MockDatabase)

	// Set up expectation for GetUserName method with any userID
	mockDB.On("GetUserName", mock.Anything).Return("John") //This specifies that you're setting an expectation for the GetUserName method of the mock object, and you're allowing any argument to be passed to this method. mock.Anything is a placeholder that matches any value passed as an argument.
	//Regardless of the argument passed to the GetUserName method, it should always return "John".

	// Create a UserService instance with the mockDB
	userService := &UserService{DB: mockDB}

	// Call the method we want to test with any userID
	result := userService.GetUserName(100)

	// Check if the result is as expected
	if result != "John" {
		t.Errorf("Expected 'John', got '%s'", result)
	}

	// Assert that the expectations are met
	mockDB.AssertExpectations(t)
}
