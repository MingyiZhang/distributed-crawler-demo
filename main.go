package main

import (
  "fmt"

  "distributed-crawler-demo/config"
  "distributed-crawler-demo/engine"
  saver "distributed-crawler-demo/persist/client"
  "distributed-crawler-demo/scheduler"
  "distributed-crawler-demo/webs/mockweb/parser"
  worker "distributed-crawler-demo/worker/client"
)

const cityListUrl = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main() {
  itemChan, err := saver.ItemSaver(
    fmt.Sprintf(":%d", config.ItemSaverPort))
  if err != nil {
    panic(err)
  }

  processor, err := worker.CreateProcessor()
  if err != nil {
    panic(err)
  }
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
