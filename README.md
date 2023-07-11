# Go-Grpc
This project demonstrates the usage of gRPC, a high-performance, open-source framework for remote procedure calls (RPC) developed by Google. gRPC allows communication between different systems and programming languages using a simplified interface definition.

## gRPC Communication Methods 🚀
gRPC provides four main communication patterns:

1. **Unary RPC** This is the traditional request-response method where the client sends a single request to the server and waits for a response.

2. **Server-side Streaming RPC:** In this pattern, the client sends a request to the server, and the server responds with a stream of messages. The client can read these messages from the stream until the server is done.

3. **Client-side Streaming RPC:** Here, the client sends a stream of messages to the server. The server processes these messages and eventually sends a single response back to the client.

4. **Bidirectional Streaming RPC:** This method allows both the client and the server to send and receive streams of messages. The client and server can independently read and write messages on the same connection.

## 🛠️ Setting up a gRPC-Go project
1. Create a new directory for your project and navigate into it:
    ```shell
    mkdir go-grpc
    cd go-grpc
    mkdir client server proto
2. Install the gRPC Go plugin:
    ```shell
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    export PATH="$PATH:$(go env GOPATH)/bin"
3. Initialize a Go module:
    ```shell
    go mod init github.com/your_username/basic-go-grpc
    go mod tidy
4. Create the proto file with the required services and messages in the proto directory.
5. Generate the .pb.go files from the proto file:

*If the greet.proto file is in the proto directory, run:*
    
    ```shell
    protoc --go_out=. --go-grpc_out=. proto/greet.proto
    
*If the greet.proto file contains a different path, adjust the command accordingly.*


6. Create the server and client directories and create the main.go files with the necessary controllers and services.


## 🏃 Running the application

1. Install the dependencies:

    ```shell
    go mod tidy

2. Run the server:

    ```shell
    go run server/main.go

3. Run the client:

    ```shell
    go run client/main.go


**Note: When you want to Run a Particular Communication Method of RPC, Just Simply Uncomment one of the line in Client/main.go**