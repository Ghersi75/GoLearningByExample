package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func Recursion() {
	fmt.Println("7!:", fact(7))

	// Closures can also be recursive, but they NEED to be declared as var
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n - 1) + fib(n - 2)
	}

	fmt.Println("Fib 25:", fib(25))
}