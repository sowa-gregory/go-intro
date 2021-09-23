package main

import (
	"fmt"
	"os"
	"sync"
)

func worker1(wg *sync.WaitGroup, num int) {
	name := fmt.Sprintf("/tmp/file%d", num)

	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		wg.Done()
	}
	for i := 0; i < 100; i++ {
		_, err = file.WriteString(fmt.Sprintf("line%d\n", i))
		if err != nil {
			fmt.Println(err)
			file.Close()
			wg.Done()
		}
	}
	_, err = file.WriteString("end\n")
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	wg.Done()
}

/*
func worker2(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	name := fmt.Sprintf("/tmp/file%d", num)
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	for i := 0; i < 100; i++ {
		_, err = file.WriteString(fmt.Sprintf("line%d\n", i))
		if err != nil {
			fmt.Println(err)
		}
	}
	_, err = file.WriteString("end\n")
	if err != nil {
		fmt.Println(err)
	}
}*/

func main() {
	var wg sync.WaitGroup

	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker1(&wg, i)
	}
	wg.Wait()

}
