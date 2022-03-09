package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/HunnTeRUS/grpc-go/greet/greetpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone function was invoked")
	for {
		req, err := stream.Recv()

		if req == nil {
			fmt.Println("End of messages")
			break
		}

		if err != nil {
			return status.Error(codes.Internal, "Error trying to receive message")
		}

		first_name := req.GetGreeting().GetFirstName()
		result := "Hello " + first_name

		fmt.Printf("Sending back the message: %v \n", result)
		errSendingMessages := stream.SendMsg(&greetpb.GreetEveryoneResponse{
			Result: result,
		})

		if errSendingMessages != nil {
			return status.Error(codes.Internal, "Error trying to send message")
		}
	}

	return nil
}

func (*server) GreetWithDeadLine(ctx context.Context, req *greetpb.GreetWithDeadLineRequest) (*greetpb.GreetWithDeadLineResponse, error) {

	fmt.Printf("GreetWithDeadLine function was invoked")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("THe client Canceled the request")
			return nil, status.Error(codes.DeadlineExceeded, "THe client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}

	res := &greetpb.GreetWithDeadLineResponse{
		Result: "Test deadline",
	}

	return res, nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatal(err)
		}

		first_name := req.GetGreeting().GetFirstName()
		result += "Hello " + first_name + "! "
	}

	return nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("Greet many times function was invoked %v", req)
	firstname := req.Greeting.FirstName

	for {
		result := "Hello " + firstname
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}

		stream.Send(res)
		time.Sleep(1 * time.Second)
	}
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked %v", req)
	firstname := req.GetGreeting().GetFirstName()

	result := "Hello " + firstname
	response := &greetpb.GreetResponse{
		Result: result,
	}

	return response, nil
}

func main() {

	list, err := net.Listen("tcp", "0.0.0.0:50054")
	if err != nil {
		log.Fatalf("%v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatalf("%v", err)
	}

}
