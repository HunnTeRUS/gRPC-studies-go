package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/HunnTeRUS/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello")

	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	//doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBidiStreaming(c)

	//doCallWithDeadline(c, 1*time.Second)
	//doCallWithDeadline(c, 5*time.Second)
}

func doBidiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Bidi streaming")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatal("Error while creating stream", err)
		return
	}

	request := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Otavio",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Celestino",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Santos",
			},
		},
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for _, x := range request {
			fmt.Printf("Sending message %v \n", x)
			stream.SendMsg(x)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			response := &greetpb.GreetEveryoneResponse{}
			err := stream.RecvMsg(response)

			if err == io.EOF {
				log.Fatalf("End of messages, err: %v", err)
				wg.Done()
				break
			}

			if err != nil {
				log.Fatal("Error while receiving messages", err)
				wg.Done()
				break
			}

			fmt.Printf("Response from server: %#v \n", response.Result)
		}
	}()

	wg.Wait()
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client streaming")

	request := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Otavio",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Celestino",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Santos",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatal("error while calling LongGreet %v", err)
	}

	for _, req := range request {
		fmt.Println("Sending req %v", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response received: %v", res)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming")

	request := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Otavio",
			LastName:  "Celestino",
		},
	}

	result, err := c.GreetManyTimes(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	for {
		message, err := result.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Response from greet many times: %#v", message.GetResult())
	}
}

func doUnary(c greetpb.GreetServiceClient) {
	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Otavio",
			LastName:  "Celestino",
		},
	}

	greet_response, err := c.Greet(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greet_response)
}

func doCallWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	request := &greetpb.GreetWithDeadLineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Otavio",
			LastName:  "Celestino",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	greet_response, err := c.GreetWithDeadLine(ctx, request)
	if err != nil {

		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout has hit! Deadline was exceeded")
			} else {
				fmt.Printf("unexpected error: %v \n", statusErr)
			}
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println(greet_response)
}
