//server and worker channel in go
package main

import (
	"fmt"
	"time"
)

//declaration of function
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

//main function
func main() {
	//creating channel
	job := make(chan int, 10)
	result := make(chan int, 10)
	for w := 1; w <= 2; w++ {
		go worker(w, job, result)
	}
	for j := 1; j <= 5; j++ {
		job <- j
	}
	//closing the channel
	close(job)
	for a := 1; a <= 5; a++ {
		<-result
	}
}
