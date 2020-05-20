package client

import (
	"net/rpc"

	"distributed-crawler-demo/config"
	"distributed-crawler-demo/engine"
)

func CreateChecker(c *rpc.Client) engine.Checker {
	return func(url string) (bool, error) {
		result := false
		err := c.Call(config.DuplicateServiceRpc, url, &result)
		if err != nil {
			return false, err
		}
		return result, nil
	}
}
