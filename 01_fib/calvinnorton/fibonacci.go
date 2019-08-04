package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	previous, current := 0, 1
	if n > 92 {
		fmt.Fprint(out, "Value too high, please choose a value less of 92 or less.")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, current)
		valueHolder := current
		current = previous + current
		previous = valueHolder
	}
}

func main() {
	fib(7)
}