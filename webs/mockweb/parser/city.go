package parser

import (
  "regexp"

  "distributed-crawler-demo/config"
  "distributed-crawler-demo/engine"
)

var (
  profileRe = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
  cityUrlRe = regexp.MustCompile(`href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
  bytes := profileRe.FindAllSubmatch(contents, -1)
  result := engine.ParseResult{}
  for _, b := range bytes {
    result.Requests = append(result.Requests, engine.Request{
      Url:    string(b[1]),
      Parser: NewProfileParser(string(b[2])),
    })
  }

  bytes = cityUrlRe.FindAllSubmatch(contents, -1)
  for _, b := range bytes {
    result.Requests = append(result.Requests, engine.Request{
      Url:    string(b[1]),
      Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
    })
  }
  return result
}
