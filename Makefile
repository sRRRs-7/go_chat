greetServer:
	go run main.go -greet

calcServer:
	go run main.go -calculate

greetClient:
	go run client/client.go -greet

greetManyClient:
	go run client/client.go -greetMany

longGreetClient:
	go run client/client.go -longGreet

greetEveryoneClient:
	go run client/client.go -everyone

calcClient:
	go run client/client.go -calculate

calcManyClient:
	go run client/client.go -calcMany

longCalcClient:
	go run client/client.go -longCalc

manyCalcClient:
	go run client/client.go -manyCalc

.PHOXY: greetServer, calcServer, greetClient, greetManyClient, longGreetClient, calcClient, calcManyClient, longCalcClient, manyCalcClient