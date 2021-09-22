package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker1(ch chan<- string) {
	sec := rand.Intn(20)
	time.Sleep(time.Second * time.Duration(sec))
	fmt.Println("worker1 done")
	ch <- "hi from worker1"
}

func worker2(ch chan<- int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(20)))
	fmt.Println("worker2 done")
	ch <- 250
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch1 := make(chan string)
	ch2 := make(chan int)

	go worker1(ch1)
	go worker2(ch2)

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)

}

/*
for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)

		}
		fmt.Println("...")
	}
*/

/*
for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		default:

			fmt.Println("...")
		}
	}
*/

/*
for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		default:
			fmt.Println("...")
		}
		time.Sleep(time.Second * 1)
	}
*/

/*
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
*/
