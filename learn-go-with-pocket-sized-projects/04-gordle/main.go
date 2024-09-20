package main

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/04-gordle/gordle/gordle"
	"os"
)

func main() {
	g := gordle.New(os.Stdin)
	g.Play()
}
