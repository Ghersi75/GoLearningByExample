package main

import "fmt"

// Just a generic, pretty much like typescript T of whatever kind
// comparable just makes sure that the type has == and != which is required for map keys
// any is an alias for interface{}, which is better than the weird interface{} syntax
func MapKeys[K comparable, V any] (m map[K]V) []K {
	// Take a map of any type and return a slice of its keys
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Linked list basic struct for head and tail
type List[T any] struct { 
	// doing [T]element would mean an array of length T which makes no sense
	// generic definitions always go after the object, hence why element[T]
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

// v is of generic type T here, and so is the List
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{
			val: v,
		}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{
			val: v,
		}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func Generics() {
	var m = map[int]string{
		1: "2",
		2: "4",
		4: "8",
	}

	// Give it the generic type and it works
	fmt.Println("Keys:", MapKeys[int, string](m))
	
	// Don't give a type and it will figure it out
	fmt.Println("Keys:", MapKeys(m))

	lst := List[int]{ }
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("List:", lst.GetAll())
}