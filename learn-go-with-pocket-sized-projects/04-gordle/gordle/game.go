// Game holds all the information required to play a game of Gordle
package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const solutionLength = 5

type Game struct {
	reader *bufio.Reader
}

// New returns a Game, which can be used to play Gordle
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")
	fmt.Println("Enter a guess:")
}

// ask reads input until a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}
		guess := []rune(string(playerInput))

		if len(guess) != solutionLength {
			fmt.Fprintf(
				os.Stderr,
				"Your attempt is invalid with Gordles solution! Expected %d characters, got %d.\n",
				solutionLength,
				len(guess),
			)
		} else {
			return guess
		}
	}
}
