package parser

import (
  "io/ioutil"
  "log"
  "net/http"
  "testing"

  "distributed-crawler-demo/webs/coronazaehler/model"
)


func TestParseTable(t *testing.T) {
  res, err := http.Get("https://www.coronazaehler.de")
  if err != nil {
    panic(err)
  }
  defer res.Body.Close()

  if res.StatusCode != http.StatusOK {
    log.Fatalf("wrong status code: %d", res.StatusCode)
  }

  content, err := ioutil.ReadAll(res.Body)
  if err != nil {
    panic(err)
  }
  result := ParseCounty(content, "")
  item0 := result.Items[0]
  county := item0.Id
  state := item0.Payload.(model.County).State
  if county != "Kreis Esslingen" {
    t.Errorf("got first item id %s; expected %s", item0.Id, "Kreis Esslingen")
  }
  if state != "Baden-Württemberg" {
    t.Errorf("got first state %s; expected %s", state, "Baden-Württemberg")
  }
  for _, item := range result.Items {
    name := item.Payload.(model.County).Name
    if item.Id != name {
      t.Errorf("got county id %s different from county name %s; expected equal", item.Id, name)
    }
  }
  if len(result.Requests) != 0 {
    t.Errorf("got %d requests; expected %d", len(result.Requests), 0)
  }
}

