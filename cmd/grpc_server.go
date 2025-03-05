package cmd

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc-test/internal/handler"
	hellopb "grpc-test/proto/hello"
)

func startGRPC() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	grpcServer := grpc.NewServer()
	serverInstance := handler.NewServer()
	// Attach the Greeter service to the server
	hellopb.RegisterGreeterServer(grpcServer, serverInstance)
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(grpcServer.Serve(lis))
	}()
}