package main

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/go_ether_parser/internal/config"
	"github.com/go_ether_parser/internal/parser/parsereth"
	storage "github.com/go_ether_parser/internal/storageengine/storage_engine_memory"
	"github.com/go_ether_parser/server"
)

func main() {
	config.Load()

	wg := &sync.WaitGroup{}

	ctx, cancelFunc := context.WithCancel(context.Background())

	memStorage := storage.NewMemoryStorage()

	ethParser := parsereth.NewEthParser(memStorage)

	go ethParser.Start(ctx, wg)

	srv := server.NewServer(ethParser, wg)
	go func() {
		port := config.Cfg.App.ServerPort

		log.Printf("Server starting on %s", port)
		if err := srv.Start(port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not start server: %s", err)
		}
	}()

	if err := srv.Shutdown(ctx, cancelFunc); err != nil {
		log.Fatalf("could not gracefully shut down the server: %s", err)
	}

	wg.Wait()
	log.Println("Server gracefully stopped")
}
