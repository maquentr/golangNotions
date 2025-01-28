package main

import "fmt"

//generics permettent de recevoir nimporte quel type en parametre et de faire donctionner la fonction selon les types recus


//SlicesIndex takes a slice of any comparable type and an element of that type and returns the index of the first occurrence of v in s, or -1 if not present.
//The comparable constraint means that we can compare values of this type with the == and != operators.
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
    for i := range s{
        if v == s[i]{
            return i
        }
    }
    return -1
}

type List[T any] struct{
    head, tail *element[T]
}

type element[T any] struct{
    next *element[T]
    val T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func main() {
    var s = []string{"foo", "bar", "zoo"}

    fmt.Println("index of zoo ", SlicesIndex(s, "zoo"))


    _ = SlicesIndex[[]string, string](s, "zoo")

    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list: ", lst.AllElements())
    lst2 := List[string]{}
    lst2.Push("a")
    lst2.Push("b")
    lst2.Push("c")
    fmt.Println("list: ", lst2.AllElements())


    fmt.Println(min[float32](0.5, 0.7))
    fmt.Println(min[int32](1, 2))
    type f float64
    var foo f = 3.14
    fmt.Println(min(foo, 0.7)) //ne marche pas car float64 est le sous type de f
    //il faut alors mettre un tilde (~) dans les types acceptes de la fonction pour accepter les sous types
    //on peut mettre le mot cle ANY pour accepter nimporte quel type ou creer une interface avec les types acceptes
}
type Number interface {
    int8 | int16 | int32 | int | float | float32 | ~float64
}

func min[T int32 | float32](x, y T) T {
    if x < y {
        return x
    } else {
        return y
    }
}

func min[T ~float64](x, y T) T {
    if x < y {
        return x
    } else {
        return y
    }
}

func min[T Number](x, y T) T {
    if x < y {
        return x
    } else {
        return y
    }
}