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

func init() {
	prometheus.MustRegister(EthBlockHeightGauge, EthAccountsGauge)
	prometheus.Unregister(collectors.NewGoCollector())
}
