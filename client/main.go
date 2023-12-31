package main

import (
	"log"

	pb "github.com/vacaramin/Go-gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Waqar", "Alice", "Bob"},
	}
	//callSayHello(client) // UNARY
	//callSayHelloServerStream(client, names) // SERVER STREAM
	//callSayHelloClientStream(client, names) // CLIENT STREAM
	callSayHelloBidirectionalStream(client, names) //  BI_DIRECTIONAL STREAM

}
