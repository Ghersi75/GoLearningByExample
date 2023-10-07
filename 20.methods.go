package main

import "fmt"

type rect struct {
	width, height int
}

// This essentially means that any rect object can now call area()
// Ex: r.area() can be called
func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2 * r.width + 2 * r.height
}

func Methods() {
	r := rect {
		width: 10,
		height: 5,
	}

	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perim())

	// Go handles dereferencing of pointers on its own
	rp := &r
	fmt.Println("Area:", rp.area())
	fmt.Println("Perimeter:", rp.perim())
}