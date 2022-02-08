package main

import (
	"context"
	"fmt"
	"greet/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse {
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

	greetpb.RegisterGreetServiceServer(s, &Server{})
}
