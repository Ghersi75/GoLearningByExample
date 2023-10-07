package main

import "fmt"

func sum(nums ...int) {
	fmt.Print("Nums: ", nums)
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(" Sum: ", total)
}

func VariadicFunctions() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// nums... basically turns the slice into a valid input for a variadic function
	sum(nums...)
}