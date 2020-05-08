package controller

import (
  "context"
  "net/http"
  "reflect"
  "regexp"
  "strconv"
  "strings"

  "distributed-crawler-demo/engine"
  "distributed-crawler-demo/frontend/model"
  "distributed-crawler-demo/frontend/view"
  "github.com/olivere/elastic/v7"
)

type SearchResultHandler struct {
  view   view.SearchResultView
  client *elastic.Client
  indices []string
}

func CreateSearchResultHandler(url, template string, indices []string) SearchResultHandler {
  client, err := elastic.NewClient(
    elastic.SetURL(url), elastic.SetSniff(false))
  if err != nil {
    panic(err)
  }
  return SearchResultHandler{
    view:    view.CreateSearchResultView(template),
    client:  client,
    indices: indices,
  }
}

// localhost:8888/search?q=...&from=...
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  q := strings.TrimSpace(req.FormValue("q"))
  from, err := strconv.Atoi(
    req.FormValue("from"))
  if err != nil {
    from = 0
  }

  page, err := h.getSearchResult(q, h.indices, from)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }
  err = h.view.Render(w, page)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }
}

func (h SearchResultHandler) getSearchResult(q string, indices []string, from int) (model.SearchResult, error) {
  var result model.SearchResult
  res, err := h.client.
    Search(indices...).
    Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
    From(from).
    Do(context.Background())
  if err != nil {
    return result, err
  }

  result.Hits = res.TotalHits()
  result.Start = from
  result.Query = q
  result.Items = res.Each(reflect.TypeOf(engine.Item{}))
  result.PrevFrom = result.Start - len(result.Items)
  result.NextFrom = result.Start + len(result.Items)
  return result, nil
}

func rewriteQueryString(q string) string {
  re := regexp.MustCompile(`([A-Z][a-z]*):`)
  return re.ReplaceAllString(q, "Payload.$1:")
}
