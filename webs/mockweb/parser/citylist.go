package parser

import (
  "regexp"

  "distributed-crawler-demo/config"
  "distributed-crawler-demo/engine"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
  re := regexp.MustCompile(cityListRe)
  bytes := re.FindAllSubmatch(contents, -1)
  result := engine.ParseResult{}
  for _, b := range bytes {
    result.Requests = append(result.Requests, engine.Request{
      Url:    string(b[1]),
      Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
    })
  }
  return result
}
