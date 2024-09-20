package gordle

import (
	"errors"
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

func TestGame_validationGuess(t *testing.T) {
	testCases := map[string]struct {
		word     []rune
		expected error
	}{
		"nominal": {
			word:     []rune("GUESS"),
			expected: nil,
		},
		"too long": {
			word:     []rune("POCKET"),
			expected: errInvalidWorldLength,
		},
		"too short": {
			word:     []rune("FOO"),
			expected: errInvalidWorldLength,
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			g := New(nil)
			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expected %q, got %q", tc.word, tc.expected, err)
			}
		})
	}
}
