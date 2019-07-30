package main

import (
	"fmt"
	//"io"
	//"os"
)

func fib(n int) {
	previous, current := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println(current)
		temp := current
		current = previous + current
		previous = temp
	}
}

func main() {
	fib(7)
}
