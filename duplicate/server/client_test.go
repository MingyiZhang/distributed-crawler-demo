package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"distributed-crawler-demo/config"
	"distributed-crawler-demo/rpchelper"
	"github.com/testcontainers/testcontainers-go"
)

func TestService(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:6.0.3-alpine",
		ExposedPorts: []string{"6379"},
	}

	redisCtn, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}
	defer redisCtn.Terminate(ctx)

	ip, err := redisCtn.Host(ctx)
	if err != nil {
		t.Error(err)
	}
	port, err := redisCtn.MappedPort(ctx, "6379")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ip, port.Port())

	const host = ":6300"
	go serveRpc(host, ip+":"+port.Port())
	time.Sleep(time.Second)

	client, err := rpchelper.NewClient(host)
	if err != nil {
		t.Error(err)
	}

	result := 0
	err = client.Call(config.DuplicateServiceRpc, "test_url", &result)
	if err != nil || result != 0 {
		t.Errorf("result: %d; err: %v", result, err)
	}

	err = client.Call(config.DuplicateServiceRpc, "test_url", &result)
	if err != nil || result != 1 {
		t.Errorf("result: %d; err: %v", result, err)
	}

}
