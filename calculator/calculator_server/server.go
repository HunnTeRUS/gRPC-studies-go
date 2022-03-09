package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"github.com/HunnTeRUS/grpc-go/calculator/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	fmt.Println("Received number for average numbers")
	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative number: %d", number),
		)
	}

	return &calculatorpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func (*server) AverageNumbers(stream calculatorpb.Sum_AverageNumbersServer) error {
	fmt.Println("Received number for average numbers")
	numbers := []int64{}

	for {
		stream_value, err := stream.Recv()

		if err == io.EOF {
			sum_values := int64(0)
			for i := 0; i < len(numbers); i++ {
				sum_values += numbers[i]
			}

			avg_value := sum_values / int64(len(numbers))
			fmt.Println("Average number: %d", avg_value)

			return stream.SendAndClose(&calculatorpb.AverageNumberResponse{
				Number: avg_value,
			})
		}

		if err != nil {
			log.Fatal("Error trying to receive values %v", err)
		}

		numbers = append(numbers, stream_value.GetNumber())
	}

}

func (*server) PrimeNumberDecomposition(
	req *calculatorpb.PrimeNumberDecompositionRequest,
	stream calculatorpb.Sum_PrimeNumberDecompositionServer,
) error {
	fmt.Println("Received prime number decomposition request: %v", req)

	divisor := int64(2)
	number := req.GetNumber()

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Println("Divisor increased")
		}
	}

	return nil
}

func (*server) NumbersSum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Println("Received calculator request: %v", req)
	response := &calculatorpb.SumResponse{
		SumResponseCount: req.FirstNumber + req.SecondNumber,
	}

	return response, nil
}

func main() {
	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("%v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterSumServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("%v", err)
	}
}
