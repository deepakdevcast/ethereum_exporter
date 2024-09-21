package ethhttprpc

import (
	"errors"
	"net/url"
)

type EthHttpRpcClient struct {
	URL *url.URL
}

type EthOptionFunc func(ethHttpRpcClient *EthHttpRpcClient)

func SetHttpRequest(url *url.URL) func(*EthHttpRpcClient) {
	return func(ethHttpRpcClient *EthHttpRpcClient) {
		ethHttpRpcClient.URL = url
	}
}

func NewEthHttpRpcClient(options ...EthOptionFunc) (*EthHttpRpcClient, error) {
	url, err := url.Parse("http://localhost:8545")
	if err != nil {
		return nil, errors.New("url Parsing Failed")
	}

	ethHttpRpcClient := &EthHttpRpcClient{
		url,
	}

	for _, opt := range options {
		opt(ethHttpRpcClient)
	}

	return ethHttpRpcClient, nil
}
