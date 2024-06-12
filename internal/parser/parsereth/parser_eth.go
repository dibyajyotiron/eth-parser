package parsereth

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/go_ether_parser/internal/blockchain"
	"github.com/go_ether_parser/internal/config"
	"github.com/go_ether_parser/internal/parser"
	"github.com/go_ether_parser/internal/rpcclient"
	"github.com/go_ether_parser/internal/storageengine"
)

type EthParser struct {
	storage    storageengine.StorageEngine
	updateFreq time.Duration
	client     blockchain.EthBlockchainClient
}

func NewEthParser(storage storageengine.StorageEngine) *EthParser {
	return &EthParser{
		storage:    storage,
		updateFreq: 5 * time.Second,
		client: *blockchain.NewEthBlockchainClient(
			rpcclient.NewRpcClient(config.Cfg.ClientURLs.EthRpcClientURL),
		),
	}
}

func (e *EthParser) fetchCurrentBlock() string {
	blockNumber, err := e.client.GetCurrentBlock()
	if err != nil {
		log.Printf("Error in GetCurrentBlock: %+v", err)
		return ""
	}

	e.storage.SetCurrentBlock(blockNumber)
	return blockNumber
}

func (e *EthParser) parseBlock(blockNumber string) {
	txs, err := e.client.GetBlockByNumber(blockNumber)
	if err != nil {
		log.Printf("Error in GetBlockByNumber: %+v", err)
		return
	}

	for _, tx := range txs {
		e.storage.AddTransaction(tx.From, tx)
		if tx.From != tx.To {
			e.storage.AddTransaction(tx.To, tx)
		}
	}
}

func (e *EthParser) GetCurrentBlock() string {
	return e.storage.GetCurrentBlock()
}

func (e *EthParser) Subscribe(address string) bool {
	return e.storage.Subscribe(address)
}

func (e *EthParser) GetTransactions(address string) []parser.Transaction {
	return e.storage.GetTransactions(address)
}

func (e *EthParser) GetStorage() interface{} {
	return e.storage.GetEntireStorage()
}

func (e *EthParser) Start(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Printf("Exiting start")
			return
		default:
			currentBlock := e.fetchCurrentBlock()
			if currentBlock == "" {
				log.Printf("Current block fetching failed! Investigate!")
				return
			}
			e.parseBlock(currentBlock)
			time.Sleep(e.updateFreq)
		}
	}
}
