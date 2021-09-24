package main

import (
	"fmt"
)

func worker(ch chan<- uint64, max int) {
	a, b := uint64(0), uint64(1)
	for i := 0; i < max; i++ {
		b, a = a+b, b
	}

	ch <- a
}

func main() {
	ch := make(chan uint64)
	numWorkers := 1000

	for i := 0; i < 1000; i++ {
		go worker(ch, i%100)
	}

	for i := 0; i < numWorkers; i++ {
		fmt.Println(<-ch)
	}
}
