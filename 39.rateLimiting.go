package main

import (
	"fmt"
	"time"
)

func RateLimiting() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Basically NewTicker but just the channel
	limiter := time.Tick(500 * time.Millisecond)

	for req := range requests {
		// Waits 200ms
		<-limiter
		fmt.Println("Request", req, time.Now())
	}

	// Allows processing of 3 requests at a time 
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		// The first 3 requests will be done right away, and the rest will have 200ms delay each
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("Request", req, time.Now())
	}
}