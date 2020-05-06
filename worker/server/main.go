package main

import (
  "fmt"
  "log"

  "distributed-crawler-demo/config"
  "distributed-crawler-demo/rpchelper"
  "distributed-crawler-demo/worker"
)

func main() {
  log.Fatal(rpchelper.ServeRpc(
    fmt.Sprintf(":%d", config.WorkerPort0),
    worker.CrawlService{}))
}
