package main

import "fmt"


// Param types go after var name
// Func return type goes after params
func plus(a int, b int) int {
	return a + b
}

// If a group of params are of same type, only last one needs the type
func plusPlus(a, b, c int, d string) int {
	fmt.Println("d:", d)
	return a + b + c
}

func Functions() {
	res := plus(1, 2)
	fmt.Println("Res:", res)

	res = plusPlus(1, 2, 3, "Test")
	fmt.Println("Res:", res)
}