package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("----------------")
	//no all condition for loop 
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}
	//while loop
	j:=0
	for j < 9 {
		fmt.Println("This is an infinite loop")
		j++
	}
}