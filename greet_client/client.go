package main

import (
	"context"
	"greet/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main(){

	fmt.Println("Hello I'm client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Some error %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	fmt.Print("Client created %f",c)
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting {
			FirstName: "Mustafa",
			LastName: "Kandirali",

		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		fmt.Print("Some error %f", err)
	}
	log.Print("Response from greet : %v", res.Result)

}

