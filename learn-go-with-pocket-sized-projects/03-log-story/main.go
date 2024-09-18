package main

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/03-log-story/pocketlog"
	"os"
)

func main() {
	// Create a logger that uses the default os.Stdout implementation of io.Writer
	log := pocketlog.New(pocketlog.LevelDebug)

	// Log debug message
	log.Debugf("This is a Debug log message from %s", "Alex Ioannides")

	// Log Info message
	log.Infof("This is an Info log message from %s", "Alex Ioannides")

	// We can also create loggers with custom writers
	log = pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(os.Stdout))

	// Log Error message
	log.Infof("This is an Error log message from %s", "Alex Ioannides")

}
