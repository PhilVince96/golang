package server

import (
	"context"
	"microservicesWithGo/samples/grpc/hello"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type HelloService struct {
	hello.UnimplementedHelloWorldServiceServer // embedded struct
}

func (hs *HelloService) SayHello(context.Context, *empty.Empty) (*hello.HelloMessage, error) {
	msg := &hello.HelloMessage{MessageText: "Hello World"}
	return msg, nil
}
