// Game holds all the information required to play a game of Gordle
package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

// New returns a Game, which can be used to play Gordle
func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    toUppercaseChars(solution),
		maxAttempts: maxAttempts,
	}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()
		fmt.Printf("Your guess is: %s\n", string(guess))

		if slices.Equal(guess, g.solution) {
			fmt.Printf(
				"🎉 You won! You found it in %d guess(es)! The word was: %s.\n",
				currentAttempt,
				string(g.solution),
			)
			return
		}
	}

	fmt.Printf("😞 You've lost! The solution was: %s. \n", string(g.solution))
}

// ask reads input until a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	for {
		fmt.Printf("Enter a %d-character guess:\n", len(g.solution))
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}
		guess := toUppercaseChars(string(playerInput))
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
	if len(guess) != len(g.solution) {
		return fmt.Errorf(
			"expected %d, got %d, %w",
			len(g.solution),
			len(guess),
			errInvalidWorldLength,
		)
	} else {
		return nil
	}
}

// toUppercaseChars converts strings to a list of uppercase characters
func toUppercaseChars(s string) []rune {
	return []rune(strings.ToUpper(s))
}
