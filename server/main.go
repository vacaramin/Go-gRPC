package main

import (
	"log"
	"net"

	pb "github.com/vacaramin/Go-gRPC/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())

	if err := grpcServer.Serve((lis)); err != nil {
		log.Fatalf("Failed to Start: %v", err)
	}
}
