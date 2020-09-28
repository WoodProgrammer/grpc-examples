package main

import (
	"context"
	"greet/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"io"
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
	/*req := &greetpb.N{
		Greeting: &greetpb.Greeting {
			FirstName: "Mustafa",
			LastName: "Kandirali",

		},
	}
	//res, err := c.Greet(context.Background(), req)

	if err != nil {
		fmt.Print("Some error %f", err)
	}
	log.Print("Response from greet : %v", res.Result)*/

	//doUnary(c)
	doServerStreaming(c)
}

func doServerStreaming(c greetpb.GreetServiceClient){
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephane",
			LastName:  "Maarek",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephane",
			LastName:  "Maarek",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
