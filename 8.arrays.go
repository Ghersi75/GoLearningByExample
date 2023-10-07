package main

import "fmt"

func Arrays() {
	// Why []int and not int[] 
	// Weird
	var a [5]int
	// [0, 0, 0, 0, 0]
	// All unassigned vars get automatically assigned 0 values, such as 0, false, "", [0] etc
	fmt.Println("Empty: ", a)

	a[4] = 100
	fmt.Println("Set: ", a)
	fmt.Println("Get: ", a[4])
	fmt.Println("Length: ", len(a))

	b := [5]int {1, 2, 3, 4, 5}
	fmt.Println("Declared: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}