package main

import "fmt"

type base struct {
	num	int
}

func (b base) describe() string {
	return fmt.Sprintf("Base with num = %v", b.num)
}

type container struct {
	base
	str 	string
}

func StructEmbedding() {
	co := container {
		base: base {
			num: 1,
		},
		str: "Container 1",
	}

	// Don't need to do co.base.num, co.num works just fine
	// Neat feature ngl
	fmt.Printf("co num = %v str: %v\n", co.num, co.str)
	fmt.Printf("also num = %v\n", co.base.num)
	
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("Describer:", d.describe())
}

