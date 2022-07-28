greetServer:
	go run main.go -greet

greetClient:
	go run client/client.go -greet

manyGreetClient:
	go run client/client.go -greetMany

calcServer:
	go run main.go -calculate

calcClient:
	go run client/client.go -calculate

calcManyClient:
	go run client/client.go -calcMany

.PHOXY: greetServer, calcServer, manyGreetClient, greetClient, calcClient, calcManyClient