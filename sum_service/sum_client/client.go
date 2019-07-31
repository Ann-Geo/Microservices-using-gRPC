package main

import (
	"context"
	"fmt"
	"log"
	"vs_workspace/Grpc_course/sum_service/sumspb"

	"google.golang.org/grpc"
)

func calculateSum(c sumspb.SumServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &sumspb.SumRequest{
		Input1: 12,
		Input2: 7,
	}

	res, err := c.FindSum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling FindSum RPC: %v", err)
	}
	log.Printf("Response from FindSum: %v", res.Result)
}

func main() {
	fmt.Println("hello I am a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()
	c := sumspb.NewSumServiceClient(conn)
	fmt.Printf("created client, %f", c)

	calculateSum(c)
}
