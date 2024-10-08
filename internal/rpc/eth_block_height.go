package rpc

import (
	"errors"
	ethHttpRpc "ethereum_exporter/internal/eth/http_rpc"
	"ethereum_exporter/internal/promclient"
	"fmt"
	"log"
	"strconv"
)

var ethBlockHeightBody = &ethHttpRpc.RpcRequestBody{
	Jsonrpc: "2.0",
	Method:  "eth_blockNumber",
	Id:      83,
	Params:  []interface{}{},
}

func getEthBlockHeight(ethRpcClient *ethHttpRpc.EthHttpRpcClient) (float64, error) {

	ethBlockHeightRpcMethod := ethHttpRpc.NewEthRpcMethod(ethRpcClient, ethBlockHeightBody)

	result := ethBlockHeightRpcMethod.Get()

	if err := result.Err; err != nil {
		return 0, errors.Join(ErrJsonMarshal, err)
	}

	strResult, ok := result.Result.(string)
	if !ok {
		return 0, ErrStringConvert
	}

	blockHeight, err := strconv.ParseInt(strResult, 0, 64)
	if err != nil {
		return 0, errors.Join(ErrNumberConvert, err)
	}

	return float64(blockHeight), nil
}

func EthBlockHeightMetric(ethRpcClient *ethHttpRpc.EthHttpRpcClient, labels map[string]string) {
	blockHeight, err := getEthBlockHeight(ethRpcClient)

	if err != nil {
		fmt.Printf("Get BlockHeight request failed, %v, %v \n", labels["name"], err)
		return
	}

	promclient.EthBlockHeightGauge.WithLabelValues(labels["name"]).Set(blockHeight)
	log.Println("BlockHeight of", labels["name"], "is", blockHeight)
}
