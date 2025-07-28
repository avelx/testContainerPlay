package main

import (
	"context"
	"fmt"
	"time"

	//"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func startContainer(ctx *context.Context) *testcontainers.Container {

	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, err := testcontainers.GenericContainer(*ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	return &redisC
	//testcontainers.CleanupContainer(t, redisC)
	//require.NoError(t, err)
}

func stopContainer(ctx *context.Context, container *testcontainers.Container) {
	duration := 1 * time.Second
	(*container).Stop(*ctx, &duration)
	//container.CleanupContainer(t, redisC)
	//require.NoError(t, err)
}

func main() {
	ctx := context.Background()
	fmt.Printf("Attempt to start test container ...\n")
	container := startContainer(&ctx)
	time.Sleep(10 * time.Second)
	stopContainer(&ctx, container)

}
