package main

import "fmt"

type person struct {
	name 	string
	age 	int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func Structs() {
	// This apparently creates a new struct, but don't know what that means exactly
	// I guess it just means that it will create a person as expected without issues here
	// Probably smart not to do it this way
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	// This is how things should be done
	// Let a function handle error checking and creation of new object
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	// Fields can be accessed with .
	fmt.Println(s.name)

	sp := &s
	// Fields can be accessed with . on pointers as well since they are automatically dereferenced
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// Anonymous structs (not named or typed as their own thing) can also be created
	dog := struct {
		name		string
		isGood 	bool
	} {
		"Rex",
		true,
	}

	fmt.Println(dog)
}	