package data

import (
	"ctpgo.app/ctpgo/core/common"
	"math/big"
)

type PoolReserves struct {
	Address common.Address
	R0      *big.Int
	R1      *big.Int
}

type Reserves struct {
	R0 *big.Int
	R1 *big.Int
}
