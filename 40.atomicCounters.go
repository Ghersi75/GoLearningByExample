package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCounters() {
	// Int that can be modifier by all goroutines without sync issues
	var ops atomic.Uint64
	var wg sync.WaitGroup

	// 50 workers will add 1000 to the atomic int each
	for i := 0; i < 50; i++ {
		// add to wait group
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}

			// finish each wait group
			wg.Done()
		}()
	}

	// wait on all tasks to finish
	wg.Wait()
	// print final count 50000
	fmt.Println("ops:", ops.Load())
}