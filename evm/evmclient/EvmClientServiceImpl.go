package evmclient

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type EvmClientServiceImpl struct {
	rpc *rpc.Client
	client *ethclient.Client
}

func NewRpcClient(isRpc bool) *rpc.Client {
	url := "/home/admin/geth-beta/node/geth.ipc"
	if (isRpc) {
		url = "http://100.64.200.1:8545"
	}
	rpccli, _ := rpc.Dial(url)
	return rpccli
}

func NewEvmClientServiceImpl(rpc bool) *EvmClientServiceImpl {
	rpccli := NewRpcClient(rpc)
	instance := &EvmClientServiceImpl{
		rpc: rpccli,
		client: ethclient.NewClient(rpccli),
	}
	return instance
}

func (obj *EvmClientServiceImpl) Client() *ethclient.Client {
	return obj.client
}
