package main

import (
	"log"
	"microservicesWithGo/samples/grpc/hello"
	"microservicesWithGo/samples/grpc/server"
	"net"

	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	hello.RegisterHelloWorldServiceServer(srv, &server.HelloService{})
	listener, _ := net.Listen("tcp", ":8080")
	log.Println("Starting Server...")
	log.Fatal(srv.Serve(listener))
}
