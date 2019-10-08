package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"vs_workspace/Grpc_course/compute_avg/avgpb"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) ComputeAverage(stream avgpb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("ComputeAverage RPC was invoked")
	var result float64
	var sum int64
	var count int64

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result = float64(sum) / float64(count)
			return stream.SendAndClose(&avgpb.OutputResponse{
				Average: result,
			})
		}

		input := req.GetNumber()
		count++
		sum += input
	}
}

func main() {
	fmt.Println("Starting server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	avgpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
