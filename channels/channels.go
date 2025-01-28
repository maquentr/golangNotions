package main

import (
    "fmt"
    "math/rand"
    "time"
)

//Channels are the pipes that connect concurrent goroutines. 
//You can send values into channels from one goroutine and receive those values into another goroutine.

func GenerateRandomNumber(n int) int {
    rand.Seed(time.Now().UnixNano()) //permet de ne pas avoir le meme rand a chaque appel
    randomNumb := rand.Intn(n)
    return randomNumb
}

func CalculateValue(intChan chan int) {
    randomNumber := GenerateRandomNumber(50) // de 0 a 50
    intChan <- randomNumber
}

//Create a new channel with make(chan val-type). Channels are typed by the values they convey.
func main() {
    foo := make(chan int)
    defer close(foo) //permet de close le channel une fois son utilisation terminee. inutile dans le main ofc mais c'est pour l'exemple

    go CalculateValue(foo)

    num := <-foo
    fmt.Println(num)


    //another example
    messages := make(chan string)

    //Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
    go func() { messages <- "ping" }()

    //The <-channel syntax receives a value from the channel. 
    //Here we’ll receive the "ping" message we sent above and print it out
    msg := <-messages
    fmt.Println(msg)
}
