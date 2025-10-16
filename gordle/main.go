package main

import (
	"main/gordlegame"
	"os"
)

func main() {
	g := gordlegame.New(os.Stdin)
	g.Play()
}
