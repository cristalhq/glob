package glob

type Glob struct{}

func Compile(pattern string) (*Glob, error) {
	return nil, nil
}

func MustCompile(pattern string) *Glob {
	return nil
}

func (m *Glob) Match(s string) bool      { return false }
func (m *Glob) MatchBytes(b []byte) bool { return false }
