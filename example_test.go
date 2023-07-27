package glob_test

import (
	"fmt"

	"github.com/cristalhq/glob"
)

func Example() {
	paths := []string{
		`foo/bar`,
		`foo/bar/a.go`,
		`foo/bar/baz`,
		`foo/bar/main.go`,
		`foo/bar/baz.txt`,
		`foo/bar/baz/noo.txt`,
		`foo/baz`,
		`foo/baz.go`,
	}

	pattern := `foo/**/*.go`
	matcher := glob.MustCompile(pattern)

	fmt.Printf("For pattern `%s`:\n", pattern)
	for _, path := range paths {
		fmt.Printf("%15s %v\n", path, matcher.Match(path))
	}

	// Output:
	// 	For pattern `foo/**/*.go`:
	//         foo/bar false
	//    foo/bar/a.go true
	//     foo/bar/baz false
	// foo/bar/main.go true
	// foo/bar/baz.txt false
	// foo/bar/baz/noo.txt false
	//         foo/baz false
	//      foo/baz.go true
}
