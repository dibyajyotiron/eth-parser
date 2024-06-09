package mockstorage

import (
	"sync"

	"github.com/go_ether_parser/internal/parser"
)

var (
	TestCurrentBlock = "0x131ed3b"
	TestAddress      = "0x00000000a991c429ee2ec6df19d40fe0c80088b8"
	TestTransactions = map[string][]parser.Transaction{
		TestAddress: {
			{
				Hash:        "0x4fb577afe9aa478b8d0a37fa50da8ea92a61a617bc34de3fcac1a2e63910554b",
				From:        "0x77ad3a15b78101883af36ad4a875e17c86ac65d1",
				To:          TestAddress,
				Value:       "0x2487e47",
				BlockHash:   "0x764bee0f55902a23cf7b19d29ad4c43e0f7010b091477df493ee6e5a0731ab0f",
				BlockNumber: "0x131ece3",
			},
			{
				Hash:        "0x52e7f7bc7f4a439a03f76b6d65d29da6b1fa1e078ab5fff76fb303f84d2bb332",
				From:        "0x77ad3a15b78101883af36ad4a875e17c86ac65d1",
				To:          TestAddress,
				Value:       "0x21dab1a",
				BlockHash:   "0x764bee0f55902a23cf7b19d29ad4c43e0f7010b091477df493ee6e5a0731ab0f",
				BlockNumber: "0x131ece3",
			},
		},
		"0x77ad3a15b78101883af36ad4a875e17c86ac65d1": {
			{
				Hash:        "0x4fb577afe9aa478b8d0a37fa50da8ea92a61a617bc34de3fcac1a2e63910554b",
				From:        "0x77ad3a15b78101883af36ad4a875e17c86ac65d1",
				To:          TestAddress,
				Value:       "0x2487e47",
				BlockHash:   "0x764bee0f55902a23cf7b19d29ad4c43e0f7010b091477df493ee6e5a0731ab0f",
				BlockNumber: "0x131ece3",
			},
			{
				Hash:        "0x52e7f7bc7f4a439a03f76b6d65d29da6b1fa1e078ab5fff76fb303f84d2bb332",
				From:        "0x77ad3a15b78101883af36ad4a875e17c86ac65d1",
				To:          TestAddress,
				Value:       "0x21dab1a",
				BlockHash:   "0x764bee0f55902a23cf7b19d29ad4c43e0f7010b091477df493ee6e5a0731ab0f",
				BlockNumber: "0x131ece3",
			},
		},
	}
)

type MockStorage struct {
	CurrentBlock  string
	Subscriptions map[string]bool
	Transactions  map[string][]parser.Transaction
	mu            sync.Mutex
}

func NewMemoryStorage() *MockStorage {
	return &MockStorage{
		CurrentBlock: TestCurrentBlock,
		Subscriptions: map[string]bool{
			TestAddress: true,
		},
		Transactions: TestTransactions,
	}
}

func (m *MockStorage) GetCurrentBlock() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.CurrentBlock
}

func (m *MockStorage) SetCurrentBlock(block string) {
	if block == TestCurrentBlock {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.CurrentBlock = block
}

func (m *MockStorage) Subscribe(address string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.Subscriptions[address]; exists {
		return false
	}
	m.Subscriptions[address] = true
	return true
}

func (m *MockStorage) GetTransactions(address string) []parser.Transaction {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.Transactions[address]
}

func (m *MockStorage) AddTransaction(address string, tx parser.Transaction) {
	if address == TestAddress {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Transactions[address] = append(m.Transactions[address], tx)
}

func (m *MockStorage) GetEntireStorage() interface{} {
	return m
}
