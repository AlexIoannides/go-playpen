package main

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/04-gordle/gordle/gordle"
	"os"
)

func main() {
	solution := "foobar"
	g := gordle.New(os.Stdin, solution, 5)
	g.Play()
}
