package main 

import "fmt"

func concat(s1 string,s2 string) string{
	return s1 + s2
}
// now both variables are of type int
func sub(a, b int) int{
	return a - b
}

func add_sub(a,b int) (int,int){
	return a+b,a-b
}

func mul_div(a,b int)(mul,div int){
	mul,div = a*b,a/b
	// you can return or you can return like this
	//return mul,div
	return
}
func main(){
	fmt.Println(concat("Hello","World"))
	fmt.Println(sub(12,-9))
	fmt.Println(add_sub(12,9))
	add, _ := add_sub(12,9)
	fmt.Println(add)
	fmt.Println(mul_div(12,9))
}