package gordle_test

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/04-gordle/gordle/gordle"
	"errors"
	"testing"
)

func TestReadCorpus(t *testing.T) {
	tt := map[string]struct {
		file   string
		length int
		err    error
	}{
		"English corpus": {
			file:   "testdata/corpus/english.txt",
			length: 34,
			err:    nil,
		},
		"empty corpus": {
			file:   "testdata/corpus/empty.txt",
			length: 0,
			err:    gordle.ErrCorpusIsEmpty,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			words, err := gordle.ReadCorpus(tc.file)
			if !errors.Is(tc.err, err) {
				t.Errorf("expected err %v, got %v", tc.err, err)
			}

			if tc.length != len(words) {
				t.Errorf("expected %d, got %d", tc.length, len(words))
			}
		})
	}
}
