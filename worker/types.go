package worker

import (
  "errors"
  "fmt"
  "log"

  "distributed-crawler-demo/config"
  "distributed-crawler-demo/engine"
  parser2 "distributed-crawler-demo/webs/coronazaehler/parser"
  "distributed-crawler-demo/webs/mockweb/parser"
)


type SerializedParser struct {
  Name string
  Args interface{}
}

type Request struct {
  Url    string
  Parser SerializedParser
}

type ParseResult struct {
  Items    []engine.Item
  Requests []Request
}

func SerializeRequest(r engine.Request) Request {
  name, args := r.Parser.Serialize()
  return Request{
    Url:    r.Url,
    Parser: SerializedParser{
      Name: name,
      Args: args,
    },
  }
}

func SerializeResult(r engine.ParseResult) ParseResult {
  result := ParseResult {
    Items: r.Items,
  }
  for _, req := range r.Requests {
    result.Requests = append(result.Requests, SerializeRequest(req))
  }
  return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
  ps, err := deserializeParser(r.Parser)
  if err != nil {
    return engine.Request{}, err
  }
  return engine.Request {
    Url: r.Url,
    Parser: ps,
  }, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
  result := engine.ParseResult{
    Items: r.Items,
  }
  for _, req := range r.Requests {
    engineReq, err := DeserializeRequest(req)
    if err != nil {
      log.Printf("error deserializing request: %v", err)
      continue
    }
    result.Requests = append(result.Requests, engineReq)
  }
  return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
  switch p.Name {
  case config.ParseCityList:
    return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
  case config.ParseCity:
    return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
  case config.NilParser:
    return engine.NilParser{}, nil
  case config.ParseProfile:
    if userName, ok := p.Args.(string); ok {
      return parser.NewProfileParser(userName), nil
    } else {
      return nil, fmt.Errorf("invalid arg: %v", p.Args)
    }
  case config.ParseCounty:
    return engine.NewFuncParser(parser2.ParseCounty, config.ParseCounty), nil
  default:
    return nil, errors.New("unknown parser name")

  }
}


