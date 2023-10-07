package main

import "fmt"

func For() {
	i := 1
	for i <= 3 {
		fmt.Println("i: ", i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println("j: ", j)
	}

	for {
		fmt.Println("\"While\" loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println("n: ", n)
	}
}