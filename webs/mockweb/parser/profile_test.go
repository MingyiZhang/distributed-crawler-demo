package parser

import (
  "io/ioutil"
  "testing"

  "distributed-crawler-demo/engine"
  "distributed-crawler-demo/model"
)

func TestParseProfile(t *testing.T) {
  contents, err := ioutil.ReadFile("profile_test_data.html")
  if err != nil {
    panic(err)
  }
  result := parseProfile(contents, "http://album.zhenai.com/u/108906739", "寂寞成影萌宝")
  if len(result.Items) != 1 {
    t.Errorf("Items should contain 1 element; but was %v", result.Items)
  }

  actual := result.Items[0]
  expected := engine.Item {
    Url: "http://album.zhenai.com/u/108906739",
    Id: "108906739",
    Payload: model.Profile{
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
  if actual != expected {
    t.Errorf("expected %v; but was %v", expected, actual)
  }
}
