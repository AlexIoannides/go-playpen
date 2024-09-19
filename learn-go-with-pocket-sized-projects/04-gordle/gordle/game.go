// Game holds all the information required to play a game of Gordle
package gordle

import "fmt"

type Game struct{}

// New returns a Game, which can be used to play Gordle
func New() *Game {
	g := &Game{}

	return g
}

// Play runs the game.
func (g Game) Play() {
	fmt.Println("Welcome to Gordle!")
	fmt.Println("Enter a guess:")
}
