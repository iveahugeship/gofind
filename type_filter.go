package finder

import (
	"os"

	"github.com/iveahugeship/gofind/utils/file/ftype"
)

type TypeFilter struct {
	ftype ftype.FileType
}

func (f TypeFilter) Match(_ string, path string) (bool, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return false, err
	}
	return f.ftype.IsTypeOf(info), nil
}

func ByType(ftype ftype.FileType) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, TypeFilter{
			ftype: ftype,
		})
	}
}
