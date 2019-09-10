package cp1

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)
type Config struct {
	Filename string
}

var config atomic.Value

func loadConfig() Config {
	// load
	return Config{Filename:"test.txt"}
}

func TestConfig(t *testing.T) {
	config.Store(loadConfig())

	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(loadConfig())
		}
	}()

	go func() {
		rs := []string{"a", "b", "c"}
		for _, val := range rs {
			c := config.Load()
			fmt.Printf("name = %s, config = %#v\n", val, c)
		}

	}()
	//<- time.After(time.Minute)
}
