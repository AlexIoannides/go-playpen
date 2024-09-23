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

func Test_computeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"nominal": {
			guess:    "HERTZ",
			solution: "HERTZ",
			expectedFeedback: feedback{
				correctPosition,
				correctPosition,
				correctPosition,
				correctPosition,
				correctPosition,
			},
		},
		"double character": {
			guess:    "HELLO",
			solution: "HELLO",
			expectedFeedback: feedback{
				correctPosition,
				correctPosition,
				correctPosition,
				correctPosition,
				correctPosition,
			},
		},
		"double character with wrong answer": {
			guess:    "HELLL",
			solution: "HELLO",
			expectedFeedback: feedback{
				correctPosition,
				correctPosition,
				correctPosition,
				correctPosition,
				absentCharacter,
			},
		},
		"five identical, but only two are there": {
			guess:    "LLLLL",
			solution: "HELLO",
			expectedFeedback: feedback{
				absentCharacter,
				absentCharacter,
				correctPosition,
				correctPosition,
				absentCharacter,
			},
		},
		"two identical, but not in the right position (from left to right)": {
			guess:    "HLLEO",
			solution: "HELLO",
			expectedFeedback: feedback{
				correctPosition,
				wrongPosition,
				correctPosition,
				wrongPosition,
				correctPosition,
			},
		},
		"three identical, but not in the right position (from left to right)": {
			guess:    "HLLLO",
			solution: "HELLO",
			expectedFeedback: feedback{
				correctPosition,
				absentCharacter,
				correctPosition,
				correctPosition,
				correctPosition,
			},
		},
		"one correct, one incorrect, one absent (left of the correct)": {
			guess:    "LLLWW",
			solution: "HELLO",
			expectedFeedback: feedback{
				wrongPosition,
				absentCharacter,
				correctPosition,
				absentCharacter,
				absentCharacter,
			},
		},
		"swapped characters": {
			guess:    "HOLLE",
			solution: "HELLO",
			expectedFeedback: feedback{
				correctPosition,
				wrongPosition,
				correctPosition,
				correctPosition,
				wrongPosition,
			},
		},
		"absent character": {
			guess:    "HULFO",
			solution: "HELFO",
			expectedFeedback: feedback{
				correctPosition,
				absentCharacter,
				correctPosition,
				correctPosition,
				correctPosition,
			},
		},
		"absent character and incorrect": {
			guess:    "HULPP",
			solution: "HELPO",
			expectedFeedback: feedback{
				correctPosition,
				absentCharacter,
				correctPosition,
				correctPosition,
				absentCharacter,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !tc.expectedFeedback.Equal(fb) {
				t.Errorf(
					"guess: %q, got the wrong feedback, expected %v, got %v",
					tc.guess,
					tc.expectedFeedback,
					fb,
				)
			}
		})
	}
}
