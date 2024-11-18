package finder

import (
	"os"

	"github.com/iveahugeship/gofind/utils/file"
)

type TypeFilter struct {
	ftype file.FileType
}

func (f TypeFilter) Match(_ string, path string) (bool, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return false, err
	}
	return f.ftype.IsTypeOf(info), nil
}

func ByType(ftype file.FileType) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, TypeFilter{
			ftype: ftype,
		})
	}
}
