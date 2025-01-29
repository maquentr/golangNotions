package main

import "fmt"

//for and range provide iteration over basic data structures. 
//We can also use this syntax to iterate over values received from a channel.

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println("elem: ", elem)
	}
}