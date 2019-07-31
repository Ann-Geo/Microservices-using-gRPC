package main

import (
	"context"
	"fmt"
	"log"
	"github.com/Ann-Geo/Microservices-using-gRPC/greet_unary/greetpb"

	"google.golang.org/grpc"
)

func doUnaryGreet(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sarah",
			LastName:  "Steve",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func main() {
	fmt.Println("hello I am a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("created client, %f", c)

	doUnaryGreet(c)
}
