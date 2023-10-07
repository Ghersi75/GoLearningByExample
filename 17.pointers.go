package main

import "fmt"

// Pass by value
// Copy made
func zeroval(ival int) {
	ival = 0
}

// Pass pointer
func zeroptr (iptr *int) {
	// Dereference pointer and change value
	*iptr = 0
}

func Pointers() {
	i := 1
	fmt.Println("Initial:", i)

	zeroval(i)
	fmt.Println("After zeroval:", i)

	// & to get pointer, like in C
	zeroptr(&i)
	fmt.Println("After zeroptr:", i)

	fmt.Println("Pointer:", &i)

}