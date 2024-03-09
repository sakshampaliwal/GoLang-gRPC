package main_test

import (
	"context"
	"testing"

	pb "crud1_proj_copy/gencode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// This function will test the CRUD operations of the gRPC server.
func TestCRUDOperations(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance.
	client := pb.NewCRUDServiceClient(conn)

	// Test Create operation.
	createResponse, err := client.Create(context.Background(), &pb.CreateRequest{
		Rollno: 123,
		Name:   "John Doe",
		Branch: "Computer Science",
	})
	if err != nil {
		t.Fatalf("Create RPC failed: %v", err)
	}
	if createResponse.Createmsg != "Creation Done!" {
		t.Errorf("Expected creation message: 'Creation Done!', got: %s", createResponse.Createmsg)
	}
	if createResponse.Createmsg == "Creation Done!" {
		t.Log("Create Test Done!")
	}

	// Test Read operation.
	readResponse, err := client.Read(context.Background(), &pb.ReadRequest{
		Rollno: 123,
	})
	if err != nil {
		t.Fatalf("Read RPC failed: %v", err)
	}
	if readResponse.Name != "John Doe" || readResponse.Branch != "Computer Science" {
		t.Errorf("Read response mismatch. Expected: Name='John Doe', Branch='Computer Science', Got: Name='%s', Branch='%s'", readResponse.Name, readResponse.Branch)
	} else {
		t.Log("Read Test Done!")
	}

	// Test Read operation for a non-existing entry.
	nonExistingReadResponse, err := client.Read(context.Background(), &pb.ReadRequest{
		Rollno: 456,
	})
	if err != nil {
		t.Fatalf("Read RPC failed: %v", err)
	}
	if nonExistingReadResponse.Name != "Roll number 456 not found" || nonExistingReadResponse.Branch != "Roll number not found" {
		t.Errorf("Expected response for non-existing entry: Name='Roll number 456 not found', Branch='Roll number not found', Got: Name='%s', Branch='%s'", nonExistingReadResponse.Name, nonExistingReadResponse.Branch)
	} else {
		t.Log("Read for non-existing Test Done!")
	}

	// Test Update operation.
	updateResponse, err := client.Update(context.Background(), &pb.UpdateRequest{
		Rollno: 123,
		Name:   "Updated Name",
		Branch: "Updated Branch",
	})
	if err != nil {
		t.Fatalf("Update RPC failed: %v", err)
	}
	if updateResponse.Umsg != "Update Successfully!" {
		t.Errorf("Expected update message: 'Update Successfully!', got: %s", updateResponse.Umsg)
	} else {
		t.Log("Update Test Done!")
	}

	// Test Update operation for a non-existing entry.
	_, updateErr := client.Update(context.Background(), &pb.UpdateRequest{
		Rollno: 456,
		Name:   "Updated Name",
		Branch: "Updated Branch",
	})
	if updateErr != nil {
		t.Error("Expected an error for updating a non-existing entry, but got nil")
	} else {
		t.Log("Update for non-existing Test Done!")
	}

	// Test Delete operation.
	deleteResponse, err := client.Delete(context.Background(), &pb.DeleteRequest{
		Rollno: 123,
	})
	if err != nil {
		t.Fatalf("Delete RPC failed: %v", err)
	}
	if deleteResponse.Dmsg != "Student Deletion Done!" {
		t.Errorf("Expected deletion message: 'Student Deletion Done!', got: %s", deleteResponse.Dmsg)
	} else {
		t.Log("Delete Test Done!")
	}

	// Test Delete operation for a non-existing entry.
	nonExistingDeleteResponse, err := client.Delete(context.Background(), &pb.DeleteRequest{
		Rollno: 456,
	})
	if err != nil {
		t.Fatalf("Delete RPC failed: %v", err)
	}
	if nonExistingDeleteResponse.Dmsg != "Data with RollNo 456 not found" {
		t.Errorf("Expected response for deleting non-existing entry: 'Data with RollNo 456 not found', got: %s", nonExistingDeleteResponse.Dmsg)
	} else {
		t.Log("Delete for non-existing Test Done!")
	}
}
