package main

import (
	"fmt"
	"math"
)

const testConstString string = "constant"

func Constants() {
	fmt.Println(testConstString)

	// Declaring this as an int would give errors since it would know its casting to other values
	// Not declaring its type allows go to infer it when it does computations such as Sin, which expects float64
	// Automatic casting only happens if no type is declared
	const n = 500_000_000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}