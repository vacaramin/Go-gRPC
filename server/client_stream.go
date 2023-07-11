package main

import (
	"io"
	"log"

	pb "github.com/vacaramin/Go-gRPC/proto"
)

func (s *helloServer) SayhelloClientStreaming(stream pb.GreetService_SayhelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf(("Got request with name:%v"), req.Message)
		messages = append(messages, "hello ", req.Message)
	}
}
