package main

import "fmt"

// chan<- means this channel can only be used to send messages
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// <-chan means it can only be used to receive messages
// chan<- means it can only be used to send messages
func pong(pings <-chan string, pongs chan<- string) {
	msg := <- pings 
	pongs <- msg
}

func ChannelDirections() {
	// For non goroutine calls, a buffer is required, otherwise an error will be thrown
	pings := make(chan string)
	pongs := make(chan string)
	// Order of goroutines doesn't matter here since the receives and sends are blocking and will wait for the data
	go ping(pings, "passed message")
	go pong(pings, pongs)
	fmt.Println(<- pongs)
}	