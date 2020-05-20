package main

import (
	"flag"
	"fmt"
	"log"

	"distributed-crawler-demo/duplicate"
	"distributed-crawler-demo/rpchelper"
	"github.com/go-redis/redis/v7"
)

var (
	port     = flag.Int("port", 0, "the port of RPC server")
	redisUrl = flag.String("redis_url", "localhost:6379", "the url of redis")
)

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
	}
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", *port),
		*redisUrl))
}

func serveRpc(host, url string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	return rpchelper.ServeRpc(host,
		&duplicate.Service{Client: client})
}
