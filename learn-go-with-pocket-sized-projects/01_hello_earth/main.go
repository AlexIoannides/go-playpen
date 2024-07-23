package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "EN", "The requred language - e.g., EN, FR")
	flag.Parse()
	greeting := greet(language(lang))
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
