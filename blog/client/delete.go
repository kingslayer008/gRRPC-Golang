package main

import (
	"context"
	"log"

	pb "github.com/kingslayer008/gRRPC-Golang/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog is invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error while deleting %v\n", err)
	}

	log.Println("Blog was deleted")
}
