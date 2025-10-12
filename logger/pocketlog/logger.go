package pocketlog

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

// Logger is used to log information
type Logger struct {
	threshold Level
	output    io.Writer
}

// logf prints the message to the output
func (l *Logger) logf(format string, args ...any) {
	// making sure we can safely write to the output
	if l.output == nil {
		l.output = os.Stdout
	}
	_, _ = fmt.Fprintf(l.output, format, args)
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold <= LevelDebug {
		l.logf(color.YellowString(format), args)
	}

}

// Infof formats and prints a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold <= LevelInfo {
		l.logf(color.GreenString(format), args)
	}

}

// Errorof formats and prints a message if the log level is error or higher
func (l *Logger) Errorof(format string, args ...any) {
	if l.threshold <= LevelError {
		l.logf(color.RedString(format), args)
	}

}

// New returns you a logger, ready to log at required threshold
// Give it a list of configuration functions to tune it at your will
// the default output is Stdout
// Struct is a value type so if passed, a copy is passed
func New(threshold Level, options ...Option) *Logger {

	logger := &Logger{
		threshold: threshold,
		output:    os.Stdout,
	}

	for _, configFunc := range options {
		configFunc(logger)
	}
	return logger
}
