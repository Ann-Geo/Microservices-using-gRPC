package main

import (
	"actual_vsc/Grpc_course/secure_rpc/greetpb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//Authentication files
const CFile string = "/home/research/goworkspace/src/actual_vsc/Grpc_course/secure_rpc/ssl/server.crt"
const KFile string = "/home/research/goworkspace/src/actual_vsc/Grpc_course/secure_rpc/ssl/server.pem"

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
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

	//Enalble Authentication section true or false for tls
	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		certFile := CFile
		keyFile := KFile
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v", sslErr)
			return
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
