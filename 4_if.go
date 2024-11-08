package main

import "fmt"

func main(){
	a := 10
	b := 20
	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is equal to b")
	}

	email := "Salam"

	if len := len(email); len > 3 {
		fmt.Println("Email is too long")
	} else {
		fmt.Println("Email is short")
	}
}