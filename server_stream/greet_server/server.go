package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
	"github.com/Ann-Geo/Microservices-using-gRPC/server_stream/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetStream(req *greetpb.GreetRequest, stream greetpb.GreetService_GreetStreamServer) error {
	fmt.Printf("GreetStream function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetStreamResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
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
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
