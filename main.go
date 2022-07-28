package main

import (
	"flag"

	"github.com/sRRRs-7/go_chat/calc"
	"github.com/sRRRs-7/go_chat/gRPC"
)

func main() {
	greet := flag.Bool("greet", false, "start greet server")
	calculate := flag.Bool("calculate", false, "start calculate server")
	flag.Parse()

	switch {
	case *greet:
		gRPC.GreetServer()
	case *calculate:
		calc.CalculateServer()
	}
}