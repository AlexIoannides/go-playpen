package main

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/04-gordle/gordle/gordle"
	"fmt"
	"os"
)

func main() {
	corpus, err := gordle.ReadCorpus("testdata/corpus/english.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err)
	}

	g, err := gordle.New(os.Stdin, corpus, 5)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to start fame: %s", err)
	}

	g.Play()
}
