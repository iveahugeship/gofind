package finder

import (
	"github.com/iveahugeship/gofind/utils/file/mtime"
)

type SinceMtimeFilter struct {
	mtime mtime.ModTime
}

func (f SinceMtimeFilter) Match(_ string, path string) (bool, error) {
	return f.mtime.NewerOf(path)
}

func BySinceMtime(s string) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, SinceMtimeFilter{
			mtime: mtime.NewModTime(s),
		})
	}
}

type UntilMtimeFilter struct {
	mtime mtime.ModTime
}

func (f UntilMtimeFilter) Match(_ string, path string) (bool, error) {
	return f.mtime.OlderOf(path)
}

func ByUntilMtime(s string) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, UntilMtimeFilter{
			mtime: mtime.NewModTime(s),
		})
	}
}
