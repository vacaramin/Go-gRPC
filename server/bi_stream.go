package main

import (
	"io"
	"log"

	pb "github.com/vacaramin/Go-gRPC/proto"
)

func (s *helloServer) sayHelloBidrectionalStreaming(stream pb.GreetService_SayHelloBidrectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return err
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
