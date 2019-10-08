package main

import (
	"context"
	"fmt"
	"log"
	"vs_workspace/Grpc_course/client_stream/greetpb"

	"google.golang.org/grpc"
)

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a client steaming RPC")
	requests := []*greetpb.GreetStreamRequest{
		&greetpb.GreetStreamRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Alex",
			},
		},

		&greetpb.GreetStreamRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Brian",
			},
		},

		&greetpb.GreetStreamRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Cathy",
			},
		},

		&greetpb.GreetStreamRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "David",
			},
		},

		&greetpb.GreetStreamRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Eric",
			},
		},
	}

	stream, err := c.GreetStream(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetStream %v", err)
	}

	for _, req := range requests {
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response %v", err)
	}

	fmt.Printf("GreetStream Response %v\n", res)

}

func main() {
	fmt.Println("Hello ths is a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Created client %v", c)

	doClientStreaming(c)
}
