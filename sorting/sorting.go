package main

import (
	"fmt"
	"slices"
)


/*Sorting functions are generic, and work for any ordered built-in type.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}
*/
func main() {
	strs := []string{"c", "a", "b"}
    slices.Sort(strs)
    fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
    slices.Sort(ints)
    fmt.Println("Ints:   ", ints)

	//We can also use the slices package to check if a slice is already in sorted order.
	s := slices.IsSorted(ints)
    fmt.Println("Sorted: ", s)
}