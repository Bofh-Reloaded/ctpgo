package common

import "math/big"

func FromWei(value *big.Int) float64 {
	f := new(big.Float).SetInt(value)
	r, _ := f.Float64()
	return r / 1e18
}
