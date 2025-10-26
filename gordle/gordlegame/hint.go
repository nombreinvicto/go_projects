package gordlegame

import "strings"

// hint describes the validity of a character in a word
type hint byte

// feedback is a list of hints one per character of the word
type feedback []hint

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

// String implements the Stringer interface
// Types that implement the Stringer are printed the same as strings by Printf
func (hint hint) String() string {
	switch hint {
	case absentCharacter:
		return "⬛️"
	case wrongPosition:
		return "⚠️"
	case correctPosition:
		return "✅"
	default:
		// this shud never happen
		return "💔"
	}
}

func (feedback feedback) StringConcat() string {
	sb := strings.Builder{}
	for _, hint := range feedback {
		sb.WriteString(hint.String())
	}
	return sb.String()
}
