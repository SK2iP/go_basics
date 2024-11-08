package main

import "fmt"

//maps are like dictionaries in python

func main() {
	ages := map[string]int{
		"alice": 31,
		"charlie": 34,
	}
	fmt.Println(ages)
	fmt.Println(ages["alice"])

	new_ages := make(map[string]int)
	new_ages["alice"] = 20
	new_ages["charlie"] = 12
	fmt.Println(new_ages)

	//delete
	delete(new_ages, "alice")
	fmt.Println(new_ages)

	//check if key is in map
	age, ok := new_ages["alice"]
	fmt.Println(age, ok)

	age, ok = new_ages["charlie"]
	fmt.Println(age, ok)
}