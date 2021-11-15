package common

import "github.com/ethereum/go-ethereum/common"

type Address = common.Address
var ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

func AsAddress(s string) common.Address {
	return common.HexToAddress(s)
}