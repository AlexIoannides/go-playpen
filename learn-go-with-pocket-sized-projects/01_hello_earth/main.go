package main

import "fmt"

func main() {
	greeting := greet("EN")
	fmt.Println(greeting)
}

// Language represents a language's two character code
type language string

// Return a greeting
func greet(lang language) string {
	switch lang {
	case "EN":
		return "Hello, world!"
	case "FR":
		return "Bonjour, le monde!"
	default:
		return ""
	}
}
