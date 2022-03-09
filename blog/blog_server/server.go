package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/HunnTeRUS/grpc-go/blog/blogpb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct{}

var (
	collection *mongo.Collection
)

type blogItemResponse struct {
	ID       primitive.ObjectID `json:"_id"`
	AuthorID string             `json:"author_id"`
	Content  string             `json:"content"`
	Title    string             `json:"title"`
}

type blogitem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	AuthorID string             `bson:"author_id" json:"author_id"`
	Content  string             `bson:"content" json:"content"`
	Title    string             `bson:"title" json:"title"`
}

func startAndHandleRoutes() {
	router := gin.Default()

	router.POST("/create-blog", CreateBlogHTTPMethod)

	if err := router.Run(":8088"); err != nil {
		panic("Error trying to start application using the port 8088")
	}
}

func main() {

	go startAndHandleRoutes()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	list, err := net.Listen("tcp", "0.0.0.0:50054")
	if err != nil {
		log.Fatalf("%v", err)
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error trying to connect into mongodb: %v", err)
	}

	collection = client.Database("mydb").Collection("blog")

	s := grpc.NewServer()
	reflection.Register(s)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		if err := s.Serve(list); err != nil {
			log.Fatalf("%v", err)
		}
	}()

	//Wait for control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	//Block until a signal is received
	<-ch
	fmt.Println("Stoping the server")
	s.Stop()
	fmt.Println("Stoping listener")
	list.Close()
	fmt.Println("Stoping mongodb")
	client.Disconnect(context.Background())
	fmt.Println("End of program")
}

func CreateBlogHTTPMethod(c *gin.Context) {
	var blogReq blogitem

	if err := c.ShouldBindJSON(&blogReq); err != nil {
		fmt.Println("Error trying to map request body")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := collection.InsertOne(context.Background(), blogReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resObj := &blogItemResponse{
		ID:       res.InsertedID.(primitive.ObjectID),
		AuthorID: blogReq.AuthorID,
		Title:    blogReq.Title,
		Content:  blogReq.Content,
	}

	c.JSON(http.StatusOK, resObj)
	return
}

func (*server) CreateBlog(stream blogpb.BlogService_CreateBlogServer) error {

	fmt.Println("Create blog grpc called")
	for {
		req, err := stream.Recv()

		if req == nil {
			fmt.Println("End of messages")
			return nil
		}

		if err != nil {
			return status.Error(codes.Internal, "Error trying to receive message")
		}

		blog := req.GetBlog()

		data := &blogitem{
			AuthorID: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		}

		res, err := collection.InsertOne(context.Background(), data)
		if err != nil {
			return status.Errorf(codes.Internal, "Error trying to insert data into database: %v", err)
		}
		responseObj := &blogpb.CreateBlogResponse{
			Blog: &blogpb.Blog{
				Id:       res.InsertedID.(primitive.ObjectID).Hex(),
				AuthorId: data.AuthorID,
				Title:    data.Title,
				Content:  data.Content,
			},
		}

		errSendingMessages := stream.SendMsg(responseObj)
		if errSendingMessages != nil {
			return status.Error(codes.Internal, "Error trying to send message")
		}
	}
}
