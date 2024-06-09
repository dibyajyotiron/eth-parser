package blockchain

import (
	"github.com/go_ether_parser/internal/parser"
	"github.com/go_ether_parser/internal/rpcclient"
	mockstorage "github.com/go_ether_parser/internal/storageengine/mock_storage"
)

type EthBlockchainMockClient struct {
	client rpcclient.RpcClient
}

func NewEthBlockchainMockClient(client rpcclient.RpcClient) *EthBlockchainMockClient {
	return &EthBlockchainMockClient{client: nil}
}

func (e *EthBlockchainMockClient) GetCurrentBlock() string {
	return mockstorage.TestAddress
}

func (e *EthBlockchainMockClient) GetBlockByNumber(blockNumber string) ([]parser.Transaction, error) {
	return mockstorage.NewMemoryStorage().GetTransactions(mockstorage.TestAddress), nil
}
