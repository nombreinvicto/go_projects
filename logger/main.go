package main

import (
	"main/pocketlog"
	"os"
)

func main() {
	logger := pocketlog.New(pocketlog.LevelError,
		pocketlog.WithOutput(os.Stdout))
	//logger.Infof("A little copying is better than a little dependency")
	logger.Errorof("Make the zero %d value useful.", 0)
}
