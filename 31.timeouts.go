package main

import (
	"fmt"
	"time"
)

func Timeouts() {
	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	
	// c1 takes 2 seconds to run
	// time.After(1 * time.Second) takes 1 second to run
	// Since select picks the first option available and moves out of the select, the timeout placed at the last option would move on if the other results havent shown up yet
	// This effectively is a timeout that limits c1's result to taking at most 1 second, after which it is skipped
	select {
	case res := <- c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "res 2"
	}()

	// It's important to know that select only selects the first value to match
	// In the case below, despite c2 being set after 2 seconds and the timeout taking 3 seconds, the timeout case is never reached since there was a successful case before it
	select {
	case res := <- c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}