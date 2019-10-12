package main

import (
	"context"
	"fmt"
	"log"
	"vs_workspace/Grpc_course/error_codes/errorcpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func doRPCWithErrorCode(c errorcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a SquareRoot RPC")

	//correct call
	doErrorCall(c, 10)

	fmt.Println("Starting to do a SquareRoot RPC")

	//errorcall
	doErrorCall(c, -1)

}

func doErrorCall(c errorcpb.CalculatorServiceClient, n float32) {
	req := &errorcpb.SquareRootRequest{
		Number: n,
	}

	res, err := c.SquareRoot(context.Background(), req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			//actual error from grpc (user error)
			fmt.Printf("Error message from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("Client sent a negative number")
				return
			}

		} else {
			log.Fatalf("Error while calling SquareRoot RPC")
			return
		}
	}

	fmt.Printf("Square root of %v is %v\n", n, res.GetRoot())

}

func main() {
	fmt.Println("Hello I am a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v\n", err)
	}

	defer conn.Close()
	c := errorcpb.NewCalculatorServiceClient(conn)
	fmt.Printf("Created client %v\n", c)

	doRPCWithErrorCode(c)

}
