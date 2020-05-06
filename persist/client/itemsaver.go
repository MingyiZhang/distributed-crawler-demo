package client

import (
  "log"

  "distributed-crawler-demo/config"
  "distributed-crawler-demo/engine"
  "distributed-crawler-demo/rpchelper"
)

func ItemSaver(host string) (chan engine.Item, error) {
  client, err := rpchelper.NewClient(host)
  if err != nil {
    return nil, err
  }
  out := make(chan engine.Item)
  go func() {
    itemCount := 0
    for {
      item := <-out
      log.Printf("Item Saver: got item #%d: %v", itemCount, item)
      itemCount++

      result := ""
      err := client.Call(config.ItemSaverRpc, item, &result)
      if err != nil {
        log.Printf("Item saver: error saving item %v: %v", item, err)
      }
    }
  }()
  return out, nil
}