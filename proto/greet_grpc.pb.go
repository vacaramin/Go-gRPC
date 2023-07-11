// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/greet.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error)
	SayHelloServerStreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (GreetService_SayHelloServerStreamingClient, error)
	SayHelloClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloClientStreamingClient, error)
	SayHelloBidrectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloBidrectionalStreamingClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/greet_service.GreetService/sayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) SayHelloServerStreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (GreetService_SayHelloServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/greet_service.GreetService/sayHelloServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayHelloServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_SayHelloServerStreamingClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceSayHelloServerStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayHelloServerStreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) SayHelloClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], "/greet_service.GreetService/SayHelloClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayHelloClientStreamingClient{stream}
	return x, nil
}

type GreetService_SayHelloClientStreamingClient interface {
	Send(*HelloResponse) error
	CloseAndRecv() (*MessagesList, error)
	grpc.ClientStream
}

type greetServiceSayHelloClientStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayHelloClientStreamingClient) Send(m *HelloResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceSayHelloClientStreamingClient) CloseAndRecv() (*MessagesList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessagesList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) SayHelloBidrectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloBidrectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[2], "/greet_service.GreetService/sayHelloBidrectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayHelloBidrectionalStreamingClient{stream}
	return x, nil
}

type GreetService_SayHelloBidrectionalStreamingClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceSayHelloBidrectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayHelloBidrectionalStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceSayHelloBidrectionalStreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	SayHello(context.Context, *NoParam) (*HelloResponse, error)
	SayHelloServerStreaming(*NamesList, GreetService_SayHelloServerStreamingServer) error
	SayHelloClientStreaming(GreetService_SayHelloClientStreamingServer) error
	SayHelloBidrectionalStreaming(GreetService_SayHelloBidrectionalStreamingServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) SayHello(context.Context, *NoParam) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloServerStreaming(*NamesList, GreetService_SayHelloServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStreaming not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloClientStreaming(GreetService_SayHelloClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStreaming not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloBidrectionalStreaming(GreetService_SayHelloBidrectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBidrectionalStreaming not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet_service.GreetService/sayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).SayHello(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_SayHelloServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NamesList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).SayHelloServerStreaming(m, &greetServiceSayHelloServerStreamingServer{stream})
}

type GreetService_SayHelloServerStreamingServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greetServiceSayHelloServerStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayHelloServerStreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_SayHelloClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).SayHelloClientStreaming(&greetServiceSayHelloClientStreamingServer{stream})
}

type GreetService_SayHelloClientStreamingServer interface {
	SendAndClose(*MessagesList) error
	Recv() (*HelloResponse, error)
	grpc.ServerStream
}

type greetServiceSayHelloClientStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayHelloClientStreamingServer) SendAndClose(m *MessagesList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceSayHelloClientStreamingServer) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_SayHelloBidrectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).SayHelloBidrectionalStreaming(&greetServiceSayHelloBidrectionalStreamingServer{stream})
}

type GreetService_SayHelloBidrectionalStreamingServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greetServiceSayHelloBidrectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayHelloBidrectionalStreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceSayHelloBidrectionalStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet_service.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sayHello",
			Handler:    _GreetService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sayHelloServerStreaming",
			Handler:       _GreetService_SayHelloServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloClientStreaming",
			Handler:       _GreetService_SayHelloClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "sayHelloBidrectionalStreaming",
			Handler:       _GreetService_SayHelloBidrectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/greet.proto",
}
