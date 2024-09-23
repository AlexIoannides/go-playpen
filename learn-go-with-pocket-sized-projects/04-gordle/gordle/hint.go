package gordle

import "strings"

// hint desctibes the validity of a character in a guess.
type hint byte

// String implements the Stringer inferface.
func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜️"
	case wrongPosition:
		return "🟡"
	case correctPosition:
		return "💚"
	default:
		return "💔"
	}
}

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

// feedback is a list of hints, one per character of the word.
type feedback []hint

// String implements the Stringer interface for a slice of hints.
func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}
