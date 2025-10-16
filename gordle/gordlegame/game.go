package gordlegame

import "fmt"

// Game holds all the information we need to play a game of gordle
type Game struct{}

// Play runs the game
func (game *Game) Play() {
	fmt.Println("Welcome to Gordle!")
	fmt.Printf("Enter a guess:\n")
}

// New returns a Game, which can be used to Play!
func New() *Game {
	game := &Game{}
	return game
}
