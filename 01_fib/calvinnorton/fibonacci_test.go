package main

import (
	"bytes"
	"testing"
)

func TestFibonacciPositiveFlow(t *testing.T) {
	var buffer bytes.Buffer
	out = &buffer

	main()

	expected := `1
1
2
3
5
8
13
`
	actual := buffer.String()

	if expected != actual {
		t.Error("TestFibonacciPositiveFlow: Output not as expected. Received:", actual, "Expected:", expected)
	}
}

func TestFibonacciMaxInt(t *testing.T) {
	var buffer bytes.Buffer
	out = &buffer
	fib(93)
	expected := "Value too high, please choose a value less of 92 or less."
	actual := buffer.String()

	if expected != actual {
		t.Error("TestFibonacciMaxInt: Output not as expected. Received:", actual, "Expected:", expected)
	}
}

func TestFibonacciZeroValue(t *testing.T) {
	var buffer bytes.Buffer
	out = &buffer
	fib(0)
	expected := ""
	actual := buffer.String()

	if expected != actual {
		t.Error("TestFibonacciZeroValue: Output not as expected. Received:", actual, "Expected:", expected)
	}
}

func TestFibonacciNegativeValue(t *testing.T) {
	var buffer bytes.Buffer
	out = &buffer
	fib(-1)
	expected := "Value too low, please choose a positive integer."
	actual := buffer.String()

	if expected != actual {
		t.Error("TestFibonacciNegativeValue: Output not as expected. Received:", actual, "Expected:", expected)
	}
}
