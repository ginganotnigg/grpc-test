package handler

import (
	"context"

	hellopb "grpc-test/proto/hello"
)


type HelloServer struct {
    hellopb.UnimplementedGreeterServer
}

func NewHelloServer() *HelloServer {
    return &HelloServer{}
}

func (s *HelloServer) SayHello(ctx context.Context, in *hellopb.HelloRequest) (*hellopb.HelloReply, error) {
    return &hellopb.HelloReply{Message: "Hello " + in.Name}, nil
}
