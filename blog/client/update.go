package main

import (
	"context"
	"log"

	pb "github.com/kingslayer008/gRRPC-Golang/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog in invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Dinesh",
		Title:    "A new title",
		Content:  "Neww content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog is updated")
}
