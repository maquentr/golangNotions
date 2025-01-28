package main

import "fmt"

//closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	//fonctions anonymes
	w := func() {
		fmt.Println("je suis une fonction anonyme sans return")
	}
	
	z := func() string {
		fmt.Println("je suis une fonction anonyme avec return")
		return "Alex"
	}() //si les parentheses sont la elles ne doivent pas etre appelees.... 

	w()
	fmt.Println(z) // .....ici ?????
	
	//mettre les parentheses a la fin des scopes, la variable reste variable et non fonction
	x, y := func() (int, int) {
		fmt.Println("aucun parametre. deux return")
		return 5, 7
	}()
	fmt.Println(x)
	fmt.Println(y)

	//on peut appeler c et y en param dune fonction
	func(a,b int) {
		fmt.Println("A * A + B * B = ", a*a+b*b)
	}(x, y)


	//We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
	nextInt := intSeq()
	
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	
	newtInt := intSeq()
	fmt.Println(newtInt())
}

