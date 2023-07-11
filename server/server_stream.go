package main

import (
	"log"
	"time"

	pb "github.com/vacaramin/Go-gRPC/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names: %v", req.Name)

	for _, name := range req.Name {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}
