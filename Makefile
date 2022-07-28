server:
	go run main.go

front:
	go run client/client.go

.PHOXY: server, client