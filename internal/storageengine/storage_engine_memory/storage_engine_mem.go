package storage_engine_memory

import (
	"sync"

	"github.com/go_ether_parser/internal/parser"
)

type MemoryStorage struct {
	CurrentBlock  string
	Subscriptions map[string]bool
	Transactions  map[string][]parser.Transaction
	mu            sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		CurrentBlock:  "",
		Subscriptions: make(map[string]bool),
		Transactions:  make(map[string][]parser.Transaction),
	}
}

func (m *MemoryStorage) GetCurrentBlock() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.CurrentBlock
}

func (m *MemoryStorage) SetCurrentBlock(block string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.CurrentBlock = block
}

func (m *MemoryStorage) Subscribe(address string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.Subscriptions[address]; exists {
		return false
	}
	m.Subscriptions[address] = true
	return true
}

func (m *MemoryStorage) GetTransactions(address string) []parser.Transaction {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.Transactions[address]
}

func (m *MemoryStorage) AddTransaction(address string, tx parser.Transaction) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Transactions[address] = append(m.Transactions[address], tx)
}

func (m *MemoryStorage) GetEntireStorage() interface{} {
	return m
}
