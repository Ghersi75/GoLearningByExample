package main

import (
	"fmt"
	"maps"
)


// Basically like any other hashmap
func Maps() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("Map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Since it's not found, it will give it the default value of the return type
	// In this case, since the type is int, v3 = 0
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// Returns the number of key/value pairs
	fmt.Println("Len:", len(m))

	// If the key doesn't exist, not changes are made
	delete(m, "k2")
	fmt.Println("Map:", m, "Len:", len(m))

	// Removes all key/value pairs
	clear(m)
	fmt.Println("Map:", m, "Len:", len(m))

	// First return value is the value returned
	// Second return value is whether that key returned anything or not
	// Used to distinguish between a 0 that's supposed to be there and a 0 returned because the key wasn't found
	_, prs := m["k2"]
	// prs = Present, aka if that key actually returned anything and a value is present
	fmt.Println("prs:", prs)

	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	n2 := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	fmt.Println("n:", n, "\nn2:", n2, "\nn == n2:", maps.Equal(n, n2))
}