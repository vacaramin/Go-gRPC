package main

import (
	"log"

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
	//client := pb.NewGreetServiceClient(conn)

	// names := &pb.NamesList{"
	// Names: [string]{"Waqar", "Alice", "Bob"}"}

}
