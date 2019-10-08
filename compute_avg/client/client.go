package main

import (
	"context"
	"fmt"
	"log"
	"vs_workspace/Grpc_course/compute_avg/avgpb"

	"google.golang.org/grpc"
)

func calculateAvg(c avgpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a client stream RPC")

	inputs := []int64{1, 2, 3, 4}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverage RPC %v\n", err)
	}

	for _, inp := range inputs {

		inpFormatted := &avgpb.InputNumber{
			Number: inp,
		}
		stream.Send(inpFormatted)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving the response %v\n", err)

	}

	fmt.Printf("Average = %v\n", res)

}

func main() {
	fmt.Println("Hello this is a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()

	c := avgpb.NewCalculatorServiceClient(conn)
	fmt.Printf("Created client %v", c)

	calculateAvg(c)
}
