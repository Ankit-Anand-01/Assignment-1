//structure in Go
package main

import "fmt"

//defining a struct type
//employee structure
type employee struct {
	name       string
	age        int
	city       string
	state      string
	experience string
}

func main() {
	// Naming fields while
	// initializing a struct
	emp1 := employee{
		name:       "Dhruv",
		age:        25,
		city:       "Jaipur",
		state:      "Rajasthan",
		experience: "2-years",
	}
	//without using fields name
	emp2 := employee{"Animesh", 32, "Pune", "Maharastra", "8-years"}
	emp3 := employee{"Harry", 28, "Roorkee", "Uttarakhand", "3-years"}
	fmt.Println("employee 1:", emp1)
	fmt.Println("employee 2:", emp2)
	fmt.Println("employee 3:", emp3)

}
