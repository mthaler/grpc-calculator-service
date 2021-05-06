package main

import (
	"context"
	"fmt"
	"github.com/mthaler/grpc-calculator-service/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
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
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)
}

func doUnaryRPC(c calculatorpb.CalculatorServiceClient) {
	log.Println("Starting to do unary RPC...")
	request := &calculatorpb.SumRequest{
		FirstNumber:  3,
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

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	log.Println("Starting to do client streaming RPC...")

	numbers := []int32{1, 3, 6, 2, 10}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, n := range numbers {
		fmt.Printf("Sending request: %v\n", n)
		stream.Send(&calculatorpb.AverageRequest{Number: n})
		time.Sleep(100 * time.Millisecond)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}
	fmt.Printf("Long greet respone: %v", response)
}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	log.Println("Starting to do BiDi streaming RPC...")

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			stream.Send(&calculatorpb.FindMaximumRequest{Number: number})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				// we've reached the end of the stream
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res)
		}
		close(waitc)
	}()

	<- waitc
}
