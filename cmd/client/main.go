package main

import (
	"context"
	"log"

	"github.com/deep-diver/test1/pkg/pbs"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pbs.NewDummyServiceClient(conn)

	response, err := c.GetHello(context.Background(), &pbs.HelloRequest{Body: "hello"})
	// response, err := c.GetHello(context.Background(), &pr12er.HelloRequest{Body: "hi server!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response)
}
