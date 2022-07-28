package greet

import (
	context "context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	grpc "google.golang.org/grpc"
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

// unary
func ( *server) Greet(ctx context.Context, req *GreetReq) (*GreetRes, error) {
	fmt.Printf("Greet function invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Fuck you " + firstName
	res := &GreetRes{
		Result: result,
	}
	return res, nil
}

// server streaming
func ( *server) GreetManyTimes(req *GreetManyTimesReq, stream GreetService_GreetManyTimesServer) error {
	fmt.Printf("Greet function invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Fuck you %s %s - %s times", firstName, lastName, strconv.Itoa(i))
		res := &GreetManyTimesRes{
			Result: result,
		}
		stream.SendMsg(res)
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}
