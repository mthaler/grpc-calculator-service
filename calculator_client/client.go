package main

import (
	"context"
	"fmt"
	"github.com/mthaler/grpc-calculator-service/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {

	fmt.Println("Hello, I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnaryRPC(c)
}

func doUnaryRPC(c calculatorpb.CalculatorServiceClient) {
	log.Println("Starting to do unary RPC...")
	request := &calculatorpb.SumRequest{
		FirstNumber: 3,
		SecondNumnber: 4,
	}

	res, err := c.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling SUM RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res)
}