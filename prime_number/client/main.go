package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"vs_workspace/Grpc_course/prime_number/primepb"

	"google.golang.org/grpc"
)

func doServerStreaming(c primepb.PrimeServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &primepb.InputRequest{
		InputNumber: 210,
	}

	resStream, err := c.PrimeDecomp(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeDecomp RPC: %v", err)
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
		log.Printf("Response from PrimeDecomp: %v", msg.GetOutputNumber())
	}

}

func main() {
	fmt.Println("hello I am a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()
	c := primepb.NewPrimeServiceClient(conn)
	fmt.Printf("created client, %f", c)

	doServerStreaming(c)
}
