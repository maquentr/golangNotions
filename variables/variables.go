package main

import "fmt"

func main(){
	var a = "initial"
	fmt.Println(a)
	var b,c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int // set to zero by default
	fmt.Println(e)
	f := "apple" //auto select type by its value
	fmt.Println(f)


}