package cmd

import (
    "context"
    "log"
    "net/http"

    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "github.com/spf13/viper"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"

    hellopb "grpc-test/proto/hello"
	cryptopb "grpc-test/proto/crypto"
)

func startGateway() {
    viper.SetConfigFile("./config/config.yaml")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }

    conn, err := grpc.NewClient(
        viper.GetString("server.host") + ":" + viper.GetString("server.port"),
        grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    if err != nil {
        log.Fatalln("Failed to dial server:", err)
    }
    gwmux := runtime.NewServeMux()
    // Register Greeter
    err = hellopb.RegisterGreeterHandler(context.Background(), gwmux, conn)
    if err != nil {
        log.Fatalln("Failed to register gateway:", err)
    }
	// Register CryptoService
    err = cryptopb.RegisterCryptoServiceHandler(context.Background(), gwmux, conn)
    if err != nil {
        log.Fatalln("Failed to register gateway:", err)
    }

    gwServer := &http.Server{
        Addr:    ":" + viper.GetString("server.gwport"),
        Handler: gwmux,
    }

    log.Println("Serving gRPC-Gateway on http://" + viper.GetString("server.host") + ":" + viper.GetString("server.gwport"))
    log.Fatalln(gwServer.ListenAndServe())
}