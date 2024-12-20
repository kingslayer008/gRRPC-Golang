package main

import (
	"context"
	"log"

	pb "github.com/kingslayer008/gRRPC-Golang/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("----readBlog is invoked----")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while reading blog %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)

	return res
}
