package data

import "ctpgo.app/ctpgo/core/common"

type PoolData struct {
	Id uint64
	Address common.Address
	Token0 common.Address
	Token1 common.Address
	BaseWeight int
	SwapFee int
	Empty bool
}

type TokenData struct {
	Id uint64
	Address common.Address
	Symbol string
	Decimals int
}

type DataContainer struct {
	pools map[string]*PoolData
}
