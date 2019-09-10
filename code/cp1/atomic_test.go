package cp1

import (
	"sync"
	"testing"
)

var total struct{
	sync.Mutex
	Value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i:=0; i <= 4; i ++ {
		total.Lock()
		total.Value += i
		total.Unlock()
	}
}

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	t.Log(total.Value)
}