package finder

import (
	"path/filepath"
	"strings"
)

type DepthFilter struct {
	n int
}

func (f DepthFilter) Match(root string, path string) (bool, error) {
	root, path = filepath.Clean(root), filepath.Clean(path)
	sep := string(filepath.Separator)
	diff := strings.TrimPrefix(path, root)
	return f.n >= strings.Count(diff, sep)+1, nil
}

func ByDepth(n int) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, DepthFilter{
			n: n,
		})
	}
}
