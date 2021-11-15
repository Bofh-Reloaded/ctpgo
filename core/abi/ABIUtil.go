package abi
import (
	"github.com/ethereum/go-ethereum/common"
	corecommon "ctpgo.app/ctpgo/core/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	Uint256Type = abi.Type{Size: 256, T: abi.UintTy}
	AddressType = abi.Type{Size: 20, T: abi.AddressTy}
)

func ABIBytes(method string) []byte {
	return crypto.Keccak256([]byte(method))[:4]
}

func ReadAddress(data []byte) corecommon.Address {
	return common.BytesToAddress(data)
}
