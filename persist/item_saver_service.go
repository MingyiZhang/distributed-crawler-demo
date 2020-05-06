package persist

import (
  "context"
  "log"

  "distributed-crawler-demo/engine"
  "github.com/olivere/elastic/v7"
)

type ItemSaverService struct {
  Client *elastic.Client
  Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
  err := save(s.Client, s.Index, item)
  log.Printf("Item %v saved.", item)
  if err == nil {
    *result = "ok"
  } else {
    log.Printf("Error saving item %v: %v", item, err)
  }
  return err
}

func save(client *elastic.Client, index string, item engine.Item) error {
  indexService := client.Index().
    Index(index).
    BodyJson(item)
  if item.Id != "" {
    indexService.Id(item.Id)
  }
  _, err := indexService.Do(context.Background())
  if err != nil {
    return err
  }

  return nil
}