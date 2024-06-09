package main

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/go_ether_parser/internal/manager"
	"github.com/go_ether_parser/internal/parser/parsereth"
	storage "github.com/go_ether_parser/internal/storageengine/storage_engine_memory"
	"github.com/go_ether_parser/server"
)

func main() {
	wg := &sync.WaitGroup{}

	ctx, cancelFunc := context.WithCancel(context.Background())

	memStorage := storage.NewMemoryStorage()

	ethParser := parsereth.NewEthParser(memStorage)
	ethManager := manager.NewEthManager(ethParser, ctx)

	go ethManager.Start(ctx, wg)

	srv := server.NewServer(ethManager, wg)
	go func() {
		log.Printf("Server starting on :8080")
		if err := srv.Start(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not start server: %s", err)
		}
	}()

	if err := srv.Shutdown(ctx, cancelFunc); err != nil {
		log.Fatalf("could not gracefully shut down the server: %s", err)
	}
	log.Println("Server gracefully stopped")
}
