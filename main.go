package main

import "fmt"

// main is the entry point of a Go program
// When you run a Go program, it starts executing from the main function
// fmt.Println is a function from the fmt package that prints text to the console
// and adds a newline character at the end
func main() {
	fmt.Println("Hello, World!")

	// Calling the add function with arguments 3 and 5
	// The result is stored in sum variable
	sum := add(3, 5)
	fmt.Println("Sum of 3 and 5 is:", sum)
}

// add is a function that takes two integer parameters and returns their sum
// Parameters:
//   - a: first integer
//   - b: second integer
// Return:
//   - int: sum of a and b
func add(a int, b int) int {
	return a + b
}
