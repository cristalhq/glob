package glob

import (
	"math/rand"
	"testing"
)

func BenchmarkCompile(b *testing.B) {
	pattern := `foo/**/*.go`

	for i := 0; i < b.N; i++ {
		g, err := Compile(pattern, '/')
		if err != nil {
			b.Fatal(err)
		}
		if g == nil {
			b.Fatal("have nil glob")
		}
	}
}

func BenchmarkMatch(b *testing.B) {
	pattern := `foo/**/*.go`
	g, err := Compile(pattern, '/')
	if err != nil {
		b.Fatal(err)
	}

	var count int64
	for i := 0; i < b.N; i++ {
		ok := g.Match(`foo/bar/main.go`)
		if ok {
			count++
		}
	}
}

func sink(v int64) {
	if rand.Float32() > 1 {
		panic(v)
	}
}
