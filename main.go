package main

import (
	"flag"

	"github.com/sRRRs-7/go_chat/calc"
	"github.com/sRRRs-7/go_chat/greet"
)

func main() {
	greetFlag := flag.Bool("greet", false, "start greet server")
	calculateFlag := flag.Bool("calculate", false, "start calculate server")
	flag.Parse()

	switch {
	case *greetFlag:
		greet.GreetServer()
	case *calculateFlag:
		calc.CalculateServer()
	}
}