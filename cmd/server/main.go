// ./cmd/server/main.go

package main

import (
	"log"
	"net"

	"github.com/koyeb/example-grpc-api/internal"
	"github.com/koyeb/example-grpc-api/protogen/golang/orders"
	"google.golang.org/grpc"
)

func main() {
	const addr = "0.0.0.0:50051"

	// create a TCP listener on the specified port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server instance
	server := grpc.NewServer()

	// create a order service instance with a reference to the db
	db := internal.NewDB()
	orderService := internal.NewOrderService(db)

	// register the order service with the grpc server
	orders.RegisterOrdersServer(server, &orderService)

	// start listening to requests
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
