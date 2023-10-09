package main

import "fmt"

func ChannelBuffering() {
	// By default channel isnt buffered, second parameter can be used to defined buffer size
	messages := make(chan string)

	go func() { 
		messages <- "buffered" 
	}()
	go func() { 
		messages <- "channel" 
	}()
	go func() { 
		messages <- "extra string" 
	}()

	// All 3 of the above will print since they use goroutines
	fmt.Println(<- messages)
	fmt.Println(<- messages)
	fmt.Println(<- messages)

	// Buffering is important for non goroutine channel accesses
	msgs := make(chan string, 3)

	msgs <- "1"
	msgs <- "2"
	msgs <- "3"
	// Attempting to store a 4th value with a 3 buffer size will throw an error
	// messages <- "4"

	fmt.Println(<- msgs)
	fmt.Println(<- msgs)
	fmt.Println(<- msgs)

}