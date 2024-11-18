package finder

import (
	"github.com/iveahugeship/gofind/utils/file"
)

type SinceMtimeFilter struct {
	mtime file.ModTime
}

func (f SinceMtimeFilter) Match(_ string, path string) (bool, error) {
	return f.mtime.Before(path)
}

func BySinceMtime(d string) Option {
	return func(f *Finder) {
		mtime, _ := file.NewModTime(d)
		f.filters = append(f.filters, SinceMtimeFilter{
			mtime: mtime,
		})
	}
}

type UntilMtimeFilter struct {
	mtime file.ModTime
}

func (f UntilMtimeFilter) Match(_ string, path string) (bool, error) {
	return f.mtime.After(path)
}

func ByUntilDate(d string) Option {
	return func(f *Finder) {
		mtime, _ := file.NewModTime(d)
		f.filters = append(f.filters, UntilMtimeFilter{
			mtime: mtime,
		})
	}
}
