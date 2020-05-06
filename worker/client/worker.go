package client

import (
  "fmt"

  "distributed-crawler-demo/config"
  "distributed-crawler-demo/engine"
  "distributed-crawler-demo/rpchelper"
  "distributed-crawler-demo/worker"
)

func CreateProcessor() (engine.Processor, error) {
  client, err := rpchelper.NewClient(
    fmt.Sprintf(":%d", config.WorkerPort0))
  if err != nil {
    return nil, err
  }

  return func(req engine.Request) (engine.ParseResult, error) {
    sReq := worker.SerializeRequest(req)

    var sResult worker.ParseResult
    err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
    if err != nil {
      return engine.ParseResult{}, nil
    }

    return worker.DeserializeResult(sResult), nil
  }, nil
}