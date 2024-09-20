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

	guess := g.ask()

	fmt.Printf("Your guess is: %s\n", string(guess))
}

// ask reads input until a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}
		guess := []rune(string(playerInput))
		err = g.validateGuess(guess)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"You attempt is invalid with Gordle's solution: %s\n",
				err.Error(),
			)
		} else {
			return guess
		}
	}
}

// errInvalidWorldLength is returned when the guess has the wrong number of characters
var errInvalidWorldLength = fmt.Errorf(
	"invalid guess, word doesn't have the same number of characters as the solution",
)

// validateGuess checks a guess to ensure that it's a viable (and this valid) guess.
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf(
			"expected %d, got %d, %w",
			solutionLength,
			len(guess),
			errInvalidWorldLength,
		)
	} else {
		return nil
	}
}
