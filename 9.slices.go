package main

import (
	"fmt"
	"slices"
)

// Article by go devs on the design of slices
// https://go.dev/blog/slices-intro

func Slices() {
	// Declare slice. Doesn't need size and by default is nil and len 0
	// Slices can vary in size, much like arraylists in java for example
	// Arrays are fixed size
	var s []string
	fmt.Println("Uninitialized:", s, s == nil, len(s) == 0)

	// By default, capacity = len
	s = make([]string, 3)
	// To declare cap as well, 3rd param is cap
	// s = make([]string, 3, 6) gives it 3 len, so s = [0, 0, 0], but cap is 6
	// Cap is capacity
	// Arrays may have only 2 values (len) but have enough space allocated for 4 items (cap)
	fmt.Println("Empty:", s, "Len:", len(s), "Cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set:", s)
	fmt.Println("Get:", s[2])
	fmt.Println("Len:", len(s), "Cap:", cap(s))

	// Cap doubles each time when appending
	s = append(s, "d")
	fmt.Println("Added 1, Len:", len(s), "Cap:", cap(s))
	s = append(s, "e", "f")
	fmt.Println("Added 2, Len:", len(s), "Cap:", cap(s))
	s = append(s, "d")
	fmt.Println("Added 1, Len:", len(s), "Cap:", cap(s))

	c := make([]string, len(s))
	// Slices can be copied using copy
	copy(c, s)

	// var a []string
	var a = s
	// Since this returns a new slice, a is its own slice
	a = append(a, "a")
	a = a[:3]
	fmt.Println(a, s)

	var b = s
	s = append(s, "t")
	// Not the same, so b isn't a reference to s by default
	fmt.Println(b, s)

	var ref = &s
	s = append(s, "t")
	// Now it's the same since c is a reference to s
	fmt.Println(ref, s)

	// Use := only on first one since its being declared for the first time
	// [2:5] from 2 (included) to 5 (excluded) like python for example
	l := s[2:5]
	fmt.Println("Slice 1:", l)

	// Use = because variable has already been declared and is not just being modified
	l = s[:5]
	fmt.Println("Slice 2:", l)

	l = s[2:]
	fmt.Println("Slice 3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("Declared:", t)

	t2 := []string{"g", "h", "i"}
	// Same len, cap, and contents, so they are equal
	fmt.Println("t == t2", slices.Equal(t, t2))

	// Slices can also be used for 2D or multi dimensional arrays, but unlike usual multi dimensional arrays, the lengths of these can vary
	twoD := make([][]int, 3)
	fmt.Println("twoD empty: ", twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)
	// Output
	// [[0] [1 2] [2 3 4]]
	// Basically pyramid
	//   0
	//  1 2
	// 2 3 4
	// Pretty cool
}