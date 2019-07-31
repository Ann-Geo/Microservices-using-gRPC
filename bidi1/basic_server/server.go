package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"vs_workspace/bidi1/basicpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) BidiStreamGreet(stream basicpb.BidiGreetService_BidiStreamGreetServer) error {
	fmt.Printf("BidiStreamGreet was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello" + firstName + "!"
		sendErr := stream.Send(&basicpb.BidiGreetStreamResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
}

func main() {
	fmt.Println("hello I am the server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)

	}

	s := grpc.NewServer()
	basicpb.RegisterBidiGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
