package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go_ether_parser/internal/manager"
)

type HttpError struct {
	ErrorString string
}

var METHOD_NOT_ALLOWED = &HttpError{"Method not allowed"}

type Handler struct {
	manager manager.Manager
	wg      *sync.WaitGroup
}

func NewHandler(manager manager.Manager, wg *sync.WaitGroup) *Handler {
	return &Handler{manager: manager, wg: wg}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.wg.Add(1)
	defer h.wg.Done()

	switch r.URL.Path {
	case "/current_block":
		if r.Method == http.MethodGet {
			h.getCurrentBlock(w, r)
		} else {
			http.Error(w, METHOD_NOT_ALLOWED.ErrorString, http.StatusMethodNotAllowed)
		}
	case "/subscribe":
		if r.Method == http.MethodPost {
			h.subscribe(w, r)
		} else {
			http.Error(w, METHOD_NOT_ALLOWED.ErrorString, http.StatusMethodNotAllowed)
		}
	case "/transactions":
		if r.Method == http.MethodPost {
			h.getTransactions(w, r)
		} else {
			http.Error(w, METHOD_NOT_ALLOWED.ErrorString, http.StatusMethodNotAllowed)
		}
	case "/storage":
		h.getStorage(w, r)
	default:
		http.NotFound(w, r)

	}
}

func (h *Handler) getCurrentBlock(w http.ResponseWriter, r *http.Request) {
	block := h.manager.GetCurrentBlock()
	w.Header().Set("Content-Type", "application/json") // setting content type as without this, postman always displays output json as text
	json.NewEncoder(w).Encode(map[string]string{"current_block": block})
}

func (h *Handler) subscribe(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Address string `json:"address"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	success := h.manager.Subscribe(req.Address)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"subscribed": success})
}

func (h *Handler) getTransactions(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Address string `json:"address"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	transactions := h.manager.GetTransactions(req.Address)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (h *Handler) getStorage(w http.ResponseWriter, r *http.Request) {
	storage := h.manager.GetStorage()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage)
}
