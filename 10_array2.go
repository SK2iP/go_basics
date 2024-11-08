package main

import "fmt"

//second way to argument a slice

func print_slice(s []string){
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

//first way to argument a slice

func initialize_slice(s ...string){
	for i := 0; i < len(s); i++ {
		s[i] = fmt.Sprintf("%d", i)
	}
	fmt.Println("slice:", s)
}
func main() {
	slices := make([]string, 3,5)
	fmt.Println("emp:", slices)
	fmt.Println("len:", len(slices))
	fmt.Println("cap:", cap(slices))
	//cap is the capacity of the slice which is the maximum number of elements that the slice can hold before it needs to allocate more memory
	//len is the length of the slice which is the number of elements that the slice currently holds
	initialize_slice(slices...)
	print_slice(slices)

	//append to
	slices = append(slices, "a")
	fmt.Println("slice:", slices)

	//append to
	slices = append(slices, "b", "c", "d")
	fmt.Println("slice:", slices)
}