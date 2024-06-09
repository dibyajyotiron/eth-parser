package rpcclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RpcClientImpl struct {
	url string
}

func NewRpcClient(url string) RpcClient {
	return &RpcClientImpl{url: url}
}

func (c *RpcClientImpl) Call(method string, params []interface{}, result interface{}) error {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1,
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	resp, err := http.Post(c.url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(result)
}
