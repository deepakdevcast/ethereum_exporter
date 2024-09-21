package ethhttprpc

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func TestRpcMethod(t *testing.T) {
	url, _ := url.Parse("https://rpc.ankr.com/eth_sepolia")
	ethClient, err := NewEthHttpRpcClient(SetHttpRequest(url))
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	rpcMethod := NewEthRpcMethod(ethClient, &RpcRequestBody{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		Id:      83,
		Params:  []interface{}{},
	})
	result := rpcMethod.Get()
	if result.Err != nil {
		log.Fatalln(err)
	}
	if result.Result == nil {
		log.Fatalln("Response is empty", result.Result)
	}
	fmt.Println(result.Result)
	// fmt.Println(strconv.ParseInt(result.Result.(string), 0, 64))
}
