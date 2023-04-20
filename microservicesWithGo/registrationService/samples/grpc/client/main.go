package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"microservicesWithGo/samples/grpc/hello"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	con, _ := grpc.Dial(":8080", grpc.WithInsecure())
	client := hello.NewHelloWorldServiceClient(con)
	ctx := context.Background()
	answer, err := client.SayHello(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(answer.GetMessageText())
}
