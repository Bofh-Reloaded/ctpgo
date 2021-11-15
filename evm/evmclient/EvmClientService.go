package evmclient

import "github.com/ethereum/go-ethereum/ethclient"

type EvmClientService interface {
	Client() *ethclient.Client
}
