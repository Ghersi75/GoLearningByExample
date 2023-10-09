package main

import (
	"fmt"
	"errors"
)

// By convention, errors are the last returned argument and they are of type error
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg 	int
	prob 	string
}

// To create a custom error you simply define a .Error() func for the error struct
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{
			arg,
			"Can't work with it",
		}
	}

	return arg + 3, nil
}

func Errors() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}

		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// To work with custom error data you'll need to get the error, check that the error type matches the custom type, and then work with it like below
	_, e := f2(42)
	// e.(*argError) should be checking if the error object matches the argError type
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}