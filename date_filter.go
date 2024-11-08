package finder

import (
	"os"

	"github.com/iveahugeship/go-finder/utils"
)

type DateFilter struct {
	op   byte
	date string
}

func (f DateFilter) Match(_ string, path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	date, err := utils.ParseDate(f.date)
	if err != nil {
		return false, err
	}
	return utils.CompareDate(f.op, date, info.ModTime()), nil
}

func ByDate(op byte, date string) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, DateFilter{
			op:   op,
			date: date,
		})
	}
}
