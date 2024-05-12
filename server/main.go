package main

import (
	p "Reblox/server/proto"
	"Reblox/shared/keys"
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type EventService struct{ p.EventServiceServer }

func (EventService) Test(ctx context.Context, in *p.SignedBaseEvent) (*p.Response, error) {
	fmt.Printf("Id:      %x\n", in.Id)
	fmt.Printf("Payload: %x\n", in.Payload)

	unsigned2 := &p.UnsignedBaseEvent{}
	proto.Unmarshal(in.Payload, unsigned2)

	pubb2 := unsigned2.Pubkey
	pub2, err := keys.Encode(pubb2)
	if err != nil {
		fmt.Println("Failed to encode pubkey:", err)
		return nil, err
	}
	fmt.Println("Pubkey:  ", pub2)

	time := time.UnixMilli(unsigned2.Created)
	fmt.Println("Created: ", time)
	return &p.Response{Message: "Message received from " + pub2 + "!"}, nil
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
	p.RegisterEventServiceServer(grpcServer, EventService{})

	fmt.Println("Server started on port 2006")
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}
}
