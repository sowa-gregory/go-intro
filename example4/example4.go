package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, i int) {
	fmt.Println(i)
	//---------
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	for i := 0; i < 10; i++ {
		go worker(&wg, i)
	}
	wg.Wait()
}
