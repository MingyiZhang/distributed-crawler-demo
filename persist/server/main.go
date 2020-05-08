package main

import (
  "flag"
  "fmt"
  "log"

  "distributed-crawler-demo/persist"
  "distributed-crawler-demo/rpchelper"
  "github.com/olivere/elastic/v7"
)

var (
	port = flag.Int("port", 0, "the port to listening on")
	elasticUrl = flag.String("elastic_url", "localhost:9200", "the url of elasticsearch")
	elasticIndex = flag.String("elastic_index", "index", "the elastic index")
)

func main() {
  flag.Parse()
  if *port == 0 {
    fmt.Println("must specify a port")
  }
  log.Fatal(serveRpc(
    fmt.Sprintf(":%d", *port),
    *elasticUrl,
    *elasticIndex))
}

func serveRpc(host, url, index string) error {
  client, err := elastic.NewClient(
    elastic.SetURL(url),
    elastic.SetSniff(false))
  if err != nil {
    return err
  }
  return rpchelper.ServeRpc(host,
    &persist.ItemSaverService{
      Client: client,
      Index:  index,
    })
}
