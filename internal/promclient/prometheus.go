package promclient

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

var EthBlockHeightGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "eth_block_height",
	Help: "Latest ethereum block number.",
}, []string{"name"})

func init() {
	prometheus.MustRegister(EthBlockHeightGauge)
	prometheus.Unregister(collectors.NewGoCollector())
}
