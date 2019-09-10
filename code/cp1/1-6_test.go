package cp1

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSync(t *testing.T) {
	done := make(chan int)
	go func() {
		fmt.Println("hello, world")
		<-done
	}()
	done <- 1
}

func TestSyncCacheChannel(t *testing.T) {
	done := make(chan int, 10)
	for i := 0; i < cap(done); i++ {
		go func(int2 int) {
			rand.Seed(time.Now().UnixNano())
			r := rand.Intn(10)
			time.Sleep(time.Millisecond * time.Duration(r))
			fmt.Printf("hello, world to %d\n", int2)
			done <- 1
		}(i)
	}
	for i:= 0; i < cap(done); i++ {
		<- done
	}
}
