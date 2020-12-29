//matching password using regex
package main

import (
	"fmt"
	"regexp"
)

//declaration of function
func matchPassword(pass string) error {
	if len(pass) < 8 {
		return fmt.Errorf("Length < 8")
	}
func main(){	
	AZ := "[A-Z]{1}"
	num := "[0-9]{1}"
	az := "[a-z]{1}"

	symbol := "[@#!%]{1}"

	//check the condition
	if x, error := regexp.MatchString(num, pass); !x || error != nil {
		return fmt.Errorf("password need num :%v", error)
	}
	if x, error := regexp.MatchString(az, pass); !x || error != nil {
		return fmt.Errorf("password need az :%v", error)
	}
	if x, error := regexp.MatchString(AZ, pass); !x || error != nil {
		return fmt.Errorf("password need AZ :%v", error)
	}
	if x, error := regexp.MatchString(symbol, pass); !x || error != nil {
		return fmt.Errorf("password need symbol :%v", error)
	}
	return nil
}

