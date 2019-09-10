package cp1

import (
	"sync"
	"sync/atomic"
)

type singleton struct{}

var (
	instance  *singleton
	initialed uint32
	mu        sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialed) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialed, 1)
		instance = &singleton{}
	}
	return instance
}
