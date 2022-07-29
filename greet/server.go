package greet

import (
	context "context"
	"fmt"
	"io"
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

// client streaming
func ( *server) LongGreet(stream GreetService_LongGreetServer) error {
	fmt.Println("LongGreet function invoked with stream request")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&LongGreetRes{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Failed to reading stream request: %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		lastName := req.GetGreeting().GetLastName()
		str := fmt.Sprintf("Fuck you %s %s !!", firstName, lastName)
		result += str
	}
}

// bidirectinal streaming
func ( *server) GreetEveryone(stream GreetService_GreetEveryoneServer) error {
	fmt.Println("starting to do bidirectional streaming")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Failed to receive requests: %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		lastName := req.GetGreeting().GetLastName()
		result := fmt.Sprintf("Fuck you %s %s", firstName, lastName)
		sendErr := stream.Send(&GreetEveryoneRes{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Failed to send message: %v", sendErr)
		}
	}
}
