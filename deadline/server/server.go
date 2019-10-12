package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	"vs_workspace/Grpc_course/deadline/deadlinepb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetWithDeadline(ctx context.Context, req *deadlinepb.GreetRequest) (*deadlinepb.GreetResponse, error) {
	fmt.Printf("GreetWithDeadline RPC was invoked with %v\n", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			//the client canceled the request
			fmt.Println("The client canceled the request")
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}

	firstName := req.GetFirstName()
	result := "Hello " + firstName
	res := &deadlinepb.GreetResponse{
		Greeting: result,
	}
	return res, nil
}

func main() {
	fmt.Println("hello I am the server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)

	}

	s := grpc.NewServer()
	deadlinepb.RegisterGreetServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
