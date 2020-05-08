package main

import (
  "flag"
  "log"
  "net/rpc"
  "strings"

  "distributed-crawler-demo/engine"
  saver "distributed-crawler-demo/persist/client"
  "distributed-crawler-demo/rpchelper"
  "distributed-crawler-demo/scheduler"
  "distributed-crawler-demo/webs/mockweb/parser"
  worker "distributed-crawler-demo/worker/client"
)

const cityListUrl = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

var (
  itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
  workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
  flag.Parse()
  itemChan, err := saver.ItemSaver(*itemSaverHost)
  if err != nil {
    panic(err)
  }

  pool := createClientPool(strings.Split(*workerHosts, ","))
  processor := worker.CreateProcessor(pool)

  e := engine.ConcurrentEngine{
    Scheduler:        &scheduler.QueuedScheduler{},
    WorkerCount:      100,
    ItemChan:         itemChan,
    RequestProcessor: processor,
  }

  e.Run(engine.Request{
    Url:    cityListUrl,
    Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
  })
}

func createClientPool(hosts []string) chan *rpc.Client {
  var clients []*rpc.Client
  for _, h := range hosts {
    client, err := rpchelper.NewClient(h)
    if err == nil {
      clients = append(clients, client)
    } else {
      log.Printf("error connecting to %s: %v", h, err)
    }
  }

  out := make(chan *rpc.Client)
  go func() {
    for {
      for _, client := range clients {
        out <- client
      }
    }
  }()
  return out
}