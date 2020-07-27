package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/toyozaki/sample_grpc/helloworld"
	"google.golang.org/grpc"
)

var address = "localhost:8080"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	c := pb.NewGreeterClient(conn)

	name := "4k1r4"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatal("Could not greet: ", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}
