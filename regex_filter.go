package finder

import (
	"regexp"
)

type RegexFilter struct {
	re *regexp.Regexp
}

func (f RegexFilter) Match(_ string, path string) (bool, error) {
	return f.re.MatchString(path), nil
}

func ByRegex(re *regexp.Regexp) Option {
	return func(f *Finder) {
		f.filters = append(f.filters, RegexFilter{
			re: re,
		})
	}
}
