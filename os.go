package glob

import (
	"context"
	"io/fs"
)

// Walk is similar to [fs.WalkDir] but process files that are matching the given glob.
func Walk(ctx context.Context, fsys fs.FS, root string, m *Matcher, fn fs.WalkDirFunc) error {
	return fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err := ctx.Err(); err != nil {
			return err
		}
		if !m.Match(path) {
			return nil
		}
		return fn(path, d, err)
	})
}

// FindAll files that are matching the given glob.
func FindAll(ctx context.Context, fsys fs.FS, root string, m *Matcher) ([]string, error) {
	var res []string

	err := fs.WalkDir(fsys, root, func(path string, _ fs.DirEntry, err error) error {
		if err := ctx.Err(); err != nil {
			return err
		}
		if m.Match(path) {
			res = append(res, path)
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
