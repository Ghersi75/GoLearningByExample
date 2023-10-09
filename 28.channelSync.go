package main

import (
	"fmt"
	"time"
)

// Can pass a channel to a func as a parameter
func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ChannelSync() {
	done := make(chan bool)
	go worker(done)
	// Adding more functions doesn't work as expected as we don't wait for all the workers to finish
	// To do so, it's better to use WaitGroup, which will be covered later
	// go worker(done)
	// go worker(done)
	// go worker(done)
	// go worker(done)
	// go worker(done)

	// Blocking operation waiting for a response from the worker
	// Without this the program would not wait on the goroutine
	<- done
}