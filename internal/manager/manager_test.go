package manager

import (
	"context"
	"testing"

	"github.com/go_ether_parser/internal/parser/parsereth"
	storage "github.com/go_ether_parser/internal/storageengine/storage_engine_memory"
)

func TestManager(t *testing.T) {
	memStorage := storage.NewMemoryStorage()
	parser := parsereth.NewEthParser(memStorage)
	ethManager := NewEthManager(parser, context.Background())

	t.Run("Current Block should be empty", func(t *testing.T) {
		b := ethManager.GetCurrentBlock()
		if b != "" {
			t.Error("Expected empty current block")
		}
	})

	t.Run("It should subscribe to address", func(t *testing.T) {
		b := ethManager.Subscribe("0X000")
		if !b {
			t.Error("Subscription should have been true")
		}
	})

	t.Run("It should not resubscribe to address", func(t *testing.T) {
		b := ethManager.Subscribe("0X000")
		if b {
			t.Error("Subscription should have been false")
		}
	})

	t.Run("Transactions should be empty", func(t *testing.T) {
		txs := ethManager.GetTransactions("0X000")
		if len(txs) > 0 {
			t.Error("Transactions should be empty for the given address")
		}
	})

}
