package main

import "fmt"

func RangeOverChannels() {
	queue := make(chan string, 2)
	queue <- "One"
	queue <- "Two"
	// Required to loop over channel with range
	close(queue)

	// Channels can be looped over with range, but queue must be closed
	for elem := range queue {
		fmt.Println(elem)
	}

}