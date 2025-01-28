package main

import(
	"fmt"
	"maps"
)

func main(){
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map = ", m)
	//get a val from the map
	v1 := m["k1"]
	fmt.Println("v1 = ", v1)
	//if a value doesnt exist, a zero is returned
	v3 := m["k3"]
	fmt.Println("v3 = ", v3)
	
	fmt.Println("len = ", len(m))
	
	//delete a pair of key/value
	delete(m, "k2")
	fmt.Println("map = ", m)
	
	//clear totality of the map
	clear(m)
	fmt.Println("map = ", m)
	
	//The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	// Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs = ", prs)
	
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map = ", n)
	
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2){
		fmt.Println("n == n2")
	}
}