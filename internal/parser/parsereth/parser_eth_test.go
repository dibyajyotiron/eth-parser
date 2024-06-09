package parsereth

import (
	"testing"

	"github.com/go_ether_parser/internal/storageengine/storage_engine_memory"
)

func TestMain(t *testing.T) {
	t.Run("TestSubscribe", func(t *testing.T) {
		storage := storage_engine_memory.NewMemoryStorage()
		parser := NewEthParser(storage)
		address := "0X000"
		if !parser.Subscribe(address) {
			t.Errorf("Expected true, got false")
		}
		if parser.Subscribe(address) {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("TestGetTransactions", func(t *testing.T) {
		storage := storage_engine_memory.NewMemoryStorage()
		parser := NewEthParser(storage)
		address := "0X000"
		txs := parser.GetTransactions(address)
		if len(txs) != 0 {
			t.Errorf("Expected 0 transactions, got %d", len(txs))
		}
	})

	t.Run("TestGetCurrentBlock", func(t *testing.T) {
		storage := storage_engine_memory.NewMemoryStorage()
		parser := NewEthParser(storage)
		block := parser.GetCurrentBlock()
		if block != "" {
			t.Errorf("Expected block 0, got %s", block)
		}
	})
}
