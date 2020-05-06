package main

import (
  "fmt"
  "testing"
  "time"

  "distributed-crawler-demo/config"
  "distributed-crawler-demo/rpchelper"
  "distributed-crawler-demo/worker"
)

func TestCrawlService(t *testing.T) {
  const host = ":9000"
  go rpchelper.ServeRpc(host, worker.CrawlService{})
  time.Sleep(time.Second)

  client, err := rpchelper.NewClient(host)
  if err != nil {
    panic(err)
  }

  req := worker.Request{
    Url: "http://localhost:8080/mock/album.zhenai.com/u/6721425675858866615",
    Parser: worker.SerializedParser{
      Name: config.ParseProfile,
      Args: "寂寞成影莓哒",
    },
  }
  var result worker.ParseResult
  err = client.Call(config.CrawlServiceRpc, req, &result)
  if err != nil {
    t.Error(err)
  } else {
    fmt.Println(result)
  }

}