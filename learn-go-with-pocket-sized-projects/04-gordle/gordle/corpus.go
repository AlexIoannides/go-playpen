package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// ErrCorpusIsEmpty is returned when we cannot find any valid solution in the given corpus.
const ErrCorpusIsEmpty = corpusError("corpus is empty")

// set the random number generator at the package level
var rng = rand.New(rand.NewSource(42))

// ReadCorpus reads the file located at the given path and returns a list of words.
func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading: %w", path, err)
	}

	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	// we expect the corpus to be a line- or space-separated list of words
	words := strings.Fields(string(data))

	return words, nil
}

// pickWord returns a random word from the corpus
func pickWord(corpus []string) string {
	// rand.Seed is only necessary if your version of Go is before 1.20.
	// It's best not to have it, if you're using go 1.21 or more recent.
	//nolint:staticcheck // Still present to explain how things used to be
	index := rng.Intn(len(corpus))

	return corpus[index]
}
