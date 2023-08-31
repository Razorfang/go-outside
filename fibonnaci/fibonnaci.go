package main

import (
	"fmt"
)

func fibonnaci(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fibonnaci(n-1) + fibonnaci(n-2)
	}
}

func main() {
	for i := 0 ; i < 20 ; i++ {
		fmt.Printf("fibonnaci(%d) = %d\n", i, fibonnaci(i))
	}
}
