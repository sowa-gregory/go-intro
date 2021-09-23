package main

import "fmt"

func main() {
	var arr []int

	for i := 0; i < 300; i++ {
		arr = append(arr, i)
		fmt.Printf("%p\n", &arr)
	}

}
