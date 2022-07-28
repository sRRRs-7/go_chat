package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sRRRs-7/go_chat/calc"
	"github.com/sRRRs-7/go_chat/greet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	greetFlag := flag.Bool("greet", false, "start greet client")
	greetManyFlag := flag.Bool("greetMany", false, "start many times greet client")
	longGreetFlag := flag.Bool("longGreet", false, "start long greet client")
	calculateFlag := flag.Bool("calculate", false, "start calc client")
	calcManyFlag := flag.Bool("calcMany", false, "start calc client")
	longCalcFlag := flag.Bool("longCalc", false, "start long calc client")
	flag.Parse()

	conn, err := grpc.Dial("0.0.0.0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	switch {
	case *greetFlag:
		client := greet.NewGreetServiceClient(conn)
		fmt.Printf("Create greet client: %f", client)
		fmt.Println()
		greetUnary(client)

	case *greetManyFlag:
		client := greet.NewGreetServiceClient(conn)
		fmt.Printf("Create many times greet client: %f", client)
		fmt.Println()
		greetServerStream(client)

	case *longGreetFlag:
		client := greet.NewGreetServiceClient(conn)
		fmt.Printf("Create long greet client: %f", client)
		fmt.Println()
		greetClientStream(client)

	case *calculateFlag:
		client := calc.NewCalcServiceClient(conn)
		fmt.Printf("Create calc client: %f", client)
		fmt.Println()
		calcUnary(client)

	case *calcManyFlag:
		client := calc.NewCalcServiceClient(conn)
		fmt.Printf("Create many calc client: %f", client)
		fmt.Println()
		calcServerStream(client)

	case *longCalcFlag:
		client := calc.NewCalcServiceClient(conn)
		fmt.Printf("Create long calc client: %f", client)
		fmt.Println()
		calcClientStream(client)
	}
}

func greetUnary(c greet.GreetServiceClient) {
	fmt.Println("Starting to do a greet unary RPC")

	fmt.Print("Enter first name: ")
	firstName, err := input(os.Stdin, flag.Args()...)
	if err != nil{
		log.Fatalf("Failed to input first_name: %v", err)
	}

	fmt.Print("Enter last name: ")
	lastName, err := input(os.Stdin, flag.Args()...)
	if err != nil{
		log.Fatalf("Failed to input last_name: %v", err)
	}

	req := &greet.GreetReq{
		Greeting: &greet.Greeting{
			FirstName: firstName,
			LastName: lastName,
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to calling RPC")
	}
	log.Printf("Response: %v", res)
}

func greetServerStream(c greet.GreetServiceClient) {
	fmt.Println("Starting to do a greet server stream RPC")

	fmt.Print("Enter first name: ")
	firstName, err := input(os.Stdin, flag.Args()...)
	if err != nil{
		log.Fatalf("Failed to input first_name: %v", err)
	}

	fmt.Print("Enter last name: ")
	lastName, err := input(os.Stdin, flag.Args()...)
	if err != nil{
		log.Fatalf("Failed to input last_name: %v", err)
	}

	req := &greet.GreetManyTimesReq{
		Greeting: &greet.Greeting{
			FirstName: firstName,
			LastName: lastName,
		},
	}
	res, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to calling RPC: %v", err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			log.Fatalf("Failed to reading stream: %v", err)
		}
		log.Printf("Response: %v", msg.GetResult())
	}
}

