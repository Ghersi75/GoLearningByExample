package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// The following structs will be used by the non-stateful goroutines to communicate with the stateful one
// These contain the important data that needs to go back and forth
type readOp struct {
	key		int
	resp	chan int
}

type writeOp struct {
	key		int
	val 	int
	resp	chan bool
}

// Stateful goroutines can be used instead of mutexes to keep track of state
// Apparently this "aligns with Go's idea of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine", much like a manager goroutine
func StatefulGoroutines() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	// stateful goroutine
	go func() {
		var state = make(map[int]int)
		// Loop that doesn't stop until the program stops running
		// Basically the state managing loop
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// reads
	for r := 0; r < 100; r++ {
		// 100 workers reading until the app stops running (roughly 1s)
		go func() {
			for {
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// writes
	for w := 0; w < 10; w++ {
		// 10 workers writing until the app stops running (roughly 1s)
		go func() {
			for {
				write := writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// give it a second
	time.Sleep(time.Second)

	// About 80k operations within a second
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

// Ultimately, choosing stateful goroutine vs mutex managed locking depends on the use case and how familiar the person is with it
// Use whatever you're most comfortable with, and if it doesn't make sense or it's too hard to use, try switching