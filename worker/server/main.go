package main

import (
  "flag"
  "fmt"
  "log"

  "distributed-crawler-demo/rpchelper"
  "distributed-crawler-demo/worker"
)

var port = flag.Int("port", 0, "the port to listen on")

func main() {
  flag.Parse()
  if *port == 0 {
    fmt.Println("must specify a port")
    return
  }
  log.Fatal(rpchelper.ServeRpc(
    fmt.Sprintf(":%d", *port),
    worker.CrawlService{}))
}
