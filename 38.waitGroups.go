package main

import (
	"fmt"
	"sync"
	"time"
)

func workerForWaitGroups(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroups() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Avoid using the same value for different goroutine
		i := i

		go func() {
			defer wg.Done()
			workerForWaitGroups(i)
		}()
	}

	// Basically synchronization but deals with in the Wait function
	wg.Wait()
}