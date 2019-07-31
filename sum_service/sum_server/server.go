package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"vs_workspace/Grpc_course/sum_service/sumspb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) FindSum(ctx context.Context, req *sumspb.SumRequest) (*sumspb.SumResponse, error) {
	fmt.Printf("FindSum function was invoked with %v\n", req)

	result := req.GetInput1() + req.GetInput2()
	res := &sumspb.SumResponse{
		Result: result,
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
	sumspb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
