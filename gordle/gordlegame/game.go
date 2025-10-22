package gordlegame

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

const solutionLength = 5

// Game holds all the information we need to play a game of gordle
type Game struct {
	// Reader is a struct
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

// Play runs the game
func (game *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	for currentAttempt := 1; currentAttempt <= game.maxAttempts; currentAttempt++ {
		guess := game.ask()

		if slices.Equal(guess, game.solution) {
			fmt.Printf("🎉 You won! You found it in the %d guess(es)!"+
				"The word was: %s.\n", currentAttempt, string(game.solution))
			return
		} else {
			fmt.Printf("Wrong guess, try again. Attempts left: %d\n", game.maxAttempts-currentAttempt)
		}
	}
	fmt.Printf("😔💔🙏 Sorry you have lost. The solution was: %s.\n", string(game.solution))
}

var errInvalidWordLength = fmt.Errorf("invalid guess, word doesnt " +
	"have the same number of characters as the solutions")

func (game *Game) validateGuess(guess []rune) error {
	if len(guess) != len(game.solution) {
		return fmt.Errorf("expected %d characters, got %d, %w",
			len(game.solution), len(guess), errInvalidWordLength)
	}
	return nil
}

// ask reads input until a valid suggestion is made and returned
func (game *Game) ask() []rune {

	// infinite loop
	for {
		fmt.Printf("Enter a %d-character guess:\n", len(game.solution))
		// the Reader struct has a ReadLine() method
		// that has a pointer receiver
		playerInput, _, err := game.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		// convert the guess []byte into []rune and check length
		guess := []rune(strings.ToUpper(string(playerInput)))
		err = game.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to validate your guess: %s\n", err.Error())
			continue
		} else {
			return guess
		}
	}
}

// New returns a Game, which can be used to Play!
func New(playerInput io.Reader,
	solution string,
	maxAttempts int) *Game {
	game := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    []rune(strings.ToUpper(solution)),
		maxAttempts: maxAttempts,
	}
	return game
}
