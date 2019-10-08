package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"vs_workspace/Grpc_course/client_stream/greetpb"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) GreetStream(stream greetpb.GreetService_GreetStreamServer) error {
	fmt.Println("GreetStream RPC was invoked")
	result := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.GreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello" + firstName + "!"
	}
}

func main() {
	fmt.Println("Starting server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
