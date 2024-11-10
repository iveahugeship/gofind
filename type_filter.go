package finder

import (
	"os"

	"github.com/iveahugeship/gofind/utils"
)

type TypeFilter struct {
	ftype utils.FileType
}

func (f TypeFilter) Match(_ string, path string) (bool, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return false, err
	}
	switch f.ftype.String {
	case 'b':
		return info.Mode().Type() == utils.DevType.Mode, nil
	case 'c':
		return info.Mode().Type() == utils.CharDevType.Mode, nil
	case 'l':
		return info.Mode().Type() == utils.SymlinkType.Mode, nil
	case 's':
		return info.Mode().Type() == utils.SymlinkType.Mode, nil
	case 'p':
		return info.Mode().Type() == utils.NamedPipeType.Mode, nil
	case 'd':
		return info.Mode().Type() == utils.DirType.Mode, nil
	case 'f':
		return info.Mode().Type() == utils.RegularType.Mode, nil
	default:
		return false, nil
	}
}

func ByType(ftype utils.FileType) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, TypeFilter{
			ftype: ftype,
		})
	}
}
