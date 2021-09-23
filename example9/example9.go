package main

import (
	"fmt"
	"math/rand"
	"time"
)

func workerString(ch chan<- string) {
	sec := rand.Intn(10)
	time.Sleep(time.Second * time.Duration(sec))
	fmt.Println("worker1 done")
	ch <- "hi from worker1"
}

func workerInt(ch chan<- int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	fmt.Println("worker2 done")
	ch <- 250
}

func start1() {
	rand.Seed(time.Now().UnixNano())
	ch1 := make(chan string)
	ch2 := make(chan int)

	go workerString(ch1)
	go workerInt(ch2)

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}

func start2() {
	rand.Seed(time.Now().UnixNano())
	ch1 := make(chan string)
	ch2 := make(chan int)

	go workerString(ch1)
	go workerInt(ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
		fmt.Println("...")
	}
}

func start3() {
	rand.Seed(time.Now().UnixNano())
	ch1 := make(chan string)
	ch2 := make(chan int)

	go workerString(ch1)
	go workerInt(ch2)

	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		default:
			fmt.Println("...")
		}
		time.Sleep(time.Second * 2)
	}
}

func start4() {
	rand.Seed(time.Now().UnixNano())
	ch1 := make(chan string)
	ch2 := make(chan int)

	go workerString(ch1)
	go workerInt(ch2)

	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		case <-time.After(time.Second * 1):
			fmt.Println("...")
		}
	}
}

func main() {
	start1()
}
