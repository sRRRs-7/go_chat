package gRPC

import (
	context "context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {}

func GreetServer() {
	fmt.Println("Start gRPC greet server")
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func ( *server) Greet(ctx context.Context, req *GreetReq) (*GreetRes, error) {
	fmt.Printf("Greet function invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Fuck you " + firstName
	res := &GreetRes{
		Result: result,
	}
	return res, nil
}
