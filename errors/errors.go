package main

import (
	"fmt"
	"errors"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}
	//A nil value in the error position indicates that there was no error.
	return arg + 3, nil
}

//A sentinel error is a predeclared variable that is used to signify a specific error condition.
var ErrOutOfTea = fmt.Errorf("No more tea available.")
var ErrPower = fmt.Errorf("can't boil water.")


//We can wrap errors with higher-level errors to add context.
//The simplest way to do this is with the %w verb in fmt.Errorf.
//Wrapped errors create a logical chain (A wraps B, which wraps C, etc.) that can be queried with functions like errors.Is and errors.As.
func makeTea(arg int) error {
	if (arg == 2) {
		return ErrOutOfTea
	} else if (arg == 4) {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed: ", e)
			} else {
			fmt.Println("f worked: ", r)
		}
	}
	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) { //errors.Is checks that a given error (or any error in its chain) matches a specific error value.
				fmt.Println("We shoud buy new tea !")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark")
				} else {
					fmt.Printf("unknown error: %s\n", err)
				}
				continue
			}
		fmt.Println("Tea is ready !")
	}
}