package main

import (
  "flag"
  "fmt"
  "net/http"

  "distributed-crawler-demo/webs/mockweb/frontend/controller"
)

var (
	url = flag.String("url", "http://localhost:9200", "elasticsearch url")
	port = flag.Int("port", 8888, "the port to listen on")
)

func main() {
  flag.Parse()

  http.Handle("/", http.FileServer(
    http.Dir("view")))
  http.Handle(
    "/search",
    controller.CreateSearchResultHandler(
      *url,
      "view/template.html"))
  err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
  if err != nil {
    panic(err)
  }
}
