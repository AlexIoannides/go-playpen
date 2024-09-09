package main

type set map[Book]struct{}

// Does set contain book?
func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}

// Create set of Books from list of Books
func newSet(books *[]Book) *set {
	bookSet := make(set)
	for _, book := range *books {
		bookSet[book] = struct{}{}
	}
	return &bookSet
}

type Recommendation struct {
	Book  Book
	Score float64
}

func recommend(allBookworms []Bookworm, targetBookworm Book, n int) []Recommendation {

	return make([]Recommendation, 0)
}
