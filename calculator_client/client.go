package main

import (
	"context"
	"fmt"
	"github.com/mthaler/grpc-calculator-service/calculatorpb"
	"google.golang.org/grpc"
	"io"
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

	//doUnaryRPC(c)
	doServerStreaming(c)
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

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	log.Println("Starting to do server streaming RPC...")
	request := &calculatorpb.PrimeNumberDecompositionRequest{Number: 42}

	stream, err := c.PrimeNumberDecomposition(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling PrimeNumberDecomposition: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from PrimeNumberDecompositin: %v", msg)
	}
}