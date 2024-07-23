package main

import "fmt"

func main() {
	greeting := greet("EN")
	fmt.Println(greeting)
}

// Language represents a language's two character code
type language string

// Hold a greeting for each supported language
var phrasebook = map[language]string{
	"EN": "Hello, world!",
	"FR": "Bonjour, le monde!",
}

// Return a greeting
func greet(lang language) string {
	msg, ok := phrasebook[lang]
	if !ok {
		return fmt.Sprintf("unsupported language: %s", lang)
	}
	return msg
}
