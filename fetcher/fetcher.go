package fetcher

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "time"

  "golang.org/x/net/html/charset"
  "golang.org/x/text/encoding"
  "golang.org/x/text/encoding/unicode"
  "golang.org/x/text/transform"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
  //<- rateLimiter
  res, err := http.Get(url)
  if err != nil {
    panic(err)
  }
  defer  res.Body.Close()

  if res.StatusCode != http.StatusOK {
    return nil, fmt.Errorf("wrong status code: %d", res.StatusCode)
  }

  br := bufio.NewReader(res.Body)
  e := determineEncoding(br)
  utf8Reader := transform.NewReader(br, e.NewDecoder())

  return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
  bytes, err := r.Peek(1024)
  if err != nil {
    log.Printf("Fetcher error: %v", err)
    return unicode.UTF8
  }
  e, _, _ := charset.DetermineEncoding(bytes, "")
  return e
}