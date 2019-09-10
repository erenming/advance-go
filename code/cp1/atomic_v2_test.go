package cp1

import (
	"sync"
	"sync/atomic"
	"testing"
)

var total2 uint64

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total2, i)
	}
}

func TestAtomicV2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker2(&wg)
	go worker2(&wg)
	wg.Wait()
	t.Log(total2)
}