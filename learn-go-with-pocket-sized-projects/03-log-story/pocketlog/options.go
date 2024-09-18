// Option defines a functional option to Logger.
package pocketlog

import "io"

type Option func(*Logger)

// WithOutput returns a configuration function that sets the output of logs.
func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
}
