package main

import (
  "net/http"

  "distributed-crawler-demo/webs/mockweb/frontend/controller"
)

func main() {
  http.Handle("/", http.FileServer(
    http.Dir("webs/mockweb/frontend/view")))
  http.Handle(
    "/search",
    controller.CreateSearchResultHandler(
      "webs/mockweb/frontend/view/template.html"))
  err := http.ListenAndServe(":8888", nil)
  if err != nil {
    panic(err)
  }
}
