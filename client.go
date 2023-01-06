package main

import (
	"context"
	"log"

	"github.com/rrd1986/goGrpc/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connewct to gRpc server hosted at 9000 port: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("error calling SayHello from client: %s", err)

	}
	log.Printf("response from the server: %s: ", response)

}
