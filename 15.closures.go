package main

import "fmt"

// Return type is a func() which itself returns an int
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Closures() {
	nextInt := intSeq()

	fmt.Println("NextInt:", nextInt())
	fmt.Println("NextInt:", nextInt())
	fmt.Println("NextInt:", nextInt())

	newInt := intSeq()

	fmt.Println("NewInt:", newInt())
}