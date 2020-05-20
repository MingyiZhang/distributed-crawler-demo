package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"distributed-crawler-demo/config"
	"distributed-crawler-demo/engine"
	"distributed-crawler-demo/model"
	"distributed-crawler-demo/rpchelper"
	"github.com/testcontainers/testcontainers-go"
)

func TestItemSaver(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "docker.elastic.co/elasticsearch/elasticsearch:7.6.2",
		ExposedPorts: []string{"9200"},
		Env: map[string]string{
			"discovery.type": "single-node",
		},
	}
	elasticCtn, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}
	defer elasticCtn.Terminate(ctx)

	ip, err := elasticCtn.Host(ctx)
	if err != nil {
		t.Error(err)
	}
	port, err := elasticCtn.MappedPort(ctx, "9200")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ip, port.Port())

	const host = ":1234"
	url := fmt.Sprintf("http://%s:%s", ip, port.Port())

	_, err = http.Get(url)
	for err != nil {
		time.Sleep(2 * time.Second)
		_, err = http.Get(url)
	}

	go serveRpc(host, url, "test1")
	time.Sleep(time.Second)

	client, err := rpchelper.NewClient(host)
	if err != nil {
		t.Error(err)
	}

	item := engine.Item{
		Url: "test_url",
		Id:  "test_id",
		Payload: model.Profile{
			Name:          "寂寞成影萌宝",
			Gender:        "女",
			Age:           83,
			Height:        105,
			Weight:        137,
			Income:        "财务自由",
			Marriage:      "离异",
			Education:     "初中",
			Occupation:    "金融",
			Residence:     "南京市",
			Constellation: "狮子座",
			House:         "无房",
			Car:           "无车",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}

}
