package main

import (
	"ethereum_exporter/config"
	ethHttpRpc "ethereum_exporter/internal/eth/http_rpc"
	"ethereum_exporter/internal/rpc"
	"log"
	"net/http"
	"net/url"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		for _, clientInfo := range config.ClientInfo {
			url, err := url.Parse(clientInfo.Url)
			if err != nil {
				log.Panicf("Incorrect Url: %s\n", url)
			}
			labels := map[string]string{"name": clientInfo.Client}
			ethHttpRpcClient, err := ethHttpRpc.NewEthHttpRpcClient(ethHttpRpc.SetHttpRequest(url))
			if err != nil {
				log.Panicf("Failed to created Eth Rpc Client %s\n", err)
			}
			rpc.EthBlockHeightMetric(ethHttpRpcClient, labels)
			rpc.EthAccountsMetric(ethHttpRpcClient, labels)
		}
		promhttp.Handler().ServeHTTP(w, r)
	})
	log.Println("Metrics Server starting on Port", config.MetricsPort)
	http.ListenAndServe(config.MetricsPort, nil)
}
