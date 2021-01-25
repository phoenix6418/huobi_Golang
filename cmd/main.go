package main

import (
	"github.com/huobirdcenter/huobi_golang/cmd/accountclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/commonclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/crossmarginclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/etfclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/isolatedmarginclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/marketclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/orderclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/walletclientexample"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/perflogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"log"
)

func main() {
	//runAll()
	runApiTest()
}

func runApiTest() {
	// common client with http
	client := new(client.MarketClient).Init(config.Host)
	resp, err := client.GetLatestTrade("btcusdt")
	if err != nil {
		log.Printf("Get Latest Trade Err: %v\n", err)
		return
	}

	for _, dt := range resp.Data {
		log.Printf("data: %+v\n", dt)
	}
}

// Run all examples
func runAll() {
	//commonclientexample.RunAllExamples()
	//accountclientexample.RunAllExamples()
	//orderclientexample.RunAllExamples()
	//algoorderclientexample.RunAllExamples()
	//marketclientexample.RunAllExamples()
	//isolatedmarginclientexample.RunAllExamples()
	//crossmarginclientexample.RunAllExamples()
	//walletclientexample.RunAllExamples()
	//subuserclientexample.RunAllExamples()
	//stablecoinclientexample.RunAllExamples()
	//etfclientexample.RunAllExamples()
	//marketwebsocketclientexample.RunAllExamples()
	//accountwebsocketclientexample.RunAllExamples()
	//orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}
