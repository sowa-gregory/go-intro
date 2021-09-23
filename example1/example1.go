package main

import (
	"fmt"
)

func main() {
	a, b := 4, 7
	b, a = 2, 8

	r := fmt.Sprintf("%x%d", 58, 39)
	fmt.Println(a, b, r)

}
