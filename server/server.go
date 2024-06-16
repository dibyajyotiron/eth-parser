package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go_ether_parser/internal/api"
	"github.com/go_ether_parser/internal/parser"
)

type Server struct {
	httpServer *http.Server
	wg         *sync.WaitGroup
}

func NewServer(parser parser.Parser, wg *sync.WaitGroup) *Server {
	handler := api.NewHandler(parser, wg)
	httpServer := &http.Server{
		Handler: handler,
	}

	return &Server{httpServer: httpServer, wg: wg}
}

func (s *Server) Start(addr string) error {
	s.wg.Add(1)
	defer s.wg.Done()
	s.httpServer.Addr = addr
	return s.httpServer.ListenAndServe()
}

func (s *Server) checkForStopSignal(cancelFunc context.CancelFunc) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	signal := <-shutdown

	log.Printf("Received signal %s, Shutting down server...", signal)
	cancelFunc()
}

func (s *Server) Shutdown(ctx context.Context, cancelFunc context.CancelFunc) error {
	s.checkForStopSignal(cancelFunc)
	s.httpServer.SetKeepAlivesEnabled(false)
	return s.httpServer.Shutdown(context.Background())
}
