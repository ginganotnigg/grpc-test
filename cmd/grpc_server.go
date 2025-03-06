package cmd

import (
	"log"
	"net"

	"grpc-test/internal/handler"
	hellopb "grpc-test/proto/hello"
	cryptopb "grpc-test/proto/crypto"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func startGRPC() {
    viper.SetConfigFile("./config/config.yaml")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }

	viper.BindEnv("coingecko.url", "COINGECKO_URL")

    // Create a listener on TCP port
    lis, err := net.Listen("tcp", viper.GetString("server.host") + ":" + viper.GetString("server.port"))
    if err != nil {
        log.Fatalln("Failed to listen:", err)
    }

    // Create a gRPC server object
    grpcServer := grpc.NewServer()
    // Create and register HelloServer
	helloHandler := handler.NewHelloServer()
	hellopb.RegisterGreeterServer(grpcServer, helloHandler)
	// Create and register CryptoServer
	cryptoHandler := handler.NewCryptoServer(viper.GetString("coingecko.url"))
	cryptopb.RegisterCryptoServiceServer(grpcServer, cryptoHandler)

    // Serve gRPC server
    log.Println("Serving gRPC on " + viper.GetString("server.host") + ":" + viper.GetString("server.port"))
    go func() {
        log.Fatalln(grpcServer.Serve(lis))
    }()
}