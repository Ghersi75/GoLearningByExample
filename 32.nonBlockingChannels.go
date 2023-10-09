package main

import "fmt"

// Adding a default to a switch causes it to be non blocking
func NonBlockingChannels() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <- messages:
		fmt.Println("Message received:", msg)
	default:
		fmt.Println("No messages received")
	}

	msg := "hi"
	// This commented code throws an error because it's not run in a goroutine and there's not channel buffer. One of the two must happen for it to run
	// messages <- msg
	select {
	// This doesn't send any message because there's not buffer for the messages channel
	// Since there is no buffer, the only messages the can be sent are from goroutines
	case messages <- msg:
		fmt.Println("Messaged sent:", msg)
	default:
		fmt.Println("No message sent")
	}

	// This message is never actually sent because default makes switch non blocking
	go func() {
		messages <- "msg"
	}()

	select {
	case msg := <- messages:
		fmt.Println("Received message:", msg)
	case sig := <- signals:
		fmt.Println("Signal received:", sig)
	default:
		fmt.Println("No activity")
	}
}