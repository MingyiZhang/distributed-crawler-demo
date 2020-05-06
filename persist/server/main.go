package main

import (
  "fmt"
  "log"

  "distributed-crawler-demo/config"
  "distributed-crawler-demo/persist"
  "distributed-crawler-demo/rpchelper"
  "github.com/olivere/elastic/v7"
)

func main() {
  log.Fatal(serveRpc(
    fmt.Sprintf(":%d", config.ItemSaverPort),
    config.ElasticIndex))
}

func serveRpc(host, index string) error {
  client, err := elastic.NewClient(elastic.SetSniff(false))
  if err != nil {
    return err
  }
  return rpchelper.ServeRpc(host,
    &persist.ItemSaverService{
      Client: client,
      Index:  index,
    })
}
