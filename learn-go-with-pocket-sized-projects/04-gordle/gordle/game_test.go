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
		"3 chars in English": {
			input: "FOO",
			want:  []rune("FOO"),
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			stringReader := strings.NewReader(tc.input)
			game := New(stringReader, tc.input, 1)

			got := game.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGame_validateGuess(t *testing.T) {
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
			g := New(nil, "GUESS", 0)
			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expected %q, got %q", tc.word, tc.expected, err)
			}
		})
	}
}

func TestToUppercaseChars(t *testing.T) {
	testCases := map[string]struct {
		word     string
		expected string
	}{
		"all upper": {
			word:     "FOO",
			expected: "FOO",
		},
		"all lower": {
			word:     "foo",
			expected: "FOO",
		},
		"mixed": {
			word:     "fOo",
			expected: "FOO",
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			got := string(toUppercaseChars(tc.word))
			if got != tc.expected {
				t.Errorf("got %s, expected %s", got, tc.expected)
			}
		})
	}
}
