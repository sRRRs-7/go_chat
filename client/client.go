package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sRRRs-7/go_chat/gRPC"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := gRPC.NewGreetServiceClient(conn)
	fmt.Printf("Create client: %f", client)
	fmt.Println()

	doUnary(client)
}

func doUnary(c gRPC.GreetServiceClient) {
	fmt.Println("Starting to do a unary RPC")
	req := &gRPC.GreetReq{
		Greeting: &gRPC.Greeting{
			FirstName: "michel",
			LastName: "jackson",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to calling RPC")
	}
	log.Printf("Response: %v", res)
}