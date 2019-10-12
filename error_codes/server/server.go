package main

import (
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
	"fmt"
	"log"
	"net"
	"vs_workspace/Grpc_course/error_codes/errorcpb"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) SquareRoot(ctx context.Context, req *errorcpb.SquareRootRequest) (*errorcpb.SquareRootResponse, error) {

	fmt.Println("SquareRoot RPC was invoked")

	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number %v", number),
		)
	}

	return &errorcpb.SquareRootResponse{
		Root: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	fmt.Println("Starting server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v\n", err)
	}

	s := grpc.NewServer()
	errorcpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}
