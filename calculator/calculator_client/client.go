package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/HunnTeRUS/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect", err)
	}

	defer conn.Close()

	c := calculatorpb.NewSumClient(conn)

	//doServerStreaming(c)

	//doClientStreaming(c)

	doErrorUnary(c)
}

func doClientStreaming(c calculatorpb.SumClient) {
	requestNumbers := []int64{
		1, 2, 3, 4, 5, 6, 32, 121, 4242,
	}

	stream, err := c.AverageNumbers(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, x := range requestNumbers {
		stream.Send(&calculatorpb.AverageNumberRequest{
			Number: x,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result is %v \n", res)
}

func doUnary(c calculatorpb.SumClient) {
	request := &calculatorpb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	}

	greet_response, err := c.NumbersSum(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greet_response)
}

func doServerStreaming(c calculatorpb.SumClient) {
	request := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 10,
	}

	greet_response, err := c.PrimeNumberDecomposition(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	for {
		response, err := greet_response.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Response: %d\n", response.GetPrimeFactor())
	}
}

func doErrorUnary(c calculatorpb.SumClient) {
	fmt.Print("Error unary")

	//correct
	doErrorCall(c, 10)

	//error call
	doErrorCall(c, -20)
}

func doErrorCall(c calculatorpb.SumClient, number int32) {
	//correct call
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{
		Number: number,
	})

	if err != nil {
		responseError, ok := status.FromError(err)
		if ok {
			//actual error from grpc (user error)
			fmt.Printf("Error message from server: %v \n", responseError.Message())
			fmt.Printf("Error code from server: %v", responseError.Code())

			if responseError.Code() == codes.InvalidArgument {
				fmt.Println("Negative number")
			}
		} else {
			log.Fatal("Big error calling server", err)
		}

		return
	}

	fmt.Printf("Result of square root: %v \n", res)
}
