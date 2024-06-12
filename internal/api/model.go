package api

import (
	"github.com/go_ether_parser/internal/parser"
)

type BaseApiSuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type GetCurrentBlockResponse struct {
	BlockNumber string `json:"block_number"`
}

type SubscribeRequest struct {
	Address string `json:"address"`
}

type SubscribeResponse struct {
	Subscribed bool `json:"subscribed"`
}

type GetTransactionsRequest struct {
	Address string `json:"address"`
}

type GetTransactionsResponse struct {
	Transactions []parser.Transaction `json:"transactions"`
}
