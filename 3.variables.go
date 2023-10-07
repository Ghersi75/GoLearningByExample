package main

import "fmt"

func Variables() {
	var a = "initial"
	fmt.Println(a)

	// Type goes after I guess
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// Shorthand for declaring and initializing
	// var f string = "apple"
	f := "apple"
	fmt.Println(f)

}