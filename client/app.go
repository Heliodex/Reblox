package main

import (
	p "Reblox/server/proto"
	"Reblox/shared/keys"
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/sha3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

const sec = "e63mgel6p15f3dn8uxv4ni4ldylmrtyt0fubkxvl0z3q20y80w"

func (a *App) ContactServer(message string) (string, error) {
	conn, err := grpc.Dial("localhost:2006", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := p.NewEventServiceClient(conn)

	secb, err := keys.Decode(sec)
	if err != nil {
		fmt.Println("Failed to decode private key:", err)
		return "", err
	}
	privateKey, err := keys.PrivateKey(secb)
	if err != nil {
		fmt.Println("Failed to get private key:", err)
		return "", err
	}
	pubb := privateKey.PublicKey().Bytes()

	unsigned := &p.UnsignedBaseEvent{
		Pubkey:  pubb,
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
