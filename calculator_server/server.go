package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/joaoh82/grpccalculator/calculatorpb"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function called with : %v\n", req)

	firstNumber := req.GetFirst_Number()
	secondNumber := req.GetSecond_Number()
	sum := firstNumber + secondNumber

	return &calculatorpb.SumResponse{
		Result: sum,
	}, nil
}

func main() {
	fmt.Println("Hello, I am gRPC Server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
