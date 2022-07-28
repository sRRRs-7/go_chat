greetServer:
	go run main.go -greet

greetClient:
	go run client/client.go -greet

calcServer:
	go run main.go -calculate

calcClient:
	go run client/client.go -calculate

.PHOXY: greetServer, calcServer, greetClient, calcClient