package main

import (
	"main/gordlegame"
	"os"
)

func main() {

	const maxAttempts = 6
	solution := "hello"
	game := gordlegame.New(os.Stdin, solution, maxAttempts)
	game.Play()
}
