package main

import "fmt"

func Range() {
	nums := []int{2, 3, 4}
	sum := 0
	// First return value is index
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index of 3:", i)
		}
	}

	kvs := map[string][]string {
		// Arrays being declared with {} instead of [] is weird ngl
		"a": {"apple", "ananas"},
		"b": {"banana", "baloon"},
	}
	// Loop over key/value
	for k, v := range kvs {
		fmt.Printf("%s -> %v\n", k, v)
	}

	// Can also just loop over keys only
	for k := range kvs {
		fmt.Println("Key:", k)
	}

	for i, c := range "go" {
		// Not sure what runes are just yet, but it seems to return the utf-8 code rather than a string itself for c
		fmt.Println("Index:", i, "Char:", c)
	}

	for i, c := range "古い世界" {
		// Indexes here aren't 0 1 2 3 because these characters are utf-8 encoded and actually take 3 bytes each
		// Indexes are 0 3 6 9
		fmt.Println("Index:", i, "Char:", c)
	}

	// More about runes in a later section I guess
}