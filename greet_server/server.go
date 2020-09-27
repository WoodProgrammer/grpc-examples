package main

import (
	"context"
	"fmt"
	"greet/greetpb"
	"log"
	"net"
	"google.golang.org/grpc"
)
type server struct{

}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	log.Print(req)
	firstName := req.GetGreeting().GetFirstName();
	result := "Hello" + firstName

	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main(){
	fmt.Print("Hello World!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Some error %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err:= s.Serve(lis); err != nil {
		log.Fatalf("Some error %v", err)
	}
}