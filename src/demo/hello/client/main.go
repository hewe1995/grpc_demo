package main

import (
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
	"fmt"
	pb "demo/hello/hello"
	"golang.org/x/net/context"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	conn1, err := grpc.Dial(address, grpc.WithInsecure())
	conn2, err := grpc.Dial(address, grpc.WithInsecure())
	conn3, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)
	c1 := pb.NewHelloClient(conn1)
	c2 := pb.NewHelloClient(conn2)
	c3 := pb.NewHelloClient(conn3)

	arr := make([]pb.HelloClient, 4)
	arr[0] = c
	arr[1] = c1
	arr[2] = c2
	arr[3] = c3


	name := defaultName

	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	c.Say(context.Background(), &pb.HelloRequest{Name: name})
	start := time.Now()
	for count := 0; count < 10000; count++ {
		_, err := arr[count % 4].Say(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}
	end := time.Since(start)
	fmt.Println(end)

}
