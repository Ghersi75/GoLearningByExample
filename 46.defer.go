package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("Creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// Defer is very much like a `finally` in other languages
// It is used to push the execution of a function to the end of the current function's call, but without having to put that call at the end
func Defer() {
	f := createFile("46.defer.txt")
	defer closeFile(f)
	writeFile(f)

	// Output is creating -> writing -> closing because defer pushing closeFile to the end of the function call
}