package main

import "fmt"

// Channels are basically pipes that allow communication across goroutines
func Channels() {
	// Create a channel by using make(chan typeOfChannel)
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <- messages
	fmt.Println(msg)
}
// Interestingly, if there's no goroutines here, errors come up about all goroutines being asleep - deadlock
// Not sure why this is the case, but it seems like channels need goroutines to work properly, understandably