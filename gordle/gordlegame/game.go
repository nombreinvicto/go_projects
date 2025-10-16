package gordlegame

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const solutionLength = 5

// Game holds all the information we need to play a game of gordle
type Game struct {
	// Reader is a struct
	reader *bufio.Reader
}

// Play runs the game
func (game *Game) Play() {
	fmt.Println("Welcome to Gordle!")
	game.ask()
}

// ask reads input until a valid suggestion is made and returned
func (game *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	// infinite loop
	for {
		// the Reader struct has a ReadLine() method
		// that has a pointer receiver
		playerInput, _, err := game.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		// convert the guess into slice of runes and check length
		guess := []rune(string(playerInput))
		if len(guess) == solutionLength {
			return guess
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "Expected %d characters, got %d.Please enter again: \n", solutionLength, len(guess))
		}
	}
}

// New returns a Game, which can be used to Play!
func New(playerInput io.Reader) *Game {
	game := &Game{
		reader: bufio.NewReader(playerInput),
	}
	return game
}
