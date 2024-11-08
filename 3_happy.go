package main

import "fmt"

func main(){
	congrats := 12
	// %T is used to print the type of the variable
	fmt.Printf("%T\n",congrats)

	avg , display := 12.5, "Hello"

	fmt.Printf("%T %T\n",avg,display)

	// Converting float to int
	float_avg := int(avg)

	fmt.Printf("%T\n",float_avg)


	// Constants
	const pi = 3.14
	fmt.Printf("%f %T\n",pi,pi)

	const name = "John"
	const openrate = 12.34

	fmt.Printf("%s %.1f\n",name,openrate)

}