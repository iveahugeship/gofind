package finder

import (
	"github.com/iveahugeship/gofind/utils/file"
)

type SinceMtimeFilter struct {
	mtime file.ModTime
}

func (f SinceMtimeFilter) Match(_ string, path string) (bool, error) {
	return f.mtime.NewerOf(path)
}

func BySinceMtime(t string) Option {
	return func(f *Finder) {
		mtime, _ := file.NewModTime(t)
		f.filters = append(f.filters, SinceMtimeFilter{
			mtime: mtime,
		})
	}
}

type UntilMtimeFilter struct {
	mtime file.ModTime
}

func (f UntilMtimeFilter) Match(_ string, path string) (bool, error) {
	return f.mtime.OlderOf(path)
}

func ByUntilMtime(t string) Option {
	return func(f *Finder) {
		mtime, _ := file.NewModTime(t)
		f.filters = append(f.filters, UntilMtimeFilter{
			mtime: mtime,
		})
	}
}