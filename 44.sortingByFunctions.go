package main

import (
	"cmp"
	"fmt"
	"slices"
)

func SortingByFunctions() {
	fruits := []string{"peach", "banana", "kiwi"}

	// Compare by length function
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println("Sorted fruits by length:", fruits)

	type Person struct {
		name string
		age int
	}

	people := []Person{
		Person{name: "Jack", age: 75},
		Person{name: "Herbert", age: 23},
		Person{name: "Zack", age: 45},
	}

	// Sort ny age
	slices.SortFunc(people, 
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		},
	)
	
	fmt.Println("Sorted people by age:", people)

}