package main

import (
	"awesomeProject/adapter/output/protos/integrator"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := integrator.NewSumClient(conn)
	sum, err := client.NumbersSum(context.Background(), &integrator.IntegratorRequest{
		FirstNumber:  20,
		SecondNumber: 0,
	})
	if err != nil {
		return
	}

	fmt.Println(sum)
}
