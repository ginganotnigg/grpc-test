package handler

import (
	"context"
	hellopb "grpc-test/proto/hello"
)

type server struct{
	hellopb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *hellopb.HelloRequest) (*hellopb.HelloReply, error) {
	return &hellopb.HelloReply{Message: "Hello " + in.Name}, nil
}