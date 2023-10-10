package main

import "os"

func Panic() {
	// Basically just throw an error
	panic("A problem I guess???????????")

	// This will never be reached
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}