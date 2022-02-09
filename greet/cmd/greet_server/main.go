package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/greet/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGreetServiceServer
}

func (*Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", in)
	
	firstName := in.GetGreeting().GetFirstName()
	
	result := "Hello " + firstName
	
	res := &pb.GreetResponse {
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

	pb.RegisterGreetServiceServer(s, &Server{})
}
