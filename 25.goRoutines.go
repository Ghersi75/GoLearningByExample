package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}


// For goroutines to work, the program needs some way to wait for them
// Since they run independently from other things, they will wait until they can run, but the program will not automatically wait on them
// The time.Sleep here makes sure there's enough time for them to execute
func GoRoutines() {
	// Blocking function call
	f("direct")

	// Async/Non blocking
	go f("goroutine")


	go func(msg string) {
		fmt.Println(msg)
	} ("going")

	time.Sleep(time.Second)
	fmt.Println("Done")
}