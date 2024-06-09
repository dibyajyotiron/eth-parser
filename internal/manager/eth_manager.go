package manager

import (
	"context"
	"sync"

	"github.com/go_ether_parser/internal/parser"
)

type EthManager struct {
	parser parser.Parser
	ctx    context.Context
}

func NewEthManager(parser parser.Parser, ctx context.Context) *EthManager {
	return &EthManager{
		parser: parser,
		ctx:    ctx,
	}
}

func (m *EthManager) GetCurrentBlock() string {
	return m.parser.GetCurrentBlock()
}

func (m *EthManager) GetCtx() context.Context {
	return m.ctx
}

func (m *EthManager) Subscribe(address string) bool {
	return m.parser.Subscribe(address)
}

func (m *EthManager) GetTransactions(address string) []parser.Transaction {
	return m.parser.GetTransactions(address)
}

func (m *EthManager) GetStorage() interface{} {
	return m.parser.GetStorage()
}

// Start will start the parser as a long running process
func (m *EthManager) Start(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	m.parser.Start(ctx)
}
