package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vacaramin/Go-gRPC/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client Streaming Started")
	stream, err := client.SayhelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("couldn't send names: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloResponse{
			Message: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the request with name %s", name)
		time.Sleep(time.Second * 2)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming Finished")
	if err != nil {
		log.Fatalf("Error while recieving %v", err)
	}
	log.Printf("%v", res.Messages)
}
