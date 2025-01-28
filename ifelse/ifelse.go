package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("is even")
	} else {
		fmt.Println("is odd")
	}

	if num := 9; num < 0 {
		fmt.Println("is negative")
	} else if num < 10 {
		fmt.Println("has 1 digit")
	} else {
		fmt.Println(num, "has multiple digit")
	}
}
