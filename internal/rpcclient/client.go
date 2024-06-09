package rpcclient

type RpcClient interface {
	Call(method string, params []interface{}, result interface{}) error
}
