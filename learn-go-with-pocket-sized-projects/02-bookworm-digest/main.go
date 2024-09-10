package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "test-data/bookworms.json", "Path to JSON data file")
	flag.Parse()

	bookwormData, err := loadUserData(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}
	commonBooks := findCommonBooks(bookwormData)

	fmt.Println("Books read by more than one bookworm:")
	displayBooks(commonBooks)

	newBookworm := Bookworm{
		Name: "Kevin",
		Books: []Book{
			{Author: "Margaret Atwood", Title: "Oryx and Crake"},
			{Author: "Charlotte Brontë", Title: "Jane Eyre"},
		},
	}
	recommendations := recommend(bookwormData, newBookworm, 1)
	var recommendedBooks = make([]Book, len(recommendations))
	for i, recommendation := range recommendations {
		recommendedBooks[i] = recommendation.Book
	}
	fmt.Printf("Recommended books for new bookworm %s:\n", newBookworm.Name)
	displayBooks((recommendedBooks))
}

// Print title and authors for a list of books
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
