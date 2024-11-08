package finder

import (
	"io/fs"
	"path/filepath"
)

type filters []Filter

type Finder struct {
	root    string
	filters filters
}

func NewFinder(root string, opts ...Option) *Finder {
	f := &Finder{
		root: root,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f Finder) Find() ([]string, error) {
	var hits []string
	filepath.WalkDir(f.root, func(path string, d fs.DirEntry, err error) error {
		isMatch, e := f.match(path)
		if e != nil {
			return e
		}
		if isMatch {
			hits = append(hits, path)
		}
		return err
	})
	return hits, nil
}

func (f Finder) match(path string) (bool, error) {
	for _, filter := range f.filters {
		isMatch, err := filter.Match(f.root, path)
		if err != nil {
			return false, err
		}
		if !isMatch {
			return false, nil
		}
	}
	return true, nil
}
