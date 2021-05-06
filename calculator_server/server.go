package main

import (
	"context"
	"fmt"
	"github.com/mthaler/grpc-calculator-service/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n", request)
	x := request.GetFirstNumber()
	y := request.GetSecondNumnber()
	result := x + y
	response := calculatorpb.SumResponse{
		SumResult: result,
	}
	return &response, nil
}

func (*server) PrimeNumberDecomposition(request *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition function was invoked with %v\n", request)

	n := request.GetNumber()
	divisor := int64(2)

	for n > 1 {
		if n % divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{Factor: divisor})
			n = n / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased: %v", divisor)
		}
	}
	return nil
}

func (*server) Average(stream calculatorpb.CalculatorService_AverageServer) error {
	fmt.Printf("Average function was invoked with %v\n", stream)
	sum := make([]int32, 0)
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			result := 0.0
			for _, n := range sum {
				result += float64(n)
			}
			result /= float64(len(sum))

			// we have finished reading the client stream
			return stream.SendAndClose(&calculatorpb.AverageResponse{Result: result})
		}
		if err != nil {
			log.Fatalf("Error while reading the client stream: %v", err)
		}

		sum = append(sum, request.GetNumber())
	}
}

func main() {
	fmt.Println("Starting calculator service...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
