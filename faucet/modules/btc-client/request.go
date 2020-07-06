package btc_client

import (
	"btc-faucet.drozd.by/modules/config"
	"github.com/KeisukeYamashita/jsonrpc"
)

var client *jsonrpc.RPCClient


func request(method string, params ...interface{}) (response *jsonrpc.RPCResponse, err error) {
	if client == nil {
		rpcClient := jsonrpc.NewRPCClient(config.Config().BitcoinRPC.URL)
		rpcClient.SetBasicAuth(config.Config().BitcoinRPC.User, config.Config().BitcoinRPC.Password)

		client = rpcClient
	}
	response, err = client.Call(method, params...)
	return
}
