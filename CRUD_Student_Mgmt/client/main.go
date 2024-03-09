package main

import (
	"context"
	pb "crud1_proj_copy/gencode"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Read the server address from the environment variable
	serverAddress := os.Getenv("PORT")
	if serverAddress == "" {
		log.Fatal("PORT environment variable is not set")
	}

	// Set up a connection to the server
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDServiceClient(conn)


	for {
		var option int
		fmt.Println(".................................")
		fmt.Println("Choose CRUD operation:")
		fmt.Println("1. Create")
		fmt.Println("2. Read")
		fmt.Println("3. Update")
		fmt.Println("4. Delete")
		fmt.Println("5. To Quit")
		fmt.Print("Enter option number: ")
		fmt.Scanln(&option)

		switch option {
		case 1: //Create
			var name string
			var branch string
			var roll int64
			fmt.Print("Enter RollNo of Student: ")
			fmt.Scanln(&roll)
			fmt.Print("Enter Name of Student:")
			fmt.Scanln(&name)
			fmt.Print("Enter Branch of Student:")
			fmt.Scanln(&branch)

			// Call Create
			createResponse, err := c.Create(context.Background(), &pb.CreateRequest{Rollno: roll, Name: name, Branch: branch})
			if err != nil {
				log.Fatalf("Create failed: %v", err)
			}
			log.Printf("Created with RollNo: %d", createResponse.Rollno)

		case 2: //Read
			var roll int64
			fmt.Print("Enter RollNo of Student: ")
			fmt.Scanln(&roll)

			// Call Read
			readResponse, err := c.Read(context.Background(), &pb.ReadRequest{Rollno: roll})
			if err != nil {
				log.Fatalf("Read failed: %v", err)
			}
			log.Printf("Name of Student: %s", readResponse.Name)
			log.Printf("Branch is: %s", readResponse.Branch)

		case 3: //Update
			var roll int64
			var choice int64
			fmt.Print("Enter RollNo of Student to Update: ")
			fmt.Scanln(&roll)
			fmt.Print("Enter 1 if you want to update Name and 2 to update Branch and 3 for Both: ")
			fmt.Scanln(&choice)

			if choice == 1 {
				var chname string
				fmt.Print("Enter Updated Name: ")
				fmt.Scanln(&chname)
				updateResponse, err := c.Update(context.Background(), &pb.UpdateRequest{Rollno: roll, Name: chname})
				if err != nil {
					log.Fatalf("Update failed: %v", err)
				}
				log.Printf("Update response: %v", updateResponse)
			}

			if choice == 2 {
				var chbranch string
				fmt.Print("Enter Updated Branch: ")
				fmt.Scanln(&chbranch)
				updateResponse, err := c.Update(context.Background(), &pb.UpdateRequest{Rollno: roll, Branch: chbranch})
				if err != nil {
					log.Fatalf("Update failed: %v", err)
				}
				log.Printf("Update response: %v", updateResponse)
			}

			if choice == 3 {
				var chname string
				fmt.Print("Enter Updated Name: ")
				fmt.Scanln(&chname)
				var chbranch string
				fmt.Print("Enter Updated Branch: ")
				fmt.Scanln(&chbranch)
				updateResponse, err := c.Update(context.Background(), &pb.UpdateRequest{Rollno: roll, Name: chname, Branch: chbranch})
				if err != nil {
					log.Fatalf("Update failed: %v", err)
				}
				log.Printf("Update response: %v", updateResponse)
			}

		case 4: //Delete
			var roll int64
			fmt.Print("Enter RollNo of Student to Delete: ")
			fmt.Scanln(&roll)

			// Call Delete
			deleteResponse, err := c.Delete(context.Background(), &pb.DeleteRequest{Rollno: roll})
			if err != nil {
				log.Fatalf("Delete failed: %v", err)
			}
			log.Printf("Delete response: %v", deleteResponse)
		case 5: //To Quit
			fmt.Println("Quitting...")
			return
		default:
			log.Fatalf("Invalid Option Selected")
		}
	}
}
