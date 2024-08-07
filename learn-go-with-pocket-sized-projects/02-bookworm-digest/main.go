package main

import (
	"fmt"
	"os"
)

func main() {
	bookwormData, err := loadUserData("test-data/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}
	printBookworms(bookwormData)
}

// Print bookworm data to stdout
func printBookworms(data []Bookworm) {
	for _, bookworm := range data {
		fmt.Println(bookworm)
	}
}
