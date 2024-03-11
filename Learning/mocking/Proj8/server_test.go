package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/mock"
)

// smsServiceMock
type smsServiceMock struct {
	mock.Mock
}

// Our mocked smsService method
func (m *smsServiceMock) SendChargeNotification(value int) bool {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)
	// this records that the method was called and passes in the value
	// it was called with
	args := m.Called(value)
	// it then returns whatever we tell it to return
	// in this case true to simulate an SMS Service Notification
	// sent out
	return args.Bool(0)
}

// We need to satisfy our MessageService interface
// which sadly means we have to stub out every method
// defined in that interface
func (m *smsServiceMock) DummyFunc() {
	fmt.Println("Dummy")
}

// TestChargeCustomer is where the magic happens
// here we create our smsService mock
func TestChargeCustomer(t *testing.T) {
	smsService := new(smsServiceMock)
	// We then define what should be returned from SendChargeNotification
	// when we pass in the value 100 to it. In this case, we want to return
	// true as it was successful in sending a notification
	smsService.On("SendChargeNotification", 100).Return(true)
	// Next, we want to define the service we wish to test
	myService := MyService{smsService}
	// And call said method
	myService.ChargeCustomer(100)
	// At the end, we verify that our myService.ChargeCustomer
	// method called our mocked SendChargeNotification method
	smsService.AssertExpectations(t)
}

// Define MyService struct and its methods
type MyService struct {
	SMS *smsServiceMock
}

func (m *MyService) ChargeCustomer(amount int) {
	// Here you would implement the actual logic to charge the customer
	// In this example, we'll just call the mock method
	m.SMS.SendChargeNotification(amount)
}
