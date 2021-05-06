package main

import (
	"context"
	"fmt"
	"github.com/mthaler/grpc-calculator-service/calculatorpb"
	"google.golang.org/grpc"
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
