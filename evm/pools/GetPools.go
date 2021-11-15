package pools

import (
	"context"
	coreabi "ctpgo.app/ctpgo/core/abi"
	corecommon "ctpgo.app/ctpgo/core/common"
	"ctpgo.app/ctpgo/core/data"
	"ctpgo.app/ctpgo/core/wiring"
	"ctpgo.app/ctpgo/evm/evmclient"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/panjf2000/ants/v2"
	"log"
	"math/big"
	"sync"
)

func GetPoolsForStdExchange(exchange *corecommon.Exchange) []*data.PoolData {
	var evmClient evmclient.EvmClientService
	wiring.Bind(&evmClient)

	size := getFactoryPools(&evmClient, exchange).Int64()

	results := make(chan *data.PoolData, size)

	maxGoroutines := corecommon.MAX_GOROUTINES
	log.Println("Starting for", exchange.Name, "with goal", size)

	var wg sync.WaitGroup
	wp, _ := ants.NewPoolWithFunc(maxGoroutines, func(i interface{}) {
		defer wg.Done()
		addr := GetPoolAddressFromFactory(i.(int64), &evmClient, exchange)
		t0 := GetPoolToken(true, &evmClient, addr)
		t1 := GetPoolToken(false, &evmClient, addr)
		p := &data.PoolData{
			Id:          uint64(i.(int64)),
			Address:     addr,
			Token0:      t0,
			Token1:      t1,
			BaseWeight:  50,
			SwapFee:     exchange.Fee,
		}
		results <- p
	})
	defer wp.Release()
	var pools []*data.PoolData
	for i := int64(0); i < size; i++ {
		wg.Add(1)
		_ = wp.Invoke(i)
	}

	for i := int64(0); i < size; i++ {
		p := <-results
		pools = append(pools, p)
	}
	wg.Wait()
	close(results)
	log.Println("Finished for ", exchange.Name, " Done @ ", len(pools))
	return pools
}

func GetPoolReserves(evmService *evmclient.EvmClientService, pool corecommon.Address) *data.PoolReserves {
	calldata := coreabi.ABIBytes("getReserves()")
	msg := ethereum.CallMsg{
		From: corecommon.ZeroAddress,
		To:   &pool,
		Data: calldata,
	}

	cli := (*evmService).Client()
	res, err := cli.CallContract(context.Background(), msg, nil)
	//cli.Close()
	corecommon.Assert("[GetPoolReserves]", err)

	dh := coreabi.NewDataHolder(res)
	r0 := dh.ReadBigInt()
	r1 := dh.ReadBigInt()

	return &data.PoolReserves{Address: pool, R0: r0, R1: r1}
}

func GetPoolAddressFromFactory(index int64, evmService *evmclient.EvmClientService, exchange *corecommon.Exchange) corecommon.Address {
	opcode := coreabi.ABIBytes("allPairs(uint256)")

	args := abi.Arguments{abi.Argument{Type: coreabi.Uint256Type}}
	inputs, _ := args.Pack(big.NewInt(index))
	calldata := append(opcode, inputs...)
	msg := ethereum.CallMsg{
		From: corecommon.ZeroAddress,
		To:   &exchange.Factory,
		Data: calldata,
	}

	cli := (*evmService).Client()
	res, err := cli.CallContract(context.Background(), msg, nil)
	// cli.Close()
	corecommon.Assert("[GetPoolAddressFromFactory]", err)
	return coreabi.ReadAddress(res)
}

func GetPoolToken(t0 bool, evmService *evmclient.EvmClientService, pool corecommon.Address) corecommon.Address {
	//defer un(trace("getPoolToken"))
	var calldata []byte
	if t0 {
		calldata = coreabi.ABIBytes("token0()")
	} else {
		calldata = coreabi.ABIBytes("token1()")
	}

	msg := ethereum.CallMsg{
		From: corecommon.ZeroAddress,
		To:   &pool,
		Data: calldata,
	}

	cli := (*evmService).Client()
	res, err := cli.CallContract(context.Background(), msg, nil)
	// cli.Close()
	corecommon.Assert("[GetPoolAddressFromFactory]", err)
	return coreabi.ReadAddress(res)
}

func getFactoryPools(evmService *evmclient.EvmClientService, exchange *corecommon.Exchange) *big.Int {
	calldata := coreabi.ABIBytes("allPairsLength()")
	msg := ethereum.CallMsg{
		From: corecommon.ZeroAddress,
		To:   &exchange.Factory,
		Data: calldata,
	}

	cli := (*evmService).Client()
	// cli.Close()
	res, err := cli.CallContract(context.Background(), msg, nil)
	corecommon.Assert("[GetFactoryPools]", err)
	size := abi.ReadInteger(coreabi.Uint256Type, res)
	return size.(*big.Int)
}
