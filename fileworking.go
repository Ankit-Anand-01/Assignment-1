//file working in go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//create a text file
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	//write this in created file
	file.WriteString("Hi...Ankit is here")
	file.Close()
	stream, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)
}
