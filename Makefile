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
greetDeadlineClient:
	go run client/client.go -deadline

calcClient:
	go run client/client.go -calculate
calcManyClient:
	go run client/client.go -calcMany
longCalcClient:
	go run client/client.go -longCalc
manyCalcClient:
	go run client/client.go -manyCalc

.PHOXY: greetServer, calcServer
.PHOXY: greetClient, greetManyClient, greetDeadlineClient, longGreetClient, greetDeadlineClient
.PHOXY: calcClient, calcManyClient, longCalcClient, manyCalcClient