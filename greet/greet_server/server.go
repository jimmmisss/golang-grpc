package main

import (
	"fmt"
	"greet/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{
	greetpb.UnimplementedGreetServiceServer
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	} 
}
