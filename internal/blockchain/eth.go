package blockchain

import (
	"github.com/go_ether_parser/internal/parser"
	"github.com/go_ether_parser/internal/rpcclient"
)

type EthBlockchainClient struct {
	client rpcclient.RpcClient
}

func NewEthBlockchainClient(client rpcclient.RpcClient) *EthBlockchainClient {
	return &EthBlockchainClient{client: client}
}

func (e *EthBlockchainClient) GetCurrentBlock() (string, error) {
	var result struct {
		Result string `json:"result"`
	}
	err := e.client.Call("eth_blockNumber", nil, &result)
	if err != nil {
		return "", err
	}

	return result.Result, nil
}

func (e *EthBlockchainClient) GetBlockByNumber(blockNumber string) ([]parser.Transaction, error) {
	var result struct {
		Result struct {
			Transactions []parser.Transaction `json:"transactions"`
		} `json:"result"`
	}
	params := []interface{}{blockNumber, true}
	err := e.client.Call("eth_getBlockByNumber", params, &result)
	if err != nil {
		return nil, err
	}

	return result.Result.Transactions, nil
}
