package main

import (
	"fmt"
	"time"
)

func Select() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// Select allows waiting for multiple channels at the same time
		// using time run go *.go shows an actual time of about 2.4s instead of 3s since both goroutines sleep concurrently
		select{
		case msg1 := <- c1:
			fmt.Println("received", msg1)
		case msg2 := <- c2:
			fmt.Println("received", msg2)
		}
	}
}