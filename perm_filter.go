package finder

import (
	"io/fs"
	"os"
)

type PermFilter struct {
	perm fs.FileMode
}

func (f PermFilter) Match(_ string, path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.Mode().Perm() == f.perm, nil
}

func ByPerm(perm fs.FileMode) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, PermFilter{
			perm: perm,
		})
	}
}
