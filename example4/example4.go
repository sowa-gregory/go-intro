package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, i int) {
	sec := rand.Intn(10)
	time.Sleep(time.Second * time.Duration(sec))
	fmt.Println(i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
}
