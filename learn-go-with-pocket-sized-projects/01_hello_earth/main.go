package main

import "fmt"

func main() {
	greeting := greet()
	fmt.Println(greeting)
}

// Return a greeting
func greet() string {
	return "Hello, world!"
}
