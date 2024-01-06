package main

import (
	"context"
	"log"

	"grpc-example/grpc_example" // Import the correct path to your generated code
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpc_example.NewItemServiceClient(conn)

	// Add Item
	addItemResponse, err := client.AddItem(context.Background(), &grpc_example.ItemRequest{Name: "Example"})
	if err != nil {
		log.Fatalf("Failed to add item: %v", err)
	}
	log.Printf("Added Item: %v", addItemResponse)

	// Get Items
	getItemsResponse, err := client.GetItems(context.Background(), &grpc_example.GetItemsRequest{})
	if err != nil {
		log.Fatalf("Failed to get items: %v", err)
	}
	log.Printf("Items: %v", getItemsResponse.Items)
}
