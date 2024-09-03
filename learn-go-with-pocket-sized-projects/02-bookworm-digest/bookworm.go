package main

import (
	"encoding/json"
	"os"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// Load data from JSON file
func loadUserData(filePath string) ([]Bookworm, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bookworms []Bookworm
	err = json.NewDecoder(file).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

// Find books common to more than one bookworm
func findCommonBooks(bookworms []Bookworm) []Book {
	bookCounter := map[Book]uint{}
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			bookCounter[book] += 1
		}
	}
	commonBooks := []Book{}
	for book, bookCount := range bookCounter {
		if bookCount > 1 {
			commonBooks = append(commonBooks, book)
		}
	}
	return commonBooks
}
