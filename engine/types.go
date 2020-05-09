package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Parser interface {
  Parse(contents []byte, url string) ParseResult
  Serialize() (name string, args interface{})
}

// Request is the instant that needs to be parsed.
type Request struct {
  // Url is the url to parse.
  Url    string
  // Parser is used for parse the Request's Url.
  Parser Parser
}

// ParseResult is the instant that contains the result after scraping.
// It contains new Requests from previous Request and
// the Items that are to save.
type ParseResult struct {
  Requests []Request
  Items    []Item
}

// Item is the instant that need to be saved to database.
type Item struct {
  Url     string
  Id      string
  // Payload is the model of the data
  Payload interface{}
}

type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
  return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
  return "NilParser", nil
}

type FuncParser struct {
  parser ParserFunc
  name string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
  return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
  return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
  return &FuncParser{
    parser: p,
    name: name,
  }
}