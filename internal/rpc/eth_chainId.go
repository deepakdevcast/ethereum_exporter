package rpc

import (
	"errors"
	ethHttpRpc "ethereum_exporter/internal/eth/http_rpc"
	"ethereum_exporter/internal/promclient"
	"fmt"
	"log"
)

var ethChainIdBody = &ethHttpRpc.RpcRequestBody{
	Jsonrpc: "2.0",
	Method:  "eth_chainId",
	Id:      0,
	Params:  []interface{}{},
}

func getEthChainId(ethRpcClient *ethHttpRpc.EthHttpRpcClient) (string, error) {
	ethChainIdRpcMethod := ethHttpRpc.NewEthRpcMethod(ethRpcClient, ethChainIdBody)

	result := ethChainIdRpcMethod.Get()

	if err := result.Err; err != nil {
		return "", errors.Join(ErrJsonMarshal, err)
	}
	chainId, ok := result.Result.(string)
	if !ok {
		return "", ErrStringConvert
	}
	return chainId, nil
}

func EthChainIdMetric(ethRpcClient *ethHttpRpc.EthHttpRpcClient, labels map[string]string) {
	ChainId, err := getEthChainId(ethRpcClient)

	if err != nil {
		fmt.Printf("Get ChainId request failed, %v, %v \n", labels["name"], err)
		return
	}
	promclient.EthChainIdGauge.WithLabelValues(labels["name"], ChainId).Set(1)
	log.Println("ChainId of", labels["name"], "is", ChainId)
}
