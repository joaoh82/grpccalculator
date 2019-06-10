package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joaoh82/grpccalculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("This is a gRPC Client!")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to Dial")
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	req := &calculatorpb.SumRequest{
		First_Number:  10,
		Second_Number: 42,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call Sum function: %v", err)
	}
	log.Printf("Response from Sum function: %v", res.Result)
}
