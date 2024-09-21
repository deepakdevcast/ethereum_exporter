package rpc

import (
	"errors"
	ethHttpRpc "ethereum_exporter/internal/eth/http_rpc"
	"ethereum_exporter/internal/promclient"
	"fmt"
	"log"
)

var ethSyncingBody = &ethHttpRpc.RpcRequestBody{
	Jsonrpc: "2.0",
	Method:  "eth_syncing",
	Id:      0,
	Params:  []interface{}{},
}

func getEthSyncing(ethRpcClient *ethHttpRpc.EthHttpRpcClient) (bool, error) {
	ethSyncingRpcMethod := ethHttpRpc.NewEthRpcMethod(ethRpcClient, ethSyncingBody)

	result := ethSyncingRpcMethod.Get()

	if err := result.Err; err != nil {
		return false, errors.Join(ErrJsonMarshal, err)
	}
	Syncing, ok := result.Result.(bool)
	if !ok {
		return false, ErrStringConvert
	}

	return Syncing, nil
}

func EthSyncingMetric(ethRpcClient *ethHttpRpc.EthHttpRpcClient, labels map[string]string) {
	Syncing, err := getEthSyncing(ethRpcClient)

	if err != nil {
		fmt.Printf("Get Syncing request failed, %v, %v \n", labels["name"], err)
		return
	}
	if Syncing {
		promclient.EthSyncingGauge.WithLabelValues(labels["name"]).Set(1)
	} else {
		promclient.EthSyncingGauge.WithLabelValues(labels["name"]).Set(0)
	}
	log.Println("Syncing of", labels["name"], "is", Syncing)
}
