package manager

import (
	"context"
	"sync"

	"github.com/go_ether_parser/internal/parser"
)

type Manager interface {
	GetCtx() context.Context
	GetCurrentBlock() string
	Subscribe(address string) bool
	GetTransactions(address string) []parser.Transaction
	GetStorage() interface{}
	Start(ctx context.Context, wg *sync.WaitGroup)
}
