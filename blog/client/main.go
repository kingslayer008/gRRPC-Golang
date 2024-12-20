package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/kingslayer008/gRRPC-Golang/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.NewClient(addr, opts...)

	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	id := createBlog(c)
	readBlog(c, id)
	//readBlog(c, "Blah blah")
	updateBlog(c, id)
	//readBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)

}
