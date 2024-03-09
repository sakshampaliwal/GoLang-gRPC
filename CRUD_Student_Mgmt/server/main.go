package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "crud1_proj_copy/gencode"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCRUDServiceServer
	data   map[int64][]string
	logger *zap.Logger
}

func (s *server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	roll := req.Rollno

	// Check if the user already exists with the given roll number
	if _, exists := s.data[roll]; exists {
		s.logger.Warn("User already exists", zap.Int64("rollno", roll))
		// Return a response indicating that the user already exists
		return &pb.CreateResponse{
			Rollno:    roll,
			Createmsg: fmt.Sprintf("User with roll number %d already exists", roll),
		}, nil
	}

	s.data[roll] = []string{req.Name, req.Branch}
	s.logger.Info("Created data", zap.Int64("rollno", roll))
	return &pb.CreateResponse{Rollno: roll, Createmsg: "Creation Done!"}, nil
}

func (s *server) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	data, ok := s.data[req.Rollno]
	if !ok {
		s.logger.Warn("Data not found", zap.Int64("rollno", req.Rollno))
		// Return a successful response with a message indicating that the roll number is not present
		return &pb.ReadResponse{
			Name:   fmt.Sprintf("Roll number %v not found", req.Rollno),
			Branch: "Roll number not found",
		}, nil
	}
	s.logger.Info("Read data", zap.Int64("rollno", req.Rollno))
	return &pb.ReadResponse{Name: data[0], Branch: data[1]}, nil
}

func (s *server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	data, ok := s.data[req.Rollno]
	if !ok {
		s.logger.Warn("Data not found", zap.Int64("rollno", req.Rollno))
		// Return a successful response with a message indicating that the roll number is not present
		// return nil, fmt.Errorf("roll number %v not found", req.Rollno)
		return &pb.UpdateResponse{Umsg: fmt.Sprintf("Roll number %v not found", req.Rollno)}, nil
	}

	if req.Name != "" {
		data[0] = req.Name
		s.data[req.Rollno] = []string{req.Name, data[1]}
	}
	if req.Branch != "" {
		data[1] = req.Branch
		s.data[req.Rollno] = []string{data[0], req.Branch}
	}
	s.logger.Info("Updated data", zap.Int64("rollno", req.Rollno))
	return &pb.UpdateResponse{Umsg: "Update Successfully!"}, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	_, ok := s.data[req.Rollno]
	if !ok {
		s.logger.Warn("Data not found", zap.Int64("rollno", req.Rollno))
		// Return a successful response with a message indicating that the roll number is not present
		return &pb.DeleteResponse{Dmsg: fmt.Sprintf("Data with RollNo %v not found", req.Rollno)}, nil
	}
	delete(s.data, req.Rollno)
	s.logger.Info("Deleted data", zap.Int64("rollno", req.Rollno))
	return &pb.DeleteResponse{Dmsg: "Student Deletion Done!"}, nil
}

func main() {
	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	defer logger.Sync()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}
	s := grpc.NewServer()
	pb.RegisterCRUDServiceServer(s, &server{data: make(map[int64][]string), logger: logger}) //creates an instance of the server struct, initializes the data field with new empty map
	logger.Info("Server started")
	if err := s.Serve(lis); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
