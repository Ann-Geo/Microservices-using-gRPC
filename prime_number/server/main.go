package main

import (
	"fmt"
	"log"
	"net"
	"vs_workspace/Grpc_course/prime_number/primepb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeDecomp(req *primepb.InputRequest, stream primepb.PrimeService_PrimeDecompServer) error {
	fmt.Printf("PrimeDecomp function was invoked with %v\n", req)
	input := req.GetInputNumber()

	var startno int32 = 2

	for input > 1 {
		if input%startno == 0 { // if k evenly divides into N
			fmt.Println(startno) // this is a factor
			res := &primepb.PrimeNumbers{
				OutputNumber: startno,
			}
			stream.Send(res)
			input = input / startno // divide N by k so that we have the rest of the number left.
		} else {
			startno = startno + 1
		}
	}

	return nil
}

func main() {
	fmt.Println("hello I am the server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)

	}

	s := grpc.NewServer()
	primepb.RegisterPrimeServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
