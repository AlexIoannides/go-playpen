package pocketlog

// Logger is used to log information
type Logger struct {
	threshold Level
}

// New returns a logger, ready to log at the required level
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

// Debugf formats and logs a message if the log level is LevelDebug
func (l *Logger) Debugf(format string, args ...any) {
	// TODO
}

// Infof formats and logs a message if the log level is LevelInfo
func (l *Logger) Infof(format string, args ...any) {
	// TODO
}

// Errorf formats and logs a message if the log level is LevelError
func (l *Logger) Errorf(format string, args ...any) {
	// TODO
}