func greetClientStream(c greet.GreetServiceClient) {
	fmt.Println("Starting to do a client stream RPC")

	var requests []*greet.LongGreetReq
	for i := 0; i < 5; i++ {
		fmt.Print("Enter first name: ")
		firstName, err := input(os.Stdin, flag.Args()...)
		if err != nil{
			log.Fatalf("Failed to input first_name: %v", err)
		}

		fmt.Print("Enter last name: ")
		lastName, err := input(os.Stdin, flag.Args()...)
		if err != nil{
			log.Fatalf("Failed to input last_name: %v", err)
		}

		req := &greet.LongGreetReq{
			Greeting: &greet.Greeting{
				FirstName: firstName,
				LastName: lastName,
			},
		}
		requests = append(requests, req)
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Failed to calling RPC: %v", err)
	}

	for _, req := range requests {
		fmt.Println("Send request: ", req)
		stream.Send(req)
		time.Sleep(500 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response from LongGreet: %v\n", err)
	}
	str := strings.Split(res.Result, "!!")
	for i, s := range str {
		if i == len(str)-1 {
			break
		}
		log.Printf("Response: %d: %v!!\n", i, s)
	}
}

func calcUnary(c calc.CalcServiceClient) {
	fmt.Println("Starting to do a calc unary RPC")

	fmt.Print("Enter num1: ")
	num1, err := input(os.Stdin, flag.Args()...)
	if err != nil{
		log.Fatalf("Failed to input num1: %v", err)
	}

	fmt.Print("Enter num2: ")
	num2, err := input(os.Stdin, flag.Args()...)
	if err != nil{
		log.Fatalf("Failed to input num2: %v", err)
	}

	intNum1, err := strconv.Atoi(num1)
	if err != nil{
		log.Fatalf("Failed to convert type num1: %v", err)
	}
	intNum2, err := strconv.Atoi(num2)
	if err != nil{
		log.Fatalf("Failed to convert type num1: %v", err)
	}

	req := &calc.CalcReq{
		Calculate: &calc.Calculate{
			Num1: int32(intNum1),
			Num2: int32(intNum2),
		},
	}
	res, err := c.Calc(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to calling RPC")
	}
	log.Printf("Response: %v", res)
}

func calcServerStream(c calc.CalcServiceClient) {
	fmt.Println("Starting to do a calc server stream RPC")

	fmt.Print("Enter num1: ")
	num1, err := input(os.Stdin, flag.Args()...)
	if err != nil{
		log.Fatalf("Failed to input num1: %v", err)
	}

	fmt.Print("Enter num2: ")
	num2, err := input(os.Stdin, flag.Args()...)
	if err != nil{
		log.Fatalf("Failed to input num2: %v", err)
	}

	intNum1, err := strconv.Atoi(num1)
	if err != nil{
		log.Fatalf("Failed to convert type num1: %v", err)
	}
	intNum2, err := strconv.Atoi(num2)
	if err != nil{
		log.Fatalf("Failed to convert type num1: %v", err)
	}

	req := &calc.CalcManyTimesReq{
		Calculate: &calc.Calculate{
			Num1: int32(intNum1),
			Num2: int32(intNum2),
		},
	}
	res, err := c.CalcManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to calling RPC")
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			log.Fatalf("Failed to reading stream: %v", err)
		}
		log.Printf("Response: %v", msg.GetResult())
	}
}

func calcClientStream(c calc.CalcServiceClient) {
	fmt.Println("Starting to calculate client stream")

	var requests []*calc.LongCalcsReq
	for i := 0; i < 5; i++ {
		fmt.Print("Enter num1: ")
		num1, err := input(os.Stdin, flag.Args()...)
		if err != nil {
			log.Fatalf("Failed to input num1: %v", err)
		}

		fmt.Print("Enter num2: ")
		num2, err := input(os.Stdin, flag.Args()...)
		if err != nil {
			log.Fatalf("Failed to input num2: %v", err)
		}

		intNum1, err := strconv.Atoi(num1)
		if err != nil{
			log.Fatalf("Failed to convert type num1: %v", err)
		}
		intNum2, err := strconv.Atoi(num2)
		if err != nil{
			log.Fatalf("Failed to convert type num2: %v", err)
		}

		req := &calc.LongCalcsReq{
			Calculate: &calc.Calculate{
				Num1: int32(intNum1),
				Num2: int32(intNum2),
			},
		}
		requests = append(requests, req)
	}

	stream, err := c.LongCalc(context.Background())
	if err != nil {
		log.Fatalf("Failed to calling RPC: %v", err)
	}

	for _, req := range requests {
		fmt.Println("Send request: ", req)
		stream.Send(req)
		time.Sleep(500 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response from LongCalc: %v\n", err)
	}

	fmt.Printf("Response: %v\n", res)
}

func input(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	text := scanner.Text()
	if len(text) == 0 {
		return "", nil
	}
	return text, nil
}