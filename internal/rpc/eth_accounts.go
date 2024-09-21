package rpc

import (
	"errors"
	ethHttpRpc "ethereum_exporter/internal/eth/http_rpc"
	"ethereum_exporter/internal/promclient"
	"fmt"
	"log"
)

var ethAccountsBody = &ethHttpRpc.RpcRequestBody{
	Jsonrpc: "2.0",
	Method:  "eth_accounts",
	Id:      83,
	Params:  []interface{}{},
}

func getEthAccounts(ethRpcClient *ethHttpRpc.EthHttpRpcClient) ([]string, error) {
	ethAccountsRpcMethod := ethHttpRpc.NewEthRpcMethod(ethRpcClient, ethAccountsBody)

	result := ethAccountsRpcMethod.Get()

	if err := result.Err; err != nil {
		return []string{}, errors.Join(ErrJsonMarshal, err)
	}

	return result.Result.([]string), nil
}

func EthAccountsMetric(ethRpcClient *ethHttpRpc.EthHttpRpcClient, labels map[string]string) {
	Accounts, err := getEthAccounts(ethRpcClient)

	if err != nil {
		fmt.Printf("Get Accounts request failed, %v, %v \n", labels["name"], err)
		return
	}
	for _, account := range Accounts {
		promclient.EthAccountsGauge.WithLabelValues(labels["name"], account).Set(1)
	}
	log.Println("Accounts of", labels["name"], "is", Accounts)
}
