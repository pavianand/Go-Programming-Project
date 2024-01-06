package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-example/grpc_example" // Import generated protobuf package
)

type server struct {
	pb.UnimplementedItemServiceServer // Embed the gRPC interface
}

func (s *server) AddItem(ctx context.Context, req *pb.ItemRequest) (*pb.ItemResponse, error) {
	newItem := pb.ItemResponse{
		Id:   "ID", // You can generate a unique ID here
		Name: req.Name,
	}
	return &newItem, nil
}

func (s *server) GetItems(ctx context.Context, req *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	// Implement logic to retrieve items
	return &pb.GetItemsResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterItemServiceServer(s, &server{})

	log.Println("gRPC server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
