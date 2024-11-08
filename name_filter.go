package finder

import (
	"github.com/ryanuber/go-glob"
)

type NameFilter struct {
	pattern string
}

func (f NameFilter) Match(_ string, path string) (bool, error) {
	return glob.Glob(f.pattern, path), nil
}

func ByName(pattern string) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, NameFilter{
			pattern: pattern,
		})
	}
}
