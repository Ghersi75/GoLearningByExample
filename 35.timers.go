package main

import (
	"fmt"
	"time"
)

func Timers() {
	timer1 := time.NewTimer(2 * time.Second)
	// Wait for timer 1 to finish
	<-timer1.C
	fmt.Println("Timer 1 finished")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// <-timer2.C
	// If <-timer2.C isn't called, blocking and waiting until the timer stops, the following will stop the timer instead and move on
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	} else {
		fmt.Println("Timer 2 finished")
	}

	// Give it a couple seconds to at least try to fire timer 2, though unlikely
	time.Sleep(2 * time.Second)
}