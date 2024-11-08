package main

import "fmt"

// Define a function type
type Operation func(int, int) int

// A function that takes another function as an argument
func compute(a int, b int, op Operation) int {
    return op(a, b)
}

// Example functions that match the Operation type
func add(x int, y int) int {
    return x + y
}

func multiply(x int, y int) int {
    return x * y
}

func main() {
    a, b := 5, 3

    // Use the compute function with different operations
    fmt.Println("Addition:", compute(a, b, add))
    fmt.Println("Multiplication:", compute(a, b, multiply))
}