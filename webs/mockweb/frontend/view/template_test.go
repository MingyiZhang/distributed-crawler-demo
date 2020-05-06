package view

import (
  "os"
  "testing"

  "distributed-crawler-demo/engine"
  common "distributed-crawler-demo/model"
  "distributed-crawler-demo/webs/mockweb/frontend/model"
)

func TestSearchResultView_Render(t *testing.T) {
  view := CreateSearchResultView("template.html")
  out, err := os.Create("template.test.html")
  page := model.SearchResult{}
  page.Hits = 123
  item := engine.Item{
    Url: "test_url",
    Id:  "test_id",
    Payload: common.Profile{
      Name:          "寂寞成影萌宝",
      Gender:        "女",
      Age:           83,
      Height:        105,
      Weight:        137,
      Income:        "财务自由",
      Marriage:      "离异",
      Education:     "初中",
      Occupation:    "金融",
      Residence:     "南京市",
      Constellation: "狮子座",
      House:         "无房",
      Car:           "无车",
    },
  }
  for i := 0; i < 10; i++ {
    page.Items = append(page.Items, item)
  }

  err = view.Render(out, page)
  if err != nil {
    panic(err)
  }
}
