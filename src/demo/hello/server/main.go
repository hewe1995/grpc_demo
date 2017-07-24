package main

import (
	"golang.org/x/net/context"
	pb "demo/hello/hello"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
}

func (s *server) Say(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "nihao" + request.Name}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("err ocuur: ", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println("err occur: ",err)
	}
	fmt.Println("serving...")
}