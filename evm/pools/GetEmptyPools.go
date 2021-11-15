package pools

import (
	"ctpgo.app/ctpgo/core/common"
	"ctpgo.app/ctpgo/core/data"
	"ctpgo.app/ctpgo/core/wiring"
	"ctpgo.app/ctpgo/evm/evmclient"
	"github.com/panjf2000/ants/v2"
	"log"
	"math/big"
	"sync"
)

var zero = big.NewInt(0)

func GetEmptyPools(pools []*data.PoolData) map[common.Address]bool {
	log.Println("GetEmptyPools")
	var evmClient evmclient.EvmClientService
	wiring.Bind(&evmClient)

	results := make(map[common.Address]bool)
	prices := make(map[common.Address]float64)
	reserves := make(map[common.Address]*data.Reserves)

	prices[common.AsAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c")] = 310
	prices[common.AsAddress("0xe9e7cea3dedca5984780bafc599bd69add087d56")] = 1.0

	reservesChan := make(chan *data.PoolReserves, len(pools))

	var wg sync.WaitGroup
	wp, _ := ants.NewPoolWithFunc(common.MAX_GOROUTINES, func(i interface{}) {
		defer wg.Done()
		p := i.(*data.PoolData)
		r := GetPoolReserves(&evmClient, p.Address)
		reservesChan <- r
	})
	defer wp.Release()

	for _, p := range pools {
		wg.Add(1)
		_ = wp.Invoke(p)
	}

	for i := 0; i < len(pools); i++ {
		p := <-reservesChan
		if p.R0.Cmp(zero) <= 0 || p.R1.Cmp(zero) <= 0 {
			results[p.Address] = true
		} else {
			res := &data.Reserves{R0: p.R0, R1: p.R1}
			reserves[p.Address] = res
		}
	}
	wg.Wait()

	log.Println("Resolve prices",len(reserves),"pools remaining",len(results),"already excluded)...")
	for _, p := range pools {
		if r, found := reserves[p.Address]; found {
			t0Price,t0f := prices[p.Token0]
			t1Price,t1f := prices[p.Token1]
			if (t0f && !t1f) {
				t1Price = common.FromWei(r.R0) * t0Price / common.FromWei(r.R1)
				prices[p.Token1] = t1Price
			} else if (!t0f && t1f) {
				t0Price = common.FromWei(r.R1) * t1Price / common.FromWei(r.R0)
				prices[p.Token0] = t0Price
			}
		}
	}

	for _, p := range pools {
		if r, found := reserves[p.Address]; found {
			t0Price,t0f := prices[p.Token0]
			t1Price,t1f := prices[p.Token1]
			if (t0f && !t1f) {
				t1Price = common.FromWei(r.R0) * t0Price / common.FromWei(r.R1)
				prices[p.Token1] = t1Price
			} else if (!t0f && t1f) {
				t0Price = common.FromWei(r.R1) * t1Price / common.FromWei(r.R0)
				prices[p.Token0] = t0Price
			}
			t0Price,t0f = prices[p.Token0]
			t1Price,t1f = prices[p.Token1]
			if (t0f && t1f) {
				value := t0Price * common.FromWei(r.R0) + t1Price * common.FromWei(r.R1)
				if (value < 5.0) {
					results[p.Address] = true
				}
			} else {
				log.Println("Garbage pool",p.Address,"will be removed")
				results[p.Address] = true
			}
		}
	}
	log.Println("Empty pools (zero reserves)", len(results))
	close(reservesChan)
	return results
}
