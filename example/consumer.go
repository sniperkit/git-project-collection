package main

import (
	"log"

	pb "github.com/leewind/git-project-collection/proto/collector"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCollectorClient(conn)
	r, err := c.GetStarredRepositories(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not Collector: %v", err)
	}
	log.Printf("Collector: %s", r.Repos)
}
