package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"Reblox/server/proto"
)

type TestServiceServer struct {
	proto.TestServiceServer
}

func (TestServiceServer) Test(ctx context.Context, in *proto.TestMessage) (*proto.TestMessage, error) {
	fmt.Println("Received:", in.GetMessage())
	return &proto.TestMessage{Message: `I received "` + in.GetMessage() + `"!`}, nil
}

func main() {
	fmt.Println("Starting server...")
	listener, err := net.Listen("tcp", ":2006")
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	// start grpc server
	grpcServer := grpc.NewServer()
	proto.RegisterTestServiceServer(grpcServer, TestServiceServer{})

	fmt.Println("Server started on port 2006")
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}
}
