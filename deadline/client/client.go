package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"vs_workspace/Grpc_course/deadline/deadlinepb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func doGreetWithDeadline(c deadlinepb.GreetClient, timeout time.Duration) {
	fmt.Println("Starting to do a Unary RPC with deadline...")
	req := &deadlinepb.GreetRequest{

		FirstName: "Sarah",
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timout was hit! Deadline was exceeded!")
			} else {
				fmt.Printf("unexpeceted error %v\n", statusErr)
			}

		} else {
			log.Fatalf("error while calling GreetWithDeadline RPC: %v", err)
		}

		return
	}
	log.Printf("Response is : %v", res.GetGreeting())
}

func main() {
	fmt.Println("hello I am a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()
	c := deadlinepb.NewGreetClient(conn)
	fmt.Println("created client")

	doGreetWithDeadline(c, 5*time.Second) // should complete

	doGreetWithDeadline(c, 1*time.Second) // should timeout

}
