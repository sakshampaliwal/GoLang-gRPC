package main

import (
	"context"
	"fmt"
	"testing"

	pb "crud1_proj_copy/gencode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the CRUDServiceServer interface
type mockServer struct {
	mock.Mock
}

func (m *mockServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	args := m.Called(ctx, req) //it is used to call create method and returned answer will be stored in the args.
	return args.Get(0).(*pb.CreateResponse), args.Error(1)
}

func (m *mockServer) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.ReadResponse), args.Error(1)
	//get is used to retrieve specific values from the return values of the mocked function.
	//error is used to handle errors returned by the mocked function.
}

func (m *mockServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.UpdateResponse), args.Error(1)
}

func (m *mockServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.DeleteResponse), args.Error(1)
}

func TestCRUDOperations(t *testing.T) {
	mockServer := new(mockServer)

	createReq := &pb.CreateRequest{
		Rollno: 123,
		Name:   "John",
		Branch: "CSE",
	}

	mockServer.On("Create", mock.Anything, createReq).Return(&pb.CreateResponse{Rollno: 123, Createmsg: "Creation Done!"}, nil)

	createResp, createErr := mockServer.Create(context.Background(), createReq)
	assert.NoError(t, createErr)
	assert.NotNil(t, createResp)
	assert.Equal(t, int64(123), createResp.Rollno)
	assert.Equal(t, "Creation Done!", createResp.Createmsg)

	createReq2 := &pb.CreateRequest{
		Rollno: 123,
		Name:   "Alice", // Using a different name
		Branch: "ECE",   // Using a different branch
	}

	mockServer.On("Create", mock.Anything, createReq2).Return(&pb.CreateResponse{
		Rollno:    123,
		Createmsg: "User with roll number 123 already exists",
	}, nil)

	createResp2, createErr2 := mockServer.Create(context.Background(), createReq2)
	assert.NoError(t, createErr2)
	assert.NotNil(t, createResp2)
	assert.Equal(t, int64(123), createResp2.Rollno)
	assert.Equal(t, "User with roll number 123 already exists", createResp2.Createmsg)

	// Test Read operation
	readReq := &pb.ReadRequest{
		Rollno: 123,
	}
	mockServer.On("Read", mock.Anything, readReq).Return(&pb.ReadResponse{Name: "John", Branch: "CSE"}, nil)
	readResp, readErr := mockServer.Read(context.Background(), readReq)
	assert.NoError(t, readErr)
	assert.NotNil(t, readResp)
	assert.Equal(t, "John", readResp.Name)
	assert.Equal(t, "CSE", readResp.Branch)

	//Testing Reading of roll no that is not present in the map

	readReq1 := &pb.ReadRequest{
		Rollno: 555, // Roll number that doesn't exist
	}

	mockServer.On("Read", mock.Anything, readReq1).Return(&pb.ReadResponse{
		Name:   fmt.Sprintf("Roll number %v not found", readReq1.Rollno),
		Branch: "Roll number not found",
	}, nil)
	readResp1, readErr1 := mockServer.Read(context.Background(), readReq1)
	assert.NoError(t, readErr1)
	assert.NotNil(t, readResp1)
	assert.Equal(t, fmt.Sprintf("Roll number %v not found", readReq1.Rollno), readResp1.Name)
	assert.Equal(t, "Roll number not found", readResp1.Branch)

	// Test Update operation
	updateReq := &pb.UpdateRequest{
		Rollno: 123,
		Name:   "John Cena",
		Branch: "ECE",
	}
	mockServer.On("Update", mock.Anything, updateReq).Return(&pb.UpdateResponse{Umsg: "Update Successfully!"}, nil)
	updateResp, updateErr := mockServer.Update(context.Background(), updateReq)
	assert.NoError(t, updateErr)
	assert.NotNil(t, updateResp)
	assert.Equal(t, "Update Successfully!", updateResp.Umsg)

	//Testing Updation if Roll number is not present in the map
	updateReq1 := &pb.UpdateRequest{
		Rollno: 555, // Roll number that doesn't exist
	}

	mockServer.On("Update", mock.Anything, updateReq1).Return(&pb.UpdateResponse{
		Umsg: fmt.Sprintf("Roll number %v not found", updateReq1.Rollno),
		// Name:   fmt.Sprintf("Roll number %v not found", updateReq1.Rollno),
		// Branch: "Roll number not found",
	}, nil)
	updateResp1, updateErr1 := mockServer.Update(context.Background(), updateReq1)
	assert.NoError(t, updateErr1)
	assert.NotNil(t, updateResp1)
	assert.Equal(t, fmt.Sprintf("Roll number %v not found", updateReq1.Rollno), updateResp1.Umsg)
	// assert.Equal(t, "Roll number not found", updateResp1.Umsg)

	// Test Delete operation
	deleteReq := &pb.DeleteRequest{
		Rollno: 123,
	}
	mockServer.On("Delete", mock.Anything, deleteReq).Return(&pb.DeleteResponse{Dmsg: "Student Deletion Done!"}, nil)
	deleteResp, deleteErr := mockServer.Delete(context.Background(), deleteReq)
	assert.NoError(t, deleteErr)
	assert.NotNil(t, deleteResp)
	assert.Equal(t, "Student Deletion Done!", deleteResp.Dmsg)

	// Testing Deletion when rollno is not present

	deleteReq1 := &pb.DeleteRequest{
		Rollno: 555, // Roll number that doesn't exist
	}

	mockServer.On("Delete", mock.Anything, deleteReq1).Return(&pb.DeleteResponse{Dmsg: fmt.Sprintf("Data with RollNo %v not found", deleteReq1.Rollno)}, nil)
	deleteResp1, deleteErr1 := mockServer.Delete(context.Background(), deleteReq1)
	assert.NoError(t, deleteErr1)
	assert.NotNil(t, deleteResp1)
	assert.Equal(t, fmt.Sprintf("Data with RollNo %v not found", deleteReq1.Rollno), deleteResp1.Dmsg)
	// assert.Equal(t, "Roll number not found", deleteResp1.Dmsg)

}
