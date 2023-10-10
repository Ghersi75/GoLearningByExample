package main

import "fmt"

func myPanic() {
	panic("A problem I guess")
}

// There's a recover function which helps get out of a panic if necessary
// Simply call recover() and see if returns something, if it does, that's the panic error
func Recover() {
	// This will execute at the end
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	myPanic()

	fmt.Println("After myPanic")
}