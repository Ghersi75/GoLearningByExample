package main

import "fmt"

// Closing channels comes in handy to know when all jobs have been completed
func ClosingChannels() {
	// Buffer is optional here, but higher buffer will allow more jobs to be sent before a receiver HAS to act
	// No buffer forces 2 jobs to be sent at once and wait for worker to receive them
	// Buffer 10 allows all 10 jobs to be sent before worker even starts receiving jobs
	// This is likely different based on the os or some variable of that sort since output may be different each time its run
	jobs := make(chan int, 10)
	// Used for synchronization from synchronization code (28)
	done := make(chan bool)

	go func() {
		for {
			// When a channel is closed, more will return false, else true
			j, more := <- jobs
			if more {
				fmt.Println("Job received:", j)
			} else {
				fmt.Println("All jobs received")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 10; j++ {
		jobs <- j
		fmt.Println("Sent job:", j)
	}
	close(jobs)
	fmt.Println("All jobs sent")

	// Wait until done has a value
	// Synchronization from previous code (28)
	<- done
}