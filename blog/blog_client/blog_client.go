package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/HunnTeRUS/grpc-go/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type blogItemResponse struct {
	ID       primitive.ObjectID `json:"_id"`
	AuthorID string             `json:"author_id"`
	Content  string             `json:"content"`
	Title    string             `json:"title"`
}

type blogitem struct {
	ID       primitive.ObjectID `json:"_id"`
	AuthorID string             `json:"author_id"`
	Content  string             `json:"content"`
	Title    string             `json:"title"`
}

func main() {
	requests := 1000000

	doRequestsWithGRPC(requests)
	//doRequestsWithREST(requests)
}

func doRequestsWithREST(requests int) {
	startTime := time.Now()
	for i := 0; i < requests; i++ {
		requestObj := blogitem{
			Title:    "otavio HTTP request",
			AuthorID: "test 123",
			Content:  "test http request",
		}

		requestByte, _ := json.Marshal(requestObj)
		resp, err := http.Post("http://localhost:8088/create-blog", "application/json", bytes.NewReader(requestByte))

		if err != nil {
			fmt.Printf("Error trying to create blog, error: %v \n", err)
			return
		}

		defer resp.Body.Close()

		var blogRes blogItemResponse
		closer, _ := io.ReadAll(resp.Body)

		if err := json.Unmarshal(closer, &blogRes); err != nil {
			fmt.Printf("Error trying to convert received blog, error: %v \n", err)
			return
		}

		fmt.Println(blogRes)
	}
	endTime := time.Now()

	fmt.Printf("Took %d Milliseconds to finish the %d inserts with REST \n", endTime.Sub(startTime).Milliseconds(), requests)
}

func doRequestsWithGRPC(requests int) {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server, error: %v", err)
	}

	defer conn.Close()

	c := blogpb.NewBlogServiceClient(conn)
	stream, err := c.CreateBlog(context.Background())

	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("Sending message")
		requestMsg := &blogpb.CreateBlogRequest{Blog: &blogpb.Blog{
			Title:    "test grpc",
			AuthorId: "otavio",
			Content:  "test",
		},
		}
		for i := 0; i < requests; i++ {
			stream.SendMsg(requestMsg)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			response := &blogpb.CreateBlogResponse{}
			err := stream.RecvMsg(response)

			if err == io.EOF {
				fmt.Printf("End of messages, err: %v \n", err)
				endTime := time.Now()
				fmt.Printf("Took %d Milliseconds to finish the %d inserts with gRPC \n", endTime.Sub(startTime).Milliseconds(), requests)
				wg.Done()

				break
			}

			fmt.Printf("Response from server: %#v \n", response.Blog)
		}
	}()

	wg.Wait()
}
