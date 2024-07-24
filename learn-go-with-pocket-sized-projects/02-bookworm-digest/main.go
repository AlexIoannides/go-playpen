package main

import (
	"fmt"
	"os"
)

func main() {
	bookwormData, err := loadUserData("bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(bookwormData)
}
