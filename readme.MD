# Ethereum Parser

## Assignment Details
### Goal
Implement Ethereum blockchain parser that will allow to query transactions for subscribed
addresses.
### Problem
Users not able to receive push notifications for incoming/outgoing transactions. By
Implementing Parser interface we would be able to hook this up to notifications service to
notify about any incoming/outgoing transactions.
### Limitations
    - Use Go Language
    - Avoid usage of external libraries
    - Use Ethereum JSONRPC to interact with Ethereum Blockchain
    - Use memory storage for storing any data (should be easily extendable to support any
    storage in the future)

Expose public interface for external usage either via code or command line or rest api that will include supported list of operations defined in the Parser interface

``` 
type Parser interface {
    // last parsed block
    GetCurrentBlock() int
    // add address to observer
    Subscribe(address string) bool
    // list of inbound or outbound transactions for an address
    GetTransactions(address string) []Transaction
}
```

### Request example
Endpoint URL: https://cloudflare-eth.com

```
curl -X POST '<https://cloudflare-eth.com>' --data
'{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":83}
'
``` 

#### Result
{
"id":83,
"jsonrpc": "2.0",
"result": "0x4b7" // 1207
}

### References
- Ethereum JSON RPC Interface
Note
- keep it simple
- try to finish the task within 4 hours. We do not track the time spent, this is just a
guidance. We do not ask for a perfect production ready service

## How to run

1. Clone the repository
2. Run `go mod tidy`
3. Run `go run cmd/main.go` or `./run.sh`
4. If you want to send sigkill for whatever reason, use `./kill.sh`

## Endpoints

- `GET /current_block`: Get the current block number
- `POST /subscribe`: Subscribe to an address
    - Request Body -> 
        ```
        {
            "address": "0x000000633b68f5d8d3a86593ebb815b4663bcbe0"
        }
- `POST /transactions`: Get transactions for an address
    - Request Body -> 
        ```
        {
            "address": "0x000000633b68f5d8d3a86593ebb815b4663bcbe0"
        }

## Approach
- Code is written in a modular approach, for easy customization and replaceable modules for loose coupling.
- Most modules have tests, to determine the sanity of the code
- Server has graceful stopping
- Only default config is provided for the sake of time constraint and simplicity as asked in requirements