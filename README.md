# glob

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![version-img]][version-url]

Glob pattern matching in Go. See [Wikipedia](https://en.wikipedia.org/wiki/Glob_(programming)).

## Rationale

TODO

## Features

* Simple API.
* Dependency-free.
* Clean and tested code.
* Support for `**`.
* Walk & find for `fs.FS`.

See [docs][pkg-url] or [GUIDE.md](GUIDE.md) for more details.

## Install

Go version 1.18+

```
go get github.com/cristalhq/glob
```

## Example

```go
pattern := `foo/**/*.go`
matcher := glob.MustCompile(pattern, '/')

ctx, cancel := context.WithCancel(context.Bacg)
glob.Walk(ctx)
fmt.Printf("For pattern `%s`:\n", pattern)
for _, path := range paths {
	fmt.Printf("%15s %v\n", path, matcher.Match(path))
}
```

See examples: [example_test.go](example_test.go).

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/glob/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/glob/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/glob
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/glob
[version-img]: https://img.shields.io/github/v/release/cristalhq/glob
[version-url]: https://github.com/cristalhq/glob/releases
