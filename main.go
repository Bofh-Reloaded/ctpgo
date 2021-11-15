package main

import (
	"ctpgo.app/ctpgo/bsc/uclone"
	"ctpgo.app/ctpgo/core/data"
	"ctpgo.app/ctpgo/evm/candidates"
	"ctpgo.app/ctpgo/evm/evmclient"
	"ctpgo.app/ctpgo/evm/pools"
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func installWiring(rpc bool) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.LUTC)

	evmclient.EthClientWiring(rpc)
	//runtime.GOMAXPROCS(16)
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	installWiring(len(os.Args) > 1 && os.Args[1] == "rpc")

	var allPools []*data.PoolData
	for _, e := range uclone.UCloneExchanges {
		epools := pools.GetPoolsForStdExchange(e)
		allPools = append(allPools, epools...)
	}
	empty := pools.GetEmptyPools(allPools)
	log.Println("Total pools", len(allPools), " empty: ", len(empty))
	candidates.CalculateCandidates(allPools, empty)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
