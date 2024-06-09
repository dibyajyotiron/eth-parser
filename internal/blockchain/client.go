package blockchain

import (
	"github.com/go_ether_parser/internal/parser"
)

type BlockchainClient interface {
	GetCurrentBlock() string
	GetBlockByNumber(blockNumber int) ([]parser.Transaction, error)
}
