package pocketlog

import (
	"fmt"
	"io"
)

// Logger is used to log information
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns a logger, ready to log at the required level
// The default output is Stdout
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

// logf prints the message to the output.
// Add decorations here, if any.
func (l *Logger) logf(format string, args ...any) {
	fmt.Fprintf(l.output, format+"\n", args...)
}

// Debugf formats and logs a message if the log level is LevelDebug
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold < LevelDebug {
		return
	}
	l.logf(format, args...)
}

// Infof formats and logs a message if the log level is LevelInfo
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold < LevelInfo {
		return
	}
	l.logf(format, args...)
}

// Errorf formats and logs a message if the log level is LevelError
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold < LevelError {
		return
	}
	l.logf(format, args...)
}
