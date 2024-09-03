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
	commonBooks := findCommonBooks(bookwormData)

	fmt.Println("Books read by more than one bookworm:")
	displayBooks(commonBooks)
}

// Print title and authors for a list of books
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
