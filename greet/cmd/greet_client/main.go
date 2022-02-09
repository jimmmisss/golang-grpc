package main

import (
	"context"
	"fmt"

	"log"

	"github.com/greet/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := pb.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c pb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary gRPC...")
	req := &pb.GreetRequest{
		Greeting: &pb.Greeting{
			FirstName: "Wesley",
			LastName: "Pereira",
		},
	}
	
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet gRPC: %v", err)
	}
	log.Printf("Response Greet: %v", res.Result)
}
