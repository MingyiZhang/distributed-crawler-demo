package config

const (
	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	ParseCounty = "ParseCounty"

	// Service Endpoints
	ItemSaverRpc        = "ItemSaverService.Save"
	CrawlServiceRpc     = "CrawlService.Process"
	DuplicateServiceRpc = "Service.Exists"

	// Rate Limit
	QPS = 20
)
