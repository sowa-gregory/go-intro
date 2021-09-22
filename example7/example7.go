package main

import (
	"fmt"
	"time"
)

func worker(ch chan<- uint64, max int) {
	a, b := uint64(0), uint64(1)
	for i := 0; i < max; i++ {
		a, b = b, a+b
	}
	time.Sleep(time.Second * 1)
	ch <- a
}

func main() {
	ch := make(chan uint64)
	numWorkers := 100

	for i := 0; i < numWorkers; i++ {
		go worker(ch, i)
	}

	for i := 0; i < numWorkers; i++ {
		fmt.Println(<-ch)
	}
}
