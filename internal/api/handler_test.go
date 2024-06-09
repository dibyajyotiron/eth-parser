package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/go_ether_parser/internal/manager"
	"github.com/go_ether_parser/internal/parser/parsereth"
	mockStorage "github.com/go_ether_parser/internal/storageengine/mock_storage"
)

func TestMain(t *testing.T) {
	memStorage := mockStorage.NewMemoryStorage()
	parser := parsereth.NewEthParser(memStorage)
	ethManager := manager.NewEthManager(parser, context.Background())

	t.Run("Subscribe", func(t *testing.T) {
		handler := NewHandler(ethManager, &sync.WaitGroup{})

		reqBody := []byte(`{"address":"0X000"}`)
		req, err := http.NewRequest("POST", "/subscribe", bytes.NewBuffer(reqBody))

		if err != nil {
			t.Errorf("Handler should return successfully")
		}

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if http.StatusOK != rr.Code {
			t.Errorf("Handler should return 200 status code")
		}

		var response map[string]bool
		err = json.NewDecoder(rr.Body).Decode(&response)
		if err != nil {
			t.Errorf("Decoder should not throw error")
		}

		if !response["subscribed"] {
			t.Errorf("Handler should subscribe successfully")
		}

	})
	t.Run("GetCurrentBlock", func(t *testing.T) {
		handler := NewHandler(ethManager, &sync.WaitGroup{})

		req, err := http.NewRequest("GET", "/current_block", nil)
		if err != nil {
			t.Errorf("Handler should return successfully")
		}

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if http.StatusOK != rr.Code {
			t.Errorf("Handler should return 200 status code")
		}

		var response map[string]string
		err = json.NewDecoder(rr.Body).Decode(&response)
		if err != nil {
			t.Errorf("Decoding Response body should not error out")
		}

		if response["current_block"] != mockStorage.TestCurrentBlock {
			t.Errorf("Handler should return mocked block number")
		}
	})

	t.Run("GetTransactions", func(t *testing.T) {
		handler := NewHandler(ethManager, &sync.WaitGroup{})
		reqBody := []byte(fmt.Sprintf(`{"address": "%s"}`, mockStorage.TestAddress))

		req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Errorf("Handler should return successfully")
		}

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if http.StatusOK != rr.Code {
			t.Errorf("Handler should return 200 status code")
		}
		fmt.Printf("%+v", rr.Body)
		var response []map[string]string
		err = json.NewDecoder(rr.Body).Decode(&response)
		if err != nil {
			t.Errorf("Decoding Response body should not error out")
		}

		if len(response) != 2 {
			t.Errorf("Response should have length of 2, Current: %d", len(response))
		}
	})

}
