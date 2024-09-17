package pocketlog_test

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/03-log-story/pocketlog"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world!")
	// Output: Hello, world!
}
