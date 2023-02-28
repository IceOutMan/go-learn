package main

import (
	"context"
	"flag"
	"fmt"
	pb "grpc_demo/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement service.MaxSize
type server struct {
	pb.UnimplementedMaxSizeServer
}

// SayHello implements service.MaxSize
func (s *server) Echo(ctx context.Context, in *pb.Empty) (*pb.StringMessage, error) {
	return &pb.StringMessage{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMaxSizeServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
