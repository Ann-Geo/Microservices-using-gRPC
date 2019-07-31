package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
	"github.com/Ann-Geo/Microservices-using-gRPC/bidi1/basicpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello I am a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()
	c := basicpb.NewBidiGreetServiceClient(conn)
	fmt.Printf("created client, %f", c)

	doBidiStreaming(c)
}

func doBidiStreaming(c basicpb.BidiGreetServiceClient) {
	fmt.Println("Starting to do a streaming RPC")

	//create a stream by invoking the client
	stream, err := c.BidiStreamGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*basicpb.BidiStreamGreetRequest{
		&basicpb.BidiStreamGreetRequest{
			Greeting: &basicpb.BidiGreeting{
				FirstName: "Sarah",
			},
		},
		&basicpb.BidiStreamGreetRequest{
			Greeting: &basicpb.BidiGreeting{
				FirstName: "John",
			},
		},
		&basicpb.BidiStreamGreetRequest{
			Greeting: &basicpb.BidiGreeting{
				FirstName: "Lucy",
			},
		},
		&basicpb.BidiStreamGreetRequest{
			Greeting: &basicpb.BidiGreeting{
				FirstName: "Mark",
			},
		},
		&basicpb.BidiStreamGreetRequest{
			Greeting: &basicpb.BidiGreeting{
				FirstName: "Piper",
			},
		},
	}

	waitc := make(chan struct{})

	//send a bunch of messages to the client
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	//receive a bunch of messages from the client
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v", res)
		}
		close(waitc)
	}()

	//block until everything is done
	<-waitc
}
