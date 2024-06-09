package parser

import "context"

type Transaction struct {
	Hash        string
	From        string
	To          string
	Value       string
	BlockHash   string
	BlockNumber string
}

type Parser interface {
	GetStorage() interface{}
	GetCurrentBlock() string
	Subscribe(address string) bool
	GetTransactions(address string) []Transaction
	Start(ctx context.Context)
}
