package main

import (
	"context"
	"log"
	"time"

	pb "server/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:5100"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// client
	c := pb.NewFactorialClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	testValues := []int64{-2, 0, 1, 2, 3, 4, 5, 6, 10, 20, 30, 1000, 10000000} //9223372036854775807

	r, err := c.Calculate(ctx, &pb.CalculateRequest{Numbers: testValues, MinSpread: &pb.SpreadData{}})
	if err != nil {
		log.Fatalf("could not calculate: %v", err)
	}

	log.Println("Calculation results:")

	for range testValues {
		result, err := r.Recv()
		if err != nil {
			log.Printf("timeout getting output: %v", err)
			break
		}

		log.Printf("input: %d  output: %s", result.InputNumber, result.FactorialResult)
	}

}
