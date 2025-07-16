package glob

import (
	"fmt"
	"regexp"
	"strings"
)

// Deprecated: use [Matcher].
type Glob = Matcher

// Matcher is the representation of a compiled glob pattern.
// A Matcher is safe for concurrent use by multiple goroutines.
type Matcher struct {
	re *regexp.Regexp
}

// Compile parses a glob pattern and returns, if successful,
// a [Matcher] object that can be used to match against text.
func Compile(pattern string, sep rune) (*Matcher, error) {
	re, err := compile(pattern, sep)
	if err != nil {
		return nil, err
	}
	return &Matcher{re: re}, nil
}

// MustCompile is like [Compile] but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding glob patterns.
func MustCompile(pattern string, sep rune) *Matcher {
	g, err := Compile(pattern, sep)
	if err != nil {
		panic(err)
	}
	return g
}

// Match reports whether the string matches the glob pattern.
func (m *Matcher) Match(s string) bool {
	return m.re.MatchString(s)
}

// Match reports whether the byte slice matches the glob pattern.
func (m *Matcher) MatchBytes(b []byte) bool {
	return m.re.Match(b)
}

func compile(pattern string, sep rune) (*regexp.Regexp, error) {
	cc := []rune(pattern)
	var b strings.Builder

	for i := 0; i < len(cc); i++ {
		switch c := cc[i]; c {
		case '?':
			b.WriteByte('.')

		case '*':
			if i < len(cc)-2 && cc[i+1] == '*' && cc[i+2] == sep {
				b.WriteString("(.*/)?")
				i += 2
			} else {
				b.WriteString("[^/]*")
			}

		// TODO(oleg): support more of this
		// case '[':
		// case '{':
		// case ',':
		// case '!':
		// case '-':

		case '\\':
			b.WriteByte('\\')

		default:
			if c == sep || isASCII(c) {
				b.WriteRune(c)
			} else {
				fmt.Fprintf(&b, "[\\x%02X]", c)
			}
		}
	}

	re, err := regexp.Compile("^" + b.String() + "$")
	if err != nil {
		return nil, err
	}
	return re, nil
}

func isASCII(c rune) bool {
	return ('0' <= c && c <= '9') ||
		('a' <= c && c <= 'z') ||
		('A' <= c && c <= 'Z') ||
		255 < c
}
