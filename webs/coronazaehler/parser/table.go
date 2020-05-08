package parser

import (
  "log"
  "regexp"
  "strconv"
  "strings"

  "distributed-crawler-demo/engine"
  "distributed-crawler-demo/webs/coronazaehler/model"
  "github.com/PuerkitoBio/goquery"
)

var attrRe = regexp.MustCompile(`(#[^']+)`)

func ParseCounty(content []byte, _ string) engine.ParseResult {
  doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
  if err != nil {
    log.Printf("error fetching counties: %v", err)
    return engine.ParseResult{}
  }

  result := engine.ParseResult{}

  doc.Find("#Deutschland").Each(func(_ int, stateTable *goquery.Selection) {
    stateTable.Find("tr").Each(func(_ int, stateRow *goquery.Selection) {
      if attr, exists := stateRow.Attr("onclick"); exists {
        id := string(attrRe.Find([]byte(attr)))
        state := stateRow.Find("td").First().Text()
        doc.Find(id).Find("tbody").
            Each(func(_ int, countyTable *goquery.Selection) {
              countyTable.Find("tr").Each(func(_ int, countyRow *goquery.Selection) {
                item := engine.Item{}
                county := model.County{
                  State: state,
                }
                info := countyRow.Find("td")
                countyName := info.Eq(0).Text()
                item.Id = countyName
                county.Name = countyName
                county.I100K = strings.Replace(info.Eq(1).Text(), ",", ".", -1)
                county.Dead, _ = numberConverter(info.Eq(2).Text())
                county.Infected, _ = numberConverter(info.Eq(3).Text())
                county.Recovered, _ = numberConverter(info.Eq(4).Text())
                item.Payload = county
                result.Items = append(result.Items, item)
              })
            })
      }
    })
  })
  return result
}

// convert number from german format to international format
func numberConverter(s string) (int, error) {
  ss := strings.Trim(strings.ReplaceAll(s, ".", ""), "\n ")
  return strconv.Atoi(ss)
}
