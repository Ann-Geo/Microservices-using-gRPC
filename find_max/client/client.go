package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"vs_workspace/Grpc_course/find_max/maxipb"

	"google.golang.org/grpc"
)

func doBidiStreaming(c maxipb.MaximServiceClient) {
	fmt.Println("Starting to do a bidi streaming RPC")

	inputs := []int64{1, 0, 7, 2, 3, 8, 5, 6}

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while calling FindMaximum RPC %v\n", err)
		return
	}

	waitc := make(chan struct{})

	go func() {
		for _, inp := range inputs {
			stream.Send(&maxipb.InputRequest{
				Number: inp,
			})

		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving :%v\n", err)
			}
			fmt.Printf("Maximum is %v\n", res.GetMaximum())
		}
		close(waitc)
	}()

	<-waitc
}

func main() {
	fmt.Println("hello this is a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v\n", err)
	}

	defer conn.Close()
	c := maxipb.NewMaximServiceClient(conn)
	fmt.Printf("Created client %v\n", c)

	doBidiStreaming(c)
}
