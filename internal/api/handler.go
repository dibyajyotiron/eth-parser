package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go_ether_parser/internal/parser"
)

type HttpError struct {
	ErrorString string
}

var METHOD_NOT_ALLOWED = &HttpError{"Method not allowed"}

type Handler struct {
	parser parser.Parser
}

func NewHandler(parser parser.Parser, wg *sync.WaitGroup) *Handler {
	return &Handler{parser: parser}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) wrapSuccessResponse(data interface{}) BaseApiSuccessResponse {
	return BaseApiSuccessResponse{
		Success: true,
		Data:    data,
	}
}

func (h *Handler) getCurrentBlock(w http.ResponseWriter, _ *http.Request) {
	block := h.parser.GetCurrentBlock()
	w.Header().Set("Content-Type", "application/json") // setting content type as without this, postman always displays output json as text
	json.NewEncoder(w).Encode(h.wrapSuccessResponse(block))
}

func (h *Handler) subscribe(w http.ResponseWriter, r *http.Request) {
	var req SubscribeRequest
	json.NewDecoder(r.Body).Decode(&req)
	isSubscribed := h.parser.Subscribe(req.Address)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.wrapSuccessResponse(SubscribeResponse{
		Subscribed: isSubscribed,
	}))
}

func (h *Handler) getTransactions(w http.ResponseWriter, r *http.Request) {
	var req GetTransactionsRequest
	json.NewDecoder(r.Body).Decode(&req)
	transactions := h.parser.GetTransactions(req.Address)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.wrapSuccessResponse(GetTransactionsResponse{
		Transactions: transactions,
	}))
}

func (h *Handler) getStorage(w http.ResponseWriter, _ *http.Request) {
	storage := h.parser.GetStorage()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.wrapSuccessResponse(map[string]interface{}{"storage": storage}))
}
