package promclient

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

var EthBlockHeightGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "eth_block_height",
	Help: "Latest ethereum block number.",
}, []string{"name"})

var EthAccountsGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "eth_accounts",
	Help: "Gives accounts own by client as label.",
}, []string{"name", "account"})

var EthChainIdGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "eth_chainId",
	Help: "Gives current network chainId",
}, []string{"name", "chainId"})

var EthGasPriceGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "eth_gasPrice",
	Help: "Gives current price per gas in wei",
}, []string{"name"})

var EthSyncingGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "eth_syncing",
	Help: "Gives Syncing Status",
}, []string{"name"})

func init() {
	prometheus.MustRegister(
		EthBlockHeightGauge,
		EthAccountsGauge,
		EthChainIdGauge,
		EthGasPriceGauge,
		EthSyncingGauge,
	)
	prometheus.Unregister(collectors.NewGoCollector())
}
