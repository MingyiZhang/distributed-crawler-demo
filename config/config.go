package config

const (
  // Parser names
  ParseCity     = "ParseCity"
  ParseCityList = "ParseCityList"
  ParseProfile  = "ParseProfile"
  NilParser     = "NilParser"

  // Service ports
  ItemSaverPort = 1234
  WorkerPort0   = 9000

  // Elasticsearch Index
  ElasticIndex = "dating_profile"

  // Service Endpoints
  ItemSaverRpc    = "ItemSaverService.Save"
  CrawlServiceRpc = "CrawlService.Process"
)
