package client

import (
  "net/rpc"

  "distributed-crawler-demo/config"
  "distributed-crawler-demo/engine"
  "distributed-crawler-demo/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
  return func(req engine.Request) (engine.ParseResult, error) {
    sReq := worker.SerializeRequest(req)

    var sResult worker.ParseResult
    c := <-clientChan
    err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
    if err != nil {
      return engine.ParseResult{}, nil
    }

    return worker.DeserializeResult(sResult), nil
  }
}