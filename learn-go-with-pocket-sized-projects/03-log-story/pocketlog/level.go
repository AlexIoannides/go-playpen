package pocketlog

// Level represents an available logging level
type Level byte

const (
	// LevelDebug is the lowest available logging level, mostly used for debugging.
	LevelDebug Level = iota

	// LevelInfo is a logging level for all useful happy-path application information.
	LevelInfo

	// LevelError is a logging level for flagging problems encountered during application runtime.
	LevelError
)
