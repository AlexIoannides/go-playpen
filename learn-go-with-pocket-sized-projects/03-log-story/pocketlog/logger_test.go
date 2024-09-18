package pocketlog_test

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/03-log-story/pocketlog"
	"testing"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world!")
	// Output: Hello, world!
}

// testWriter is a struct that implements io.Writer
// We use it to validate that we can write to a specific output
type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents += string(p)
	return len(p), nil
}

const (
	debugMessage = "Foo"
	infoMessage  = "Bar"
	errorMessage = "La La La"
)

func TestLogger_DebugInfoError(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	testCases := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: debugMessage + "\n" + infoMessage + "\n" + errorMessage + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: infoMessage + "\n" + errorMessage + "\n",
		},

		"error": {
			level:    pocketlog.LevelError,
			expected: errorMessage + "\n",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			mockWriter := testWriter{contents: ""}
			logOptions := []pocketlog.Option{pocketlog.WithOutput(&mockWriter)}
			log := pocketlog.New(testCase.level, logOptions...) // we can unpack a slice into variadic args

			log.Debugf(debugMessage)
			log.Infof(infoMessage)
			log.Errorf(errorMessage)

			want := testCase.expected
			got := mockWriter.contents
			if got != want {
				t.Errorf("\n- want log message:\n%s\n- got:\n%s", want, got)
			}
		})
	}

}
