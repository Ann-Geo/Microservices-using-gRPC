package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"vs_workspace/Grpc_course/find_max/maxipb"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) FindMaximum(stream maxipb.MaximService_FindMaximumServer) error {
	fmt.Println("Starting the server")

	var count int64
	var maxi int64

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		count++
		input := req.GetNumber()
		if count == 1 {
			maxi = input
			sendErr := stream.Send(&maxipb.OutputResponse{
				Maximum: input,
			})

			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return sendErr
			}
		} else {
			if input > maxi {
				maxi = input
				sendErr := stream.Send(&maxipb.OutputResponse{
					Maximum: input,
				})
				if sendErr != nil {
					log.Fatalf("Error while sending data to client: %v", sendErr)
					return sendErr
				}
			} else {
				sendErr := stream.Send(&maxipb.OutputResponse{
					Maximum: maxi,
				})

				if sendErr != nil {
					log.Fatalf("Error while sending data to client: %v", sendErr)
					return sendErr
				}
			}
		}

	}

}

func main() {
	fmt.Println("Starting the server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v\n", err)
	}
	s := grpc.NewServer()
	maxipb.RegisterMaximServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

}
