package calc

import (
	context "context"
	"fmt"
	"log"
	"net"
	"time"

	grpc "google.golang.org/grpc"
)

type server struct {}

func CalculateServer() {
	fmt.Println("Start gRPC greet server")
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterCalcServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func ( *server) Calc(ctx context.Context, req *CalcReq) (*CalcRes, error) {
	fmt.Println("Request num1: ", req.Calculate.Num1)
	fmt.Println("Request num2: ", req.Calculate.Num2)
	num1 := req.GetCalculate().Num1
	num2 := req.GetCalculate().Num2
	result := num1 + num2
	res := &CalcRes{
		Result: result,
	}
	return res, nil
}

func ( *server) CalcManyTimes(req *CalcManyTimesReq, stream CalcService_CalcManyTimesServer) error {
	fmt.Println("Request num1: ", req.Calculate.Num1)
	fmt.Println("Request num2: ", req.Calculate.Num2)
	num1 := req.GetCalculate().Num1
	num2 := req.GetCalculate().Num2
	result := num1 + num2
	for i := 0; i < 16; i++ {
		res := &CalcRes{
			Result: result,
		}
		result += result
		stream.SendMsg(res)
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

