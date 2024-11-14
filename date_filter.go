package finder

import (
	"os"

	"github.com/iveahugeship/gofind/utils"
)

type SinceDateFilter struct {
	date string
}

func (f SinceDateFilter) Match(_ string, path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	date, err := utils.ParseDate(f.date)
	if err != nil {
		return false, err
	}
	return info.ModTime().After(date), nil
}

func BySinceDate(date string) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, SinceDateFilter{
			date: date,
		})
	}
}

type UntilDateFilter struct {
	date string
}

func (f UntilDateFilter) Match(_ string, path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	date, err := utils.ParseDate(f.date)
	if err != nil {
		return false, err
	}
	return info.ModTime().Before(date), nil
}

func ByUntilDate(date string) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, UntilDateFilter{
			date: date,
		})
	}
}
