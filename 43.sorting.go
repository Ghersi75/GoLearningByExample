package main

import (
	"fmt"
	"slices"
)

func Sorting() {
	strs := []string{"a", "c", "d", "b"}
	slices.Sort(strs)
	fmt.Println("Sorted strs:", strs)

	ints := []int{7, 2, 4, 19, 1}
	slices.Sort(ints)
	fmt.Println("Sorted ints:", ints)

	s := slices.IsSorted(ints)
	fmt.Println("IsSorted(ints):", s)
}