package main

import (
	"testing"
)

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

// test helper function for asserting if two lists of Books are equal
func equalBooks(got, want []Book) bool {
	if len(want) != len(got) {
		return false
	}

	for i := range want {
		if got[i] != want[i] {
			return false
		}
	}

	return true
}

// test helper function for asserting if two lists of Bookworms are equal
func equalBookworms(got, want []Bookworm) bool {
	if len(got) != len(want) {
		return false
	}

	for i := range got {
		if got[i].Name != want[i].Name {
			return false
		}

		if !equalBooks(got[i].Books, want[i].Books) {
			return false
		}
	}

	return true
}

func TestLoadUserData(t *testing.T) {
	type testCase struct {
		dataFile string
		want     []Bookworm
		wantErr  bool
	}

	tests := map[string]testCase{
		"file_exists": {
			dataFile: "test-data/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file_does_not_exist": {
			dataFile: "test-data/foo.json",
			want:     nil,
			wantErr:  true,
		},
	}

	for testName, testData := range tests {
		t.Run(
			testName,
			func(t *testing.T) {
				gotData, gotErr := loadUserData(testData.dataFile)

				if testData.wantErr && gotErr == nil {
					t.Fatalf("expected an error, got none")
				}

				if !testData.wantErr && gotErr != nil {
					t.Fatalf("expected no error, got one %s", gotErr.Error())
				}

				if !equalBookworms(testData.want, gotData) {
					t.Fatalf("different results got %v, expected %v", gotData, testData.want)
				}
			},
		)
	}
}

func TestFindCommonBooks(t *testing.T) {
	type testCase struct {
		data []Bookworm
		want []Book
	}

	tests := map[string]testCase{
		"common_books": {
			data: []Bookworm{
				{Name: "A", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "B", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: []Book{handmaidsTale},
		},
		"no_common_books": {
			data: []Bookworm{
				{Name: "A", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "B", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: []Book{},
		},
		"empty_bookworms": {
			data: []Bookworm{},
			want: []Book{},
		},
	}

	for testName, testData := range tests {
		t.Run(
			testName,
			func(t *testing.T) {
				want := testData.want
				got := findCommonBooks(testData.data)
				if !equalBooks(got, testData.want) {
					t.Fatalf("different results got %v, expected %v", got, want)
				}
			},
		)
	}
}
