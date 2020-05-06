package config

const (
  // Parser names
  ParseCity     = "ParseCity"
  ParseCityList = "ParseCityList"
  ParseProfile  = "ParseProfile"
  NilParser     = "NilParser"

  // Elasticsearch Index
  ElasticIndex = "dating_profile"

  // Service Endpoints
  ItemSaverRpc    = "ItemSaverService.Save"
  CrawlServiceRpc = "CrawlService.Process"

  // Rate Limit
  QPS = 20
)
