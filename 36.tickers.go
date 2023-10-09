package main

import (
	"fmt"
	"time"
)

func Tickers() {
	// Every 500 seconds a new value is added to the channel, thus creating a 500ms ticker
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at:", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	// Stop the ticker from sending any more signals
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}