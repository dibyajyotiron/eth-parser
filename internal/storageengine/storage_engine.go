package storageengine

import "github.com/go_ether_parser/internal/parser"

type StorageEngine interface {
	GetCurrentBlock() string
	SetCurrentBlock(block string)
	Subscribe(address string) bool
	GetTransactions(address string) []parser.Transaction
	AddTransaction(address string, tx parser.Transaction)
	GetEntireStorage() interface{}
}
