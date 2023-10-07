package main

import (
	"fmt"
	"math"
)

// Cool blog post about interfaces I guess. Haven't read it, but don't wanna lose the link
// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

type geometry interface {
	area()	float64
	perim()	float64
}

// rect name conflicts with 20.methods.go declaration
type rect2 struct {
	width, height float64
}

type circle struct {
	radius float64
}

// To implement an interface, we simply need to implement all of its methods
// For example, below is the geometry interface implemented for the rect struct
func (r rect2) area() float64 {
	return r.width * r.height
}

func (r rect2) perim() float64 {
	return 2 * r.width + 2 * r.height
}

// Here's the implementation of the geometry interface for the circle struct
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("Geometry:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
}

func Interfaces() {
	r := rect2 {
		width: 3,
		height: 4,
	}

	c := circle {
		radius: 2,
	}

	measure(r)
	measure(c)
}