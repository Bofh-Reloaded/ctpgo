package candidates

import (
	"ctpgo.app/ctpgo/core/common"
	"ctpgo.app/ctpgo/core/data"
	"github.com/gammazero/deque"
	"github.com/panjf2000/ants/v2"
	"log"
	"sync"
	"time"
)

var startingTokens = []common.Address{
	common.AsAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"),
	common.AsAddress("0xe9e7cea3dedca5984780bafc599bd69add087d56"),
	common.AsAddress("0x55d398326f99059ff775485246999027b3197955"),
	common.AsAddress("0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82"),
	common.AsAddress("0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c"),
	common.AsAddress("0x2170ed0880ac9a755fd29b2688956bd959f933f8"),
}

type CandidateEntry struct {
	sealed bool
	path   []*data.PoolData
	tokens []*data.TokenData
	next   *data.TokenData
}

func newCandidateEntry(next *data.TokenData) *CandidateEntry {
	path := make([]*data.PoolData, 0, 4)
	tokens := make([]*data.TokenData, 0, 4)
	return &CandidateEntry{
		path:   path,
		tokens: tokens,
		sealed: false,
		next:   next,
	}
}

func (this *CandidateEntry) addPool(pool *data.PoolData, t2i map[common.Address]*data.TokenData) {
	next := pool.Token0

	if next == this.next.Address {
		next = pool.Token1
	}
	this.tokens = append(this.tokens, this.next)
	this.path = append(this.path, pool)
	this.next = t2i[next]
	this.sealed = next == this.tokens[0].Address
}

func (this *CandidateEntry) cloneWithPool(pool *data.PoolData, t2i map[common.Address]*data.TokenData) *CandidateEntry {
	path := make([]*data.PoolData, len(this.path), cap(this.path))
	tokens := make([]*data.TokenData, len(this.tokens), cap(this.tokens))
	copy(path, this.path)
	copy(tokens, this.tokens)

	instance := &CandidateEntry{
		sealed: this.sealed,
		path:   path,
		tokens: tokens,
		next:   this.next,
	}
	instance.addPool(pool, t2i)
	return instance
}

func CalculateCandidates(pools []*data.PoolData, emptyPools map[common.Address]bool) []*data.SimpleExecutionPath {
	min := 2
	max := 3
	log.Println("Searching candidates betwewn:", min, "and", max)
	i := uint64(0)
	t2i := make(map[common.Address]*data.TokenData)
	t2Pool := make(map[uint64][]*data.PoolData)
	for _, p := range pools {
		// If the pool is considered "empty" skip it
		if _, f := emptyPools[p.Address]; f {
			continue
		}

		// Calculate the indices - we do not want to use strings...
		var t0i, t1i uint64
		if t0idx, t0f := t2i[p.Token0]; t0f {
			t0i = t0idx.Id
		} else {
			t2i[p.Token0] = &data.TokenData{
				Id:      i,
				Address: p.Token0,
			}
			t0i = i
			i += 1
		}
		if t1idx, t1f := t2i[p.Token1]; t1f {
			t1i = t1idx.Id
		} else {
			t2i[p.Token1] = &data.TokenData{
				Id:      i,
				Address: p.Token1,
			}
			t1i = i
			i += 1
		}

		// Add the pool to the relevant slices
		if ex, found := t2Pool[t0i]; found {
			t2Pool[t0i] = append(ex, p)
		} else {
			e := make([]*data.PoolData, 0)
			t2Pool[t0i] = append(e, p)
		}
		if ex, found := t2Pool[t1i]; found {
			t2Pool[t1i] = append(ex, p)
		} else {
			e := make([]*data.PoolData, 0)
			t2Pool[t1i] = append(e, p)
		}
	}
	log.Println("t2Pool Map created")

	c := make([]*CandidateEntry, 0)
	results := make(chan *data.SimpleExecutionPath, 100000000)
	for _, addr := range startingTokens {
		i := t2i[addr]
		if stPools, found := t2Pool[i.Id]; found {
			for _, p := range stPools {
				entry := newCandidateEntry(i)
				entry.addPool(p, t2i)
				c = append(c, entry)
			}
		}
	}
	var wg sync.WaitGroup
	var rg sync.WaitGroup

	seps := make([]*data.SimpleExecutionPath, 0, 10000000)
	wp, _ := ants.NewPoolWithFunc(common.MAX_GOROUTINES, func(i interface{}) {
		stepSingle(i.(*CandidateEntry), t2i, t2Pool, min, max, results)
		defer wg.Done()
	})

	rg.Add(1)
	adv := 0
	go func() {
		for r := range results {
			seps = append(seps, r)
			if len(seps)%10000 == 0 && len(seps) > adv {
				adv = len(seps)
				log.Println("Seps", "len", len(seps))
			}
		}
		defer rg.Done()
	}()

	log.Println("1-Way", len(c))

	for _, it := range c {
		wg.Add(1)
		_ = wp.Invoke(it)
	}
	wg.Wait()
	log.Println("Done...")
	time.Sleep(4000 * time.Millisecond)
	close(results)
	rg.Wait()
	log.Println("Found ", len(seps), "candidates")
	wp.Release()
	return seps
}

func stepSingle(entry *CandidateEntry, t2i map[common.Address]*data.TokenData, t2Pool map[uint64][]*data.PoolData, min int, max int, results chan *data.SimpleExecutionPath) {
	//start := time.Now()
	cs := deque.New(256, 32)

	cs.PushBack(entry)
	for cs.Len() != 0 {
		it := cs.PopBack().(*CandidateEntry)
		if ex, found := t2Pool[it.next.Id]; found {
			for _, p := range ex {
				e := it.cloneWithPool(p, t2i)
				if e.sealed {
					if len(e.path) >= min {
						results <- data.NewSimpleExecutionPath(e.path, e.tokens)
					}
				} else if len(e.path) < max {
					cs.PushBack(e)
				}
			}
		}
	}
	//log.Println("Perf:", i, " -> ", (time.Now().Sub(start)))
}
