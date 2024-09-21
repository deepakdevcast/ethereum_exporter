package ethhttprpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type RpcMethod struct {
	ethHttpRpcClient *EthHttpRpcClient
	rpcRequestBody   *RpcRequestBody
}

type RpcRequestBody struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      uint          `json:"id"`
}

type RpcResponse struct {
	Id      uint   `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  any    `json:"result,omitempty"`
	Error   *struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`
}

type MethodRpcResult struct {
	Result any
	Err    error
}

func NewEthRpcMethod(ethClient *EthHttpRpcClient, methodBody *RpcRequestBody) *RpcMethod {
	return &RpcMethod{
		ethClient,
		methodBody,
	}
}

func (rm *RpcMethod) Get() MethodRpcResult {
	jsonData, err := json.Marshal(rm.rpcRequestBody)
	if err != nil {
		return MethodRpcResult{nil, errors.Join(err)}
	}

	resp, err := http.Post(rm.ethHttpRpcClient.URL.String(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return MethodRpcResult{nil, errors.Join(err)}
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return MethodRpcResult{nil, errors.Join(err)}
	}

	var respResult RpcResponse
	if err := json.Unmarshal(body, &respResult); err != nil {
		return MethodRpcResult{nil, errors.Join(err)}
	}
	if respResult.Error != nil {
		return MethodRpcResult{nil, errors.New(respResult.Error.Message)}
	}
	return MethodRpcResult{respResult.Result, nil}
}

func (rm *RpcMethod) GetAsync(resultChan chan<- MethodRpcResult) {
	go func() {
		resultChan <- rm.Get()
	}()
}
