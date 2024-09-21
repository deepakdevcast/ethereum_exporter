package rpc

import (
	"errors"
	ethHttpRpc "ethereum_exporter/internal/eth/http_rpc"
	"ethereum_exporter/internal/promclient"
	"fmt"
	"log"
	"strconv"
)

var ethGasPriceBody = &ethHttpRpc.RpcRequestBody{
	Jsonrpc: "2.0",
	Method:  "eth_gasPrice",
	Id:      0,
	Params:  []interface{}{},
}

func getEthGasPrice(ethRpcClient *ethHttpRpc.EthHttpRpcClient) (float64, error) {
	ethGasPriceRpcMethod := ethHttpRpc.NewEthRpcMethod(ethRpcClient, ethGasPriceBody)

	result := ethGasPriceRpcMethod.Get()

	if err := result.Err; err != nil {
		return 0, errors.Join(ErrJsonMarshal, err)
	}
	GasPrice, ok := result.Result.(string)
	if !ok {
		return 0, ErrStringConvert
	}
	gasPriceGwi, err := strconv.ParseInt(GasPrice, 0, 64)
	if err != nil {
		return 0, ErrNumberConvert
	}
	return float64(gasPriceGwi), nil
}

func EthGasPriceMetric(ethRpcClient *ethHttpRpc.EthHttpRpcClient, labels map[string]string) {
	GasPrice, err := getEthGasPrice(ethRpcClient)

	if err != nil {
		fmt.Printf("Get GasPrice request failed, %v, %v \n", labels["name"], err)
		return
	}
	promclient.EthGasPriceGauge.WithLabelValues(labels["name"]).Set(GasPrice)
	log.Println("GasPrice of", labels["name"], "is", GasPrice)
}
