package rpc

import (
	"errors"
	ethHttpRpc "ethereum_exporter/internal/eth/http_rpc"
	"ethereum_exporter/internal/promclient"
	"fmt"
	"log"
	"strconv"
)

var ethRpcMethod = &ethHttpRpc.RpcRequestBody{
	Jsonrpc: "2.0",
	Method:  "eth_blockNumber",
	Id:      83,
	Params:  []interface{}{},
}

func getEthBlockHeight(ethRpcClient *ethHttpRpc.EthHttpRpcClient) (int64, error) {

	ethBlockHeightRpcMethod := ethHttpRpc.NewEthRpcMethod(ethRpcClient, ethRpcMethod)

	result := ethBlockHeightRpcMethod.Get()

	if err := result.Err; err != nil {
		return 0, errors.Join(errors.New("FAILED TO MARSHAL REQUEST"), err)
	}

	strResult, ok := result.Result.(string)
	if !ok {
		return 0, errors.New("FAILED TO CONVERT BLOCK HEIGHT TO STRING")
	}

	blockHeight, err := strconv.ParseInt(strResult, 0, 64)
	if err != nil {
		return 0, errors.Join(errors.New("FAILED TO CONVERT BLOCK HEIGHT TO NUMBER"), err)
	}

	log.Println(blockHeight)
	return blockHeight, nil
}

func EthBlockHeightMetric(ethRpcClient *ethHttpRpc.EthHttpRpcClient, labels map[string]string) {
	blockHeight, err := getEthBlockHeight(ethRpcClient)
	if err != nil {
		fmt.Printf("Get BlockHeight request failed, %v, %v \n", labels["name"], err)
	}
	promclient.EthBlockHeightGauge.WithLabelValues(labels["name"]).Set(float64(blockHeight))
	log.Println("BlockHeight of", labels["name"], "is", blockHeight)
}
