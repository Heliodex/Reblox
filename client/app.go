package main

import (
	"context"
	"fmt"
	"time"

	p "Reblox/server/proto"
	"Reblox/shared/keys"

	"golang.org/x/crypto/sha3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

type App struct {
	ctx context.Context
}

// startup is called when the app starts. The context is saved so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

const sk = "resec:eco44kv2fj-8md68h4f0s-vut2oybxpr-idpkdjinuj-0l2z7fqmlc"

func (a *App) ContactServer(message string) (string, error) {
	conn, err := grpc.Dial("localhost:2006", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := p.NewEventServiceClient(conn)

	skBytes, err := keys.DecodeFormatted(sk)
	if err != nil {
		fmt.Println("Failed to decode secret key:", err)
		return "", err
	}
	pkBytes, err := keys.PublicKey(skBytes)
	if err != nil {
		fmt.Println("Failed to get public key:", err)
		return "", err
	}

	unsigned := &p.UnsignedBaseEvent{
		Pubkey:  pkBytes,
		Created: time.Now().UnixMilli(),
	}

	payload, err := proto.Marshal(unsigned)
	if err != nil {
		fmt.Println("Failed to marshal unsigned:", err)
		return "", err
	}
	hash := sha3.Sum512(payload)

	resp, err := client.Test(a.ctx, &p.SignedBaseEvent{
		Id:      hash[:],
		Payload: payload,
	})
	if err != nil {
		fmt.Println("Failed to call Test:", err)
		return "", err
	}

	return resp.GetMessage(), nil
}
