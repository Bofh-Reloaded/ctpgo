package abi

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
)

type DataHolder struct {
	data []byte
	offset int
}

func NewDataHolder(data []byte) *DataHolder {
	return &DataHolder{
		data: data,
		offset: 0,
	}
}

func (this *DataHolder) incOffset (size int) {
	this.offset += size
}

func (this *DataHolder) ReadBigInt() *big.Int {
	res := abi.ReadInteger(Uint256Type, this.data[this.offset:this.offset+32])
	this.incOffset(32)
	return res.(*big.Int)
}
