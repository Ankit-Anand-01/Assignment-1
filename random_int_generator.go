//generate random integer in Go
package main

//import "math/rand" to generate random int number

import (
	"fmt"
	"math/rand"
)

func main() {
	//generate a random number
	message := make(chan string)
	go func() { message <- "here is the random number" }()
	msg := <-message
	fmt.Println(msg)
	num1 := rand.Int()
	num2 := rand.Int()
	num3 := rand.Int()
	//print number
	fmt.Println("random number 1 :", num1)
	fmt.Println("random number 2 :", num2)
	fmt.Println("random number 3 :", num3)

}
