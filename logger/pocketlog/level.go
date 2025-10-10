package pocketlog

// Level is the available logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, used for debugging
	LevelDebug Level = iota

	// LevelInfo represents logging level that contains info deemed valuable
	LevelInfo

	// LevelError represents logging level used to trace errors
	LevelError
)
