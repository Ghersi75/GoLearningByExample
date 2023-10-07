package main

import "fmt"

// As simple as putting multiple return types for the values
func vals() (int, int) {
	return 3, 7
}

func MultipleReturns() {
	a, b := vals()
	fmt.Println("a:", a, "b:", b)

	_, c := vals()
	fmt.Println("c:", c)
}