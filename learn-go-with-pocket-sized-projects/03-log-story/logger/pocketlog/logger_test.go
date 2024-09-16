package pocketlog_test

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/03-log-story/logger/logger/pocketlog"
	"os"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello, %s", "world!")
	// Output: Hello, world!
}
