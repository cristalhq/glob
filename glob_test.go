package glob

import (
	"math/rand"
	"testing"
)

func BenchmarkCompile(b *testing.B) {
	pattern := `foo/**/*.go`

	var count int64
	for i := 0; i < b.N; i++ {
		g, err := Compile(pattern, '/')
		if err != nil {
			b.Fatal(err)
		}
		if g != nil {
			count++
		}
	}
	sink(b, count)
}

func BenchmarkMatch(b *testing.B) {
	pattern := `foo/**/*.go`
	g, err := Compile(pattern, '/')
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	var count int64
	for i := 0; i < b.N; i++ {
		ok := g.Match(`foo/bar/main.go`)
		if ok {
			count++
		}
	}
	sink(b, count)
}

func sink(tb testing.TB, v int64) {
	if rand.Float32() > 1 {
		tb.Fatal(v)
	}
}
