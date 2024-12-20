package main

import (
	"context"
	"log"
	"net"

	pb "github.com/kingslayer008/gRRPC-Golang/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	pb.BlogServiceServer
}

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

func main() {

	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	//client, err := mongo.Connect(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		certfile := "ssl/server.crt"
		keyfile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certfile, keyfile)

		if err != nil {
			log.Fatalf("Failed to load certificate %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)
	pb.RegisterBlogServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
