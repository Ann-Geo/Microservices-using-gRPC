package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"vs_workspace/Grpc_course/server_stream/greetpb"

	"google.golang.org/grpc"
)

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sarah",
			LastName:  "Mathew",
		},
	}

	resStream, err := c.GreetStream(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetStream RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetStream: %v", msg.GetResult())
	}

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

	doServerStreaming(c)
}
