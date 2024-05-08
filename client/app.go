package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"Reblox/server/proto"
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

func (a *App) ContactServer(message string) (string, error) {
	conn, err := grpc.Dial("localhost:2006", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := proto.NewTestServiceClient(conn)
	resp, err := client.Test(context.Background(), &proto.TestMessage{Message: message})
	if err != nil {
		fmt.Println("Failed to call Test:", err)
		return "", err
	}

	return resp.GetMessage(), nil
}
