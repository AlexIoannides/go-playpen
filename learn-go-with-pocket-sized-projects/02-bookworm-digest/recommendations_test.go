package main

import "testing"

func TestSet_Contains(t *testing.T) {
	s := set{handmaidsTale: {}, oryxAndCrake: {}, theBellJar: {}}
	t.Run("contains_matching_book", func(t *testing.T) {
		if !s.Contains(theBellJar) {
			t.Errorf("expected set to contain book: %v", theBellJar)
		}
	})
	t.Run("contains_no_matching_book", func(t *testing.T) {
		if s.Contains(janeEyre) {
			t.Errorf("expected set not to contain book: %v", janeEyre)
		}
	})
}

func TestNewSet(t *testing.T) {
	bookList := []Book{handmaidsTale, handmaidsTale, oryxAndCrake}
	bookSet := *newSet(&bookList)

	if len(bookSet) != 2 {
		t.Fatal("expected set to have length 2")
	}

	_, ok := bookSet[handmaidsTale]
	if !ok {
		t.Fatalf("expected book %v to be in set", handmaidsTale)
	}

	_, ok = bookSet[oryxAndCrake]
	if !ok {
		t.Fatalf("expected book %v to be in set", oryxAndCrake)
	}
}

func TestRecommend(t *testing.T) {
	targetBookworm := Bookworm{Name: "Kevin", Books: []Book{oryxAndCrake, janeEyre}}
	t.Run("recommendations_when_books_in_common", func(t *testing.T) {
		wantBook := handmaidsTale
		recommendation := recommend(bookWorms, targetBookworm, 1)
		gotBook := recommendation[0].Book
		if gotBook != wantBook {
			t.Errorf("Expected recommended book: %v, got: %v", wantBook, gotBook)
		}
	})
}
