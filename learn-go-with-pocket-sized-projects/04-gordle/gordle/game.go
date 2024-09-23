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

		fb := computeFeedback(guess, g.solution)
		fmt.Println(fb)

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

// computeFeedback verifies every character of the guess against the solution.
func computeFeedback(guess, solution []rune) feedback {
	// initialise holders for marks
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		fmt.Fprintf(
			os.Stderr,
			"Internal error! Guess and solution have different lengths: %d vs %d",
			len(guess),
			len(solution),
		)
		return result
	}

	// check for correct letters
	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctPosition
			used[posInGuess] = true
		}
	}

	// look for letters in the wrong position
	for posInGuess, character := range guess {
		if result[posInGuess] != absentCharacter {
			// The character has already been marked, ignore it.
			continue
		}

		for posInSolution, target := range solution {
			if used[posInSolution] {
				// The letter of the solution is already assigned to a letter of the guess.
				// Skip to the next letter of the solution.
				continue
			}
			if character == target {
				result[posInGuess] = wrongPosition
				used[posInSolution] = true
				// Skip to the next letter of the guess.
				break
			}
		}
	}
	return result
}
