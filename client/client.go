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

	"github.com/sRRRs-7/go_chat/calc"
	"github.com/sRRRs-7/go_chat/gRPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	greet := flag.Bool("greet", false, "start greet client")
	calculate := flag.Bool("calculate", false, "start calc client")
	flag.Parse()

	conn, err := grpc.Dial("0.0.0.0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	switch {
	case *greet:
		client := gRPC.NewGreetServiceClient(conn)
		fmt.Printf("Create greet client: %f", client)
		fmt.Println()
		greetUnary(client)
	case *calculate:
		client := calc.NewCalcServiceClient(conn)
		fmt.Printf("Create calc client: %f", client)
		fmt.Println()
		calcUnary(client)
	}
}

func greetUnary(c gRPC.GreetServiceClient) {
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

	req := &gRPC.GreetReq{
		Greeting: &gRPC.Greeting{
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