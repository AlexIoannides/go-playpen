package main

import (
	"math"
	"sort"
)

type set map[Book]struct{}

// Does set contain book?
func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}

// Create set of Books from list of Books
func newSet(books *[]Book) *set {
	bookSet := make(set, 0)
	for _, book := range *books {
		bookSet[book] = struct{}{}
	}
	return &bookSet
}

type Recommendation struct {
	Book  Book
	Score float64
}

// Define sort.Interface for a collection of Recommendations
type Recommedations []Recommendation

func (r Recommedations) Len() int { return len(r) }

func (r Recommedations) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Recommedations) Less(i, j int) bool {
	return r[i].Score < r[j].Score
}

// Recommend books for a bookworm
func recommend(allBookworms []Bookworm, targetBookworm Bookworm, n int) []Recommendation {
	booksOnShelf := newSet(&targetBookworm.Books)
	recommendScores := map[Book]float64{}

	for _, bookWorm := range allBookworms {
		if bookWorm.Name == targetBookworm.Name {
			continue
		}

		var similarity float64
		for _, book := range bookWorm.Books {
			if booksOnShelf.Contains(book) {
				similarity++
			}
		}

		if similarity == 0 {
			continue
		}

		score := math.Log(similarity) + 1
		for _, book := range bookWorm.Books {
			if !booksOnShelf.Contains(book) {
				recommendScores[book] += score
			}
		}
	}

	var allRecommendations []Recommendation
	for book, score := range recommendScores {
		allRecommendations = append(allRecommendations, Recommendation{book, score})
	}
	sort.Sort(Recommedations(allRecommendations))

	if len(allRecommendations) > 0 {
		if n > len(allRecommendations) {
			return allRecommendations
		} else {
			return allRecommendations[:n]
		}
	} else {
		return make([]Recommendation, 0)
	}
}
