package main

import (
	"fmt"
	"log"
	"net"

	pb "server/proto"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"math/big"
)

type server struct {
	pb.UnimplementedFactorialServer
}

func factorialInt(v int64) string {

	if v < 0 {
		return "Number should be greater or equal to 0"
	}

	if v == 0 {
		return "0"
	}

	calc := big.NewInt(1)
	calc.MulRange(1, v)

	return calc.String()
}

func factorialFloat(v int64) string {

	var res float64 = 1.0

	var i int64

	for i = 1; i <= v; i++ {
		res = res * float64(i)
	}

	return fmt.Sprintf("%g", res)
}

func (s *server) factorial(v int64) string {

	if v < 100 {
		return factorialInt(v)
	} else {
		return factorialFloat(v)
	}

}

func (s *server) Calculate(in *pb.CalculateRequest, stream pb.Factorial_CalculateServer) error {

	log.Printf("Received: %v", in.Numbers)

	for _, v := range in.Numbers {

		log.Printf("Input value: %d", v)

		res := pb.CalculateResult{
			InputNumber:     v,
			FactorialResult: s.factorial(v),
		}

		log.Printf("Calculated: %v", res.FactorialResult)

		err := stream.Send(&res)
		if err != nil {
			log.Printf("Error sening response: %v", err)
			break
		}

	}

	return nil
}

func main() {

	// Setup config
	viper.SetConfigName("server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Reading config
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error reading config: %v", err)
	}

	viper.SetDefault("port", ":5100")
	port := viper.GetString("port")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFactorialServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
