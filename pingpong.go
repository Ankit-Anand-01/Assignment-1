//ping pong program in go using channel
package main

import (
	"fmt"
	"time"
)

//pinger prints a ping
//and waits for pong
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

// ponger prints a pong and waits for a ping
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	// The main goroutine starts the ping/pong by sending into the ping channel
	ping <- 1

	for {
		// Block the main thread
		//untill any interruption
		time.Sleep(time.Second)
	}
}
