package glob_test

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/cristalhq/glob"
)

func Example() {
	pattern := `**/*co?e*.go`
	matcher := glob.MustCompile(pattern, os.PathSeparator)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fsys := os.DirFS(filepath.Join(os.Getenv("GOPATH"), "src"))

	err := glob.Walk(ctx, fsys, ".", matcher, func(path string, d fs.DirEntry, err error) error {
		// fmt.Println(path)
		return err
	})
	if err != nil {
		panic(err)
	}

	// Output:
}

func ExampleMatch() {
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
	matcher := glob.MustCompile(pattern, '/')

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
