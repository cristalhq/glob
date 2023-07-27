package glob

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

const (
	sepChar = '/'
	sepStr  = "/"
)

// Glob is the representation of a compiled glob pattern.
// A Glob is safe for concurrent use by multiple goroutines.
type Glob struct {
	re *regexp.Regexp
}

// Compile parses a glob pattern and returns, if successful,
// a Glob object that can be used to match against text.
func Compile(pattern string) (*Glob, error) {
	re, err := compile(pattern)
	if err != nil {
		return nil, err
	}
	return &Glob{re: re}, nil
}

// MustCompile is like Compile but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding glob patterns.
func MustCompile(pattern string) *Glob {
	g, err := Compile(pattern)
	if err != nil {
		panic(err)
	}
	return g
}

// Match reports whether the string matches the glob pattern.
func (m *Glob) Match(s string) bool {
	return m.re.MatchString(s)
}

// Match reports whether the byte slice matches the glob pattern.
func (m *Glob) MatchBytes(b []byte) bool {
	return m.re.Match(b)
}

func compile(pattern string) (*regexp.Regexp, error) {
	root := ""
	globmask := ""

	for _, s := range strings.Split(pattern, sepStr) {
		if root == "" && hasSpecial(s) {
			root = globmask
		}

		globmask = path.Join(globmask, s)
	}

	globmask = path.Clean(globmask)

	cc := []rune(globmask)
	dirmask := ""
	filemask := ""
	staticDir := true

	for i := 0; i < len(cc); i++ {
		switch {
		case cc[i] == '*':
			staticDir = false
			if i < len(cc)-2 && cc[i+1] == '*' && cc[i+2] == sepChar {
				filemask += "(.*/)?"
				i += 2
			} else {
				filemask += "[^/]*"
			}
		default:
			c := cc[i]
			if c == sepChar || ('0' <= c && c <= '9') || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || 255 < c {
				filemask += string(c)
			} else {
				filemask += fmt.Sprintf("[\\x%02X]", c)
			}
			if staticDir {
				dirmask += string(c)
			}
		}
	}

	re, err := regexp.Compile("^" + filemask + "$")
	if err != nil {
		return nil, err
	}
	return re, nil
}

func hasSpecial(s string) bool {
	return strings.IndexByte(s, '*') != -1 ||
		strings.IndexByte(s, '{') != -1
}
