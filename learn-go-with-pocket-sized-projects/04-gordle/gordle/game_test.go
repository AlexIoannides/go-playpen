package gordle

import (
	"slices"
	"strings"
	"testing"
)

func TestGame_ask(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  []rune
	}{
		"5 chars in English": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 chars in Arabic": {
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		"3 chars in Japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			stringReader := strings.NewReader(testCase.input)
			game := New(stringReader)

			got := game.ask()
			if !slices.Equal(got, testCase.want) {
				t.Errorf("got = %v, want %v", string(got), string(testCase.want))
			}
		})
	}
}
