package main

import (
	"fmt"
	"time"
)

// 生产者消费者

func Producer(factor int, out chan<- int) {
	for i:=0; ; i++ {
		out <- i*factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)

	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)

	time.Sleep(time.Second*5)
}