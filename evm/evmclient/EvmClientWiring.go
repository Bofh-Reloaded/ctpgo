package evmclient

import (
	"ctpgo.app/ctpgo/core/wiring"
)

func EthClientWiring(rpc bool) {
	wiring.Wire(func() EvmClientService {
	    return NewEvmClientServiceImpl(rpc)
	})
}
