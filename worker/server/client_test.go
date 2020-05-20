package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"distributed-crawler-demo/config"
	"distributed-crawler-demo/rpchelper"
	"distributed-crawler-demo/worker"
)

func TestCrawlService(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(w, "Hi there!")
		}))
	defer ts.Close()

	const host = ":9000"
	go rpchelper.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpchelper.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		//Url: "http://localhost:8080/mock/album.zhenai.com/u/6721425675858866615",
		Url: ts.URL,
		Parser: worker.SerializedParser{
			//Name: config.ParseProfile,
			Name: config.NilParser,
			//Args: "寂寞成影莓哒",
			Args: "",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

}
