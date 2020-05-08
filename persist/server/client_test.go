package main

import (
  "testing"
  "time"

  "distributed-crawler-demo/config"
  "distributed-crawler-demo/engine"
  "distributed-crawler-demo/model"
  "distributed-crawler-demo/rpchelper"
)

func TestItemSaver(t *testing.T) {
  const host = ":1234"
  go serveRpc(host, "localhost:9200", "test1")
  time.Sleep(time.Second)

  client, err := rpchelper.NewClient(host)
  if err != nil {
    panic(err)
  }

  item := engine.Item{
    Url: "test_url",
    Id: "test_id",
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
