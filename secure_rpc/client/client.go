package main

import (
	"actual_vsc/Grpc_course/secure_rpc/greetpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//Authentication
const CFile string = "/home/research/goworkspace/src/actual_vsc/Grpc_course/secure_rpc/ssl/ca.crt"

func doUnaryGreet(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sarah",
			LastName:  "Steve",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func main() {
	fmt.Println("hello I am a client")

	//Enable Authentication - true or false
	tls := true
	opts := grpc.WithInsecure()
	if tls {
		certFile := CFile
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
		}

		opts = grpc.WithTransportCredentials(creds)
	}

	conn, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	fmt.Println("created client")

	doUnaryGreet(c)
}
