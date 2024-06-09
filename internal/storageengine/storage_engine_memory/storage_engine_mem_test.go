package storage_engine_memory

import (
	"testing"

	"github.com/go_ether_parser/internal/parser"
)

func TestMemoryStorage(t *testing.T) {
	storage := NewMemoryStorage()

	address := "0X000"
	// Test SetCurrentBlock and GetCurrentBlock
	t.Run("SetCurrentBlock and GetCurrentBlock", func(t *testing.T) {
		storage.SetCurrentBlock("10")
		if block := storage.GetCurrentBlock(); block != "10" {
			t.Errorf("Expected block 10, got %s", block)
		}

		if !storage.Subscribe(address) {
			t.Errorf("Expected true, got false")
		}
		if storage.Subscribe(address) {
			t.Errorf("Expected false, got true")
		}
		if len(storage.GetTransactions(address)) != 0 {
			t.Errorf("Expected 0 transactions, got %d", len(storage.GetTransactions(address)))
		}
	})

	t.Run("AddTransaction and GetTransactions", func(t *testing.T) {
		tx := parser.Transaction{
			Hash:        "somehash",
			From:        "someaddress",
			To:          "anotheraddress",
			Value:       "100",
			BlockHash:   "randomblockhash",
			BlockNumber: "0x111",
		}
		storage.AddTransaction(address, tx)
		if len(storage.GetTransactions(address)) != 1 {
			t.Errorf("Expected 1 transaction, got %d", len(storage.GetTransactions(address)))
		}
	})
}
